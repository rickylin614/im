package msggateway

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/vmihailenco/msgpack/v5"

	"im/internal/models/po"
	"im/internal/pkg/config"
	"im/internal/pkg/consts/enums"
	"im/internal/pkg/consts/topic"
	"im/internal/pkg/prom"
	"im/internal/pkg/signalctx"
	"im/internal/util/errs"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
)

type IWsManager interface {
	Run(ctx *signalctx.Context) error
	// wsHandler(w http.ResponseWriter, r *http.Request)
	NewClient(ctx *gin.Context, conn IClientConn, isBackground, isCompress bool, token string) *Client
	GetUserAllConnects(userID string) ([]*Client, bool)
	GetUserPlatformCons(userID string, platform int) ([]*Client, bool, bool)
	GetUserAll() []string
	Validate(s any) error
	// SetCacheHandler(cache cache.MsgModel)
	// SetDiscoveryRegistry(client discoveryregistry.SvcDiscoveryRegistry)
	KickUserConn(client *Client) error
	Register(c *Client) error
	UnRegister(c *Client)

	SendMessage2Client(context context.Context, data *po.Message) error
	SendMessage2Queue(context context.Context, msg *po.Message) error
	// SetKickHandlerInfo(i *kickHandler)
	Compressor
	Encoder
}

type digIn struct {
	dig.In

	Conf      *config.Config
	Ctx       *signalctx.Context
	Prom      *prom.Manager
	Publisher message.Publisher
}

type WsManager struct {
	wsMaxConnNum    int64             // 設定檔限制的最大連線數量
	registerChan    chan *Client      // 註冊Chan
	unregisterChan  chan *Client      // 註銷Chan
	kickHandlerChan chan *kickHandler // 踢人Chan

	clients *UserMap // 管理用戶列表

	onlineUserNum     atomic.Int64     // 統計在線用戶數量
	onlineUserConnNum atomic.Int64     // 統計在線總連線數量(含同用戶多個連線)
	onlineUserGauge   prometheus.Gauge // 給prometheus監聽統計的數據
	handshakeTimeout  time.Duration
	writeBufferSize   int
	validate          *validator.Validate
	clientPool        sync.Pool
	Publisher         message.Publisher
	// cache             cache.MsgModel
	// userClient        *rpcclient.UserRpcClient
	// disCov discoveryregistry.SvcDiscoveryRegistry

	Compressor
	Encoder
}

func NewWsManger(in digIn) IWsManager {
	manager := &WsManager{
		wsMaxConnNum:     int64(in.Conf.WsConfig.MaxConnNum),
		writeBufferSize:  in.Conf.WsConfig.WriteBufferSize,
		handshakeTimeout: time.Duration(in.Conf.WsConfig.HandshakeTimeoutSec * time.Now().Second()),
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
		registerChan:    make(chan *Client, 1000),
		unregisterChan:  make(chan *Client, 1000),
		kickHandlerChan: make(chan *kickHandler, 1000),
		onlineUserGauge: in.Prom.OnlineUserGauge,
		validate:        validator.New(),
		clients:         newUserMap(),
		Compressor:      NewGzipCompressor(),
		Encoder:         NewGobEncoder(),
		Publisher:       in.Publisher,
	}
	manager.Run(in.Ctx)
	return manager
}

// NewClient implements LongConnServer.
func (w *WsManager) NewClient(ctx *gin.Context, conn IClientConn, isBackground, isCompress bool, token string) *Client {
	client, _ := w.clientPool.Get().(*Client)
	client.ResetClient(ctx, conn, isBackground, isCompress, token)
	return client
}

// GetUserAllConnects implements LongConnServer.
func (*WsManager) GetUserAllConnects(userID string) ([]*Client, bool) {
	panic("unimplemented")
}

// GetUserAll get all online username
func (w *WsManager) GetUserAll() []string {
	resp := w.clients.GetAll()
	names := make([]string, 0)
	for _, v := range resp {
		names = append(names, v.User.Username+":"+v.UserID)
	}
	return names
}

// GetUserPlatformCons implements LongConnServer.
func (*WsManager) GetUserPlatformCons(userID string, platform int) ([]*Client, bool, bool) {
	panic("unimplemented")
}

// KickUserConn implements LongConnServer.
func (*WsManager) KickUserConn(client *Client) error {
	panic("unimplemented")
}

// Run implements LongConnServer.
func (w *WsManager) Run(ctx *signalctx.Context) error {

	var client *Client

	go func() {
		ctx.RunFunc(func() {
			for {
				select {
				case <-ctx.Done():
					// TODO unregister all client
					return
				case client = <-w.registerChan:
					w.registerClient(client)
				case client = <-w.unregisterChan:
					w.unregisterClient(client)
				case onlineInfo := <-w.kickHandlerChan:
					w.kickClient(onlineInfo)
				}
			}
		})
	}()

	return nil
}

// Register implements IWsManager.
func (w *WsManager) Register(c *Client) error {
	if w.onlineUserConnNum.Load() > w.wsMaxConnNum {
		return errs.WebSocketMaxConnectionsError
	}

	c.wsManager = w // TODO client不保存整個wsManager
	w.registerChan <- c
	go c.ReadMessage()
	return nil
}

// UnRegister implements LongConnServer.
func (w *WsManager) UnRegister(c *Client) {
	w.unregisterChan <- c
}

// Validate implements LongConnServer.
func (*WsManager) Validate(s any) error {
	panic("unimplemented")
}

// wsHandler implements LongConnServer.
func (*WsManager) wsHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (w *WsManager) registerClient(client *Client) {
	old, userOK := w.clients.Get(client.UserID)
	if !userOK { // 用戶完全沒連線
		w.clients.Set(client.UserID, client)
		w.onlineUserNum.Add(1)
		w.onlineUserGauge.Add(1)
	} else { // 該用戶有同平台的連線
		w.multiTerminalLoginChecker(old, client)
	}
	w.onlineUserConnNum.Add(1)

	// 向倉儲註冊用戶資訊
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// TODO 向其他節點更新在線資訊
	}()
	go func() {
		defer wg.Done()
		w.SetUserOnlineStatus(client.ctx, client, enums.Online)
	}()
	wg.Wait()
}

func (w *WsManager) unregisterClient(client *Client) {
	// 放回池子減少記憶體增減
	defer w.clientPool.Put(client)

	isDelete := w.clients.delete(client.UserID, client.ctx.RemoteIP())
	if isDelete { // 該用戶沒有其他在線連線
		w.onlineUserNum.Add(-1)
		w.onlineUserGauge.Add(-1)
	}
	w.onlineUserConnNum.Add(-1)
	w.SetUserOnlineStatus(client.ctx, client, enums.Offline)
	slog.InfoContext(client.ctx, "user offline",
		", close reason: ", client.closedErr,
		", online user Num: ", w.onlineUserNum.Load(),
		", online user conn Num: ", w.onlineUserConnNum.Load(),
	)
}

// multiTerminalLoginChecker 同用戶不同裝置登入是否剔除判斷
func (w *WsManager) multiTerminalLoginChecker(oldClients []*Client, newClient *Client) {
	// 目前不阻擋
	return
}

// kickClient 踢出指定用戶所有連線
func (w *WsManager) kickClient(kick *kickHandler) {
	isDelete := w.clients.deleteClients(kick.newClient.UserID, kick.oldClients)
	if isDelete {
		w.onlineUserNum.Add(-1)
	}
	for _, c := range kick.oldClients {
		err := c.KickOnlineMessage()
		if err != nil {
			slog.WarnContext(c.ctx, "KickOnlineMessage", err)
		}

	}
	// TODO 向倉儲刪除所有用戶連線資訊
	return
}

// SetUserOnlineStatus 更新user在線資訊
func (w *WsManager) SetUserOnlineStatus(ctx *gin.Context, client *Client, status enums.IsOnline) {
	// TODO 向倉儲更新user在線資訊
}

type kickHandler struct {
	clientOK   bool
	oldClients []*Client
	newClient  *Client
}

func (w *WsManager) SendMessage2Client(context context.Context, data *po.Message) error {
	if cs, ok := w.clients.Get(data.RecipientId); ok {
		for _, c := range cs {
			switch data.MsgType {
			case enums.SingleChatType:
				msg, err := json.Marshal(data)
				if err != nil {
					return err
				}
				return c.writeBinaryMsg(msg)
			case enums.GroupChatType:
			case enums.NotificationChatType:
			case enums.SuperGroupChatType:
			}
		}
	}

	return nil
}

func (w *WsManager) SendMessage2Queue(context context.Context, msg *po.Message) error {
	var err error
	queueMsg := &message.Message{}
	queueMsg.UUID = uuid.New()
	queueMsg.Payload, err = msgpack.Marshal(msg)
	if err != nil {
		return err
	}

	return w.Publisher.Publish(topic.MSG, queueMsg)
}
