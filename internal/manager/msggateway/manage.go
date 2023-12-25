package msggateway

import (
	"log/slog"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"im/internal/pkg/config"
	"im/internal/pkg/consts/enums"
	"im/internal/pkg/signalctx"
	"im/internal/util/errs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
)

type LongConnPoolMgmt interface {
	Run(ctx *signalctx.Context) error
	// wsHandler(w http.ResponseWriter, r *http.Request)
	NewClient(ctx *gin.Context, conn LongConn, isBackground, isCompress bool, token string) *Client
	GetUserAllCons(userID string) ([]*Client, bool)
	GetUserPlatformCons(userID string, platform int) ([]*Client, bool, bool)
	Validate(s any) error
	// SetCacheHandler(cache cache.MsgModel)
	// SetDiscoveryRegistry(client discoveryregistry.SvcDiscoveryRegistry)
	KickUserConn(client *Client) error
	Register(c *Client) error
	UnRegister(c *Client)
	// SetKickHandlerInfo(i *kickHandler)
	Compressor
	Encoder
	MessageHandler
}

type digIn struct {
	dig.In

	Conf config.Config
	Ctx  *signalctx.Context
}

type WsManager struct {
	wsMaxConnNum      int64             // 設定檔限制的最大連線數量
	registerChan      chan *Client      // 註冊Chan
	unregisterChan    chan *Client      // 註銷Chan
	kickHandlerChan   chan *kickHandler // 踢人Chan
	clients           *UserMap
	onlineUserNum     atomic.Int64     // 統計在線用戶數量
	onlineUserConnNum atomic.Int64     // 統計在線總連線數量(含同用戶多個連線)
	onlineUserGauge   prometheus.Gauge // 給prometheus監聽統計的數據
	handshakeTimeout  time.Duration
	writeBufferSize   int
	validate          *validator.Validate
	clientPool        sync.Pool
	// cache             cache.MsgModel
	// userClient        *rpcclient.UserRpcClient
	// disCov discoveryregistry.SvcDiscoveryRegistry

	Compressor
	Encoder
	MessageHandler
}

func NewWsManger(in digIn) LongConnPoolMgmt {
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
		onlineUserGauge: prometheus.NewGauge(prometheus.GaugeOpts{ // TODO 另外創建一個集中管理的結構體
			Name: "online_user_num",
			Help: "The number of online user num",
		}),
		validate:   validator.New(),
		clients:    newUserMap(),
		Compressor: NewGzipCompressor(),
		Encoder:    NewGobEncoder(),
	}
	manager.Run(in.Ctx)
	return manager
}

// NewClient implements LongConnServer.
func (w *WsManager) NewClient(ctx *gin.Context, conn LongConn, isBackground, isCompress bool, token string) *Client {
	client, _ := w.clientPool.Get().(*Client)
	client.ResetClient(ctx, conn, isBackground, isCompress, token)
	return client
}

// GetUserAllCons implements LongConnServer.
func (*WsManager) GetUserAllCons(userID string) ([]*Client, bool) {
	panic("unimplemented")
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
func (ws *WsManager) Run(ctx *signalctx.Context) error {
	ctx.Increment()
	defer ctx.Decrement()

	var client *Client

	go func() {
		for {
			select {
			case <-ctx.Done():
				// TODO unregister all client
				return
			case client = <-ws.registerChan:
				ws.registerClient(client)
			case client = <-ws.unregisterChan:
				ws.unregisterClient(client)
			case onlineInfo := <-ws.kickHandlerChan:
				ws.kickClient(onlineInfo)
			}
		}
	}()

	return nil
}

// Register implements LongConnPoolMgmt.
func (w *WsManager) Register(c *Client) error {
	if w.onlineUserConnNum.Load() > w.wsMaxConnNum {
		return errs.WebSocketMaxConnectionsError
	}

	c.longConnManager = w // TODO client不保存整個wsManager
	w.registerChan <- c
	go c.ReadMessage()
	return nil
}

// UnRegister implements LongConnServer.
func (*WsManager) UnRegister(c *Client) {
	panic("unimplemented")
}

// Validate implements LongConnServer.
func (*WsManager) Validate(s any) error {
	panic("unimplemented")
}

// wsHandler implements LongConnServer.
func (*WsManager) wsHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (ws *WsManager) registerClient(client *Client) {
	old, userOK := ws.clients.Get(client.UserID)
	if !userOK { // 用戶完全沒連線
		ws.clients.Set(client.UserID, client)
		ws.onlineUserNum.Add(1)
		ws.onlineUserGauge.Add(1)
	} else { // 該用戶有同平台的連線
		ws.multiTerminalLoginChecker(old, client)
	}
	ws.onlineUserConnNum.Add(1)

	// 向倉儲註冊用戶資訊
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// TODO 向其他節點更新在線資訊
	}()
	go func() {
		defer wg.Done()
		ws.SetUserOnlineStatus(client.ctx, client, enums.Online)
	}()
	wg.Wait()
}

func (ws *WsManager) unregisterClient(client *Client) {
	// 放回池子減少記憶體增減
	defer ws.clientPool.Put(client)

	isDelete := ws.clients.delete(client.UserID, client.ctx.RemoteIP())
	if isDelete { // 該用戶沒有其他在線連線
		ws.onlineUserNum.Add(-1)
		ws.onlineUserGauge.Add(-1)
	}
	ws.onlineUserConnNum.Add(-1)
	ws.SetUserOnlineStatus(client.ctx, client, enums.Offline)
	slog.InfoContext(client.ctx, "user offline",
		", close reason: ", client.closedErr,
		", online user Num: ", ws.onlineUserNum.Load(),
		", online user conn Num: ", ws.onlineUserConnNum.Load(),
	)
}

// multiTerminalLoginChecker 同用戶不同裝置登入是否剔除判斷
func (ws *WsManager) multiTerminalLoginChecker(oldClients []*Client, newClient *Client) {
	// 目前不阻擋
	return
}

// kickClient 踢出指定用戶所有連線
func (ws *WsManager) kickClient(kick *kickHandler) {
	isDelete := ws.clients.deleteClients(kick.newClient.UserID, kick.oldClients)
	if isDelete {
		ws.onlineUserNum.Add(-1)
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
func (ws *WsManager) SetUserOnlineStatus(ctx *gin.Context, client *Client, status enums.IsOnline) {
	// TODO 向倉儲更新user在線資訊
}

type kickHandler struct {
	clientOK   bool
	oldClients []*Client
	newClient  *Client
}
