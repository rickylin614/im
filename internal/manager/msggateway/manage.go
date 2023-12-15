package msggateway

import (
	"fmt"
	"im/internal/pkg/config"
	"im/internal/util/errs"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
	"golang.org/x/sync/errgroup"
)

type LongConnPoolMgmt interface {
	Run() error
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

	conf config.Config
}

type WsManager struct {
	wsMaxConnNum      int64             // 設定檔限制的最大連線數量
	registerChan      chan *Client      // 註冊Chan
	unregisterChan    chan *Client      // 註銷Chan
	kickHandlerChan   chan *kickHandler // 踢人Chan
	clients           *sync.Map
	onlineUserNum     atomic.Int64
	onlineUserConnNum atomic.Int64
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
		wsMaxConnNum:     int64(in.conf.WsConfig.MaxConnNum),
		writeBufferSize:  in.conf.WsConfig.WriteBufferSize,
		handshakeTimeout: time.Duration(in.conf.WsConfig.HandshakeTimeoutSec * time.Now().Second()),
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
		registerChan:    make(chan *Client, 1000),
		unregisterChan:  make(chan *Client, 1000),
		kickHandlerChan: make(chan *kickHandler, 1000),
		validate:        validator.New(),
		// clients:         newUserMap(),
		Compressor: NewGzipCompressor(),
		Encoder:    NewGobEncoder(),
	}
	manager.Run()
	return manager
}

// GetUserAllCons implements LongConnServer.
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
func (ws *WsManager) Run() error {
	var (
		client *Client
		wg     errgroup.Group

		sigs = make(chan os.Signal, 1)
		done = make(chan struct{}, 1)
	)

	wg.Go(func() error {
		for {
			select {
			case <-done:
				return nil

			case client = <-ws.registerChan:
				fmt.Println(client)
				// ws.registerClient(client)
			case client = <-ws.unregisterChan:
				fmt.Println(client)
				// ws.unregisterClient(client)
			case onlineInfo := <-ws.kickHandlerChan:
				fmt.Println(onlineInfo)
				// ws.multiTerminalLoginChecker(onlineInfo.clientOK, onlineInfo.oldClients, onlineInfo.newClient)
			}
		}
	})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs
	close(done)
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

type kickHandler struct {
	clientOK   bool
	oldClients []*Client
	newClient  *Client
}
