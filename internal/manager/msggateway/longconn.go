package msggateway

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type PingPongHandler func(string) error

type LongConn interface {
	// Close this connection
	Close() error
	// WriteMessage Write message to connection,messageType means data type,can be set binary(2) and text(1).
	WriteMessage(messageType int, message []byte) error
	// ReadMessage Read message from connection.
	ReadMessage() (int, []byte, error)
	// SetReadDeadline sets the read deadline on the underlying network connection,
	// after a read has timed out, will return an error.
	SetReadDeadline(timeout time.Duration) error
	// SetWriteDeadline sets to write deadline when send message,when read has timed out,will return error.
	SetWriteDeadline(timeout time.Duration) error
	// Dial Try to dial a connection,url must set auth args,header can control compress data
	Dial(urlStr string, requestHeader http.Header) (*http.Response, error)
	// IsNil Whether the connection of the current long connection is nil
	IsNil() bool
	// SetConnNil Set the connection of the current long connection to nil
	SetConnNil()
	// SetReadLimit sets the maximum size for a message read from the peer.bytes
	SetReadLimit(limit int64)
	SetPongHandler(handler PingPongHandler)
	SetPingHandler(handler PingPongHandler)
	// GenerateLongConn Check the connection of the current and when it was sent are the same
	GenerateConnection(w http.ResponseWriter, r *http.Request) error
}
type GWebSocket struct {
	protocolType     int
	conn             *websocket.Conn
	handshakeTimeout time.Duration
	writeBufferSize  int
}

func NewGWebSocket(protocolType int, handshakeTimeout time.Duration, wbs int) *GWebSocket {
	return &GWebSocket{protocolType: protocolType, handshakeTimeout: handshakeTimeout, writeBufferSize: wbs}
}

func (d *GWebSocket) Close() error {
	return d.conn.Close()
}

var mu = sync.Mutex{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (d *GWebSocket) GenerateConnection(w http.ResponseWriter, r *http.Request) error {
	mu.Lock()
	defer mu.Unlock()

	if d.writeBufferSize > 0 {
		upgrader.WriteBufferSize = d.writeBufferSize
	}
	if d.handshakeTimeout > 0 {
		upgrader.HandshakeTimeout = d.handshakeTimeout
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	d.conn = conn
	return nil
}

func (d *GWebSocket) WriteMessage(messageType int, message []byte) error {
	// d.setSendConn(d.conn)
	return d.conn.WriteMessage(messageType, message)
}

//func (d *GWebSocket) setSendConn(sendConn *websocket.Conn) {
//	d.sendConn = sendConn
//}

func (d *GWebSocket) ReadMessage() (int, []byte, error) {
	return d.conn.ReadMessage()
}

func (d *GWebSocket) SetReadDeadline(timeout time.Duration) error {
	return d.conn.SetReadDeadline(time.Now().Add(timeout))
}

func (d *GWebSocket) SetWriteDeadline(timeout time.Duration) error {
	return d.conn.SetWriteDeadline(time.Now().Add(timeout))
}

func (d *GWebSocket) Dial(urlStr string, requestHeader http.Header) (*http.Response, error) {
	conn, httpResp, err := websocket.DefaultDialer.Dial(urlStr, requestHeader)
	if err == nil {
		d.conn = conn
	}
	return httpResp, err
}

func (d *GWebSocket) IsNil() bool {
	if d.conn != nil {
		return false
	}
	return true
}

func (d *GWebSocket) SetConnNil() {
	d.conn = nil
}

func (d *GWebSocket) SetReadLimit(limit int64) {
	d.conn.SetReadLimit(limit)
}

func (d *GWebSocket) SetPongHandler(handler PingPongHandler) {
	d.conn.SetPongHandler(handler)
}

func (d *GWebSocket) SetPingHandler(handler PingPongHandler) {
	d.conn.SetPingHandler(handler)
}

//func (d *GWebSocket) CheckSendConnDiffNow() bool {
//	return d.conn == d.sendConn
//}
