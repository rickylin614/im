package msggateway

import (
	"im/internal/pkg/config"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
)

type MsgGatewayManager interface {
	Run() error
	wsHandler(w http.ResponseWriter, r *http.Request)
	GetUserAllCons(userID string) ([]*Client, bool)
	GetUserPlatformCons(userID string, platform int) ([]*Client, bool, bool)
	Validate(s any) error
	// SetCacheHandler(cache cache.MsgModel)
	// SetDiscoveryRegistry(client discoveryregistry.SvcDiscoveryRegistry)
	KickUserConn(client *Client) error
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
	port              string
	wsMaxConnNum      int64
	registerChan      chan *Client
	unregisterChan    chan *Client
	kickHandlerChan   chan *kickHandler
	clients           *sync.Map
	clientPool        sync.Pool
	onlineUserNum     atomic.Int64
	onlineUserConnNum atomic.Int64
	handshakeTimeout  time.Duration
	writeBufferSize   int
	validate          *validator.Validate
	// cache             cache.MsgModel
	// userClient        *rpcclient.UserRpcClient
	// disCov discoveryregistry.SvcDiscoveryRegistry
	Compressor
	Encoder
	MessageHandler
}

func NewWsManger(in digIn) MsgGatewayManager {
	manager := &WsManager{
		port:             in.conf.WsConfig.Port,
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
func (*WsManager) Run() error {
	// TODO 啟動基本監聽相關JOB
	panic("unimplemented")
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
