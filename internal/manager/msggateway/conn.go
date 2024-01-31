package msggateway

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type PingPongHandler func(string) error

type IClientConn interface {
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
type ClientConn struct {
	protocolType     int
	conn             *websocket.Conn
	handshakeTimeout time.Duration
	writeBufferSize  int
}

func NewClientConn(protocolType int, handshakeTimeout time.Duration, wbs int) IClientConn {
	return &ClientConn{protocolType: protocolType, handshakeTimeout: handshakeTimeout, writeBufferSize: wbs}
}

func (d *ClientConn) Close() error {
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

func (d *ClientConn) GenerateConnection(w http.ResponseWriter, r *http.Request) error {
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

func (d *ClientConn) WriteMessage(messageType int, message []byte) error {
	// d.setSendConn(d.conn)
	return d.conn.WriteMessage(messageType, message)
}

//func (d *ClientConn) setSendConn(sendConn *websocket.Conn) {
//	d.sendConn = sendConn
//}

func (d *ClientConn) ReadMessage() (int, []byte, error) {
	return d.conn.ReadMessage()
}

func (d *ClientConn) SetReadDeadline(timeout time.Duration) error {
	return d.conn.SetReadDeadline(time.Now().Add(timeout))
}

func (d *ClientConn) SetWriteDeadline(timeout time.Duration) error {
	return d.conn.SetWriteDeadline(time.Now().Add(timeout))
}

func (d *ClientConn) Dial(urlStr string, requestHeader http.Header) (*http.Response, error) {
	conn, httpResp, err := websocket.DefaultDialer.Dial(urlStr, requestHeader)
	if err == nil {
		d.conn = conn
	}
	return httpResp, err
}

func (d *ClientConn) IsNil() bool {
	if d.conn != nil {
		return false
	}
	return true
}

func (d *ClientConn) SetConnNil() {
	d.conn = nil
}

func (d *ClientConn) SetReadLimit(limit int64) {
	d.conn.SetReadLimit(limit)
}

func (d *ClientConn) SetPongHandler(handler PingPongHandler) {
	d.conn.SetPongHandler(handler)
}

func (d *ClientConn) SetPingHandler(handler PingPongHandler) {
	d.conn.SetPingHandler(handler)
}

//func (d *ClientConn) CheckSendConnDiffNow() bool {
//	return d.conn == d.sendConn
//}
