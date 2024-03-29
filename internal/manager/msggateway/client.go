package msggateway

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"

	"im/internal/models/po"
	"im/internal/util/ctxs"
	"im/internal/util/uuid"
)

var (
	ErrConnClosed                = errors.New("conn has closed")
	ErrNotSupportMessageProtocol = errors.New("not support message protocol")
	ErrClientClosed              = errors.New("client actively close the connection")
	ErrPanic                     = errors.New("panic error")
)

const (
	// MessageText is for UTF-8 encoded text messages like JSON.
	MessageText = iota + 1
	// MessageBinary is for binary messages like protobufs.
	MessageBinary
	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

type Client struct {
	w    *sync.Mutex
	conn IClientConn
	// PlatformID     int    `json:"platformID"`
	IsCompress   bool      `json:"isCompress"` // 訊息是否要壓縮、目前寫死 false 在建立連線的時候
	UserID       string    `json:"userID"`     // 用戶ID
	User         *po.Users `json:"user"`       // 用戶結構體
	IsBackground bool      `json:"isBackground"`
	ctx          *gin.Context
	closed       atomic.Bool
	closedErr    error
	token        string

	// 紀錄管理器位址供呼叫
	wsManager IWsManager
}

func newClient(ctx *gin.Context, conn IClientConn, isCompress bool) *Client {
	return &Client{
		w:          new(sync.Mutex),
		conn:       conn,
		IsCompress: isCompress,
		UserID:     ctxs.GetUserInfo(ctx).ID,
		User:       ctxs.GetUserInfo(ctx),
		ctx:        ctx,
	}
}

func (c *Client) ResetClient(
	ctx *gin.Context,
	conn IClientConn,
	isBackground, isCompress bool,
	token string,
) {
	c.w = new(sync.Mutex)
	c.conn = conn
	c.IsCompress = isCompress
	c.IsBackground = isBackground
	c.UserID = ctxs.GetUserInfo(ctx).ID
	c.User = ctxs.GetUserInfo(ctx)
	c.ctx = ctx
	c.closed.Store(false)
	c.closedErr = nil
	c.token = token
}

func (c *Client) pingHandler(_ string) error {
	_ = c.conn.SetReadDeadline(pongWait)
	return c.writePongMsg()
}

func (c *Client) ReadMessage() {
	defer func() {
		if r := recover(); r != nil {
			c.closedErr = ErrPanic
			fmt.Println("socket have panic err:", r, string(debug.Stack()))
		}
		c.close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	//_ = c.conn.SetReadDeadline(pongWait)  // close when debug mode
	c.conn.SetPingHandler(c.pingHandler)

	count := 1

	for {
		messageType, message, returnErr := c.conn.ReadMessage()
		if returnErr != nil {
			slog.WarnContext(c.ctx, "readMessage", "err", returnErr, "messageType", messageType)
			c.closedErr = returnErr
			return
		}

		slog.DebugContext(c.ctx, "readMessage", "err", returnErr, "messageType", messageType)
		if c.closed.Load() { // 连接刚置位已经关闭，但是协程还没退出的场景
			c.closedErr = ErrConnClosed
			return
		}

		switch messageType {
		case MessageBinary:
			_ = c.conn.SetReadDeadline(pongWait)
			parseDataErr := c.handleMessage(message)
			if parseDataErr != nil {
				c.closedErr = parseDataErr
				return
			}
		case MessageText:
			parseDataErr := c.handleMessage(message)
			testMsg1 := &po.Message{ID: uuid.New()}
			testMsg := message
			if err := json.Unmarshal(message, testMsg1); err == nil {
				testMsg = []byte(testMsg1.MsgContent)
			}
			testMsg2 := &po.Message{
				Sender:     c.User.Username,
				MsgContent: fmt.Sprintf("第%d次溝通 收到訊息:%s", count, testMsg),
				CreatedAt:  time.Now(),
			}
			jsonByte, err := json.Marshal(testMsg2)
			c.writeBinaryMsg(jsonByte)

			count++
			slog.Debug("ws recieve", "msg", string(message))
			if parseDataErr != nil {
				c.closedErr = parseDataErr
				return
			}

			poMsg := &po.Message{}
			err = json.Unmarshal(message, poMsg)
			if err != nil {
				slog.Error("ws recieve json.Unmarshal fail", "err", err.Error(), "user", c.User.Username, "data", string(message))
				continue
			}
			// 確保來源資訊沒被竄改
			poMsg.ID = uuid.New()
			poMsg.Sender = c.User.Username
			poMsg.SenderId = c.UserID
			poMsg.CreatedAt, poMsg.UpdatedAt = time.Now(), time.Now()

			err = c.wsManager.SendMessage2Queue(context.Background(), poMsg)
			if err != nil {
				slog.Error("ws recieve SendMessage2Queue fail", "err", err.Error(), "user", c.User.Username, "msg", poMsg)
				continue
			}

		case PingMessage:
			err := c.writePongMsg()
			if err != nil {
				slog.ErrorContext(c.ctx, "writePongMsg", err)
			}

		case CloseMessage:
			c.closedErr = ErrClientClosed
			return
		default:
		}
	}
}

func (c *Client) handleMessage(message []byte) error {
	if c.IsCompress {
		var err error
		message, err = c.wsManager.DecompressWithPool(message)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (c *Client) handleStrMessage(message string) error {
	return c.handleMessage([]byte(message))
}

func (c *Client) setAppBackgroundStatus(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (c *Client) close() {
	if c.closed.Load() {
		return
	}

	c.w.Lock()
	defer c.w.Unlock()

	c.closed.Store(true)
	c.conn.Close()
	c.wsManager.UnRegister(c)
}

func (c *Client) replyMessage(ctx context.Context) error {
	return nil
}

func (c *Client) PushMessage(ctx context.Context) error {
	return nil
}

func (c *Client) KickOnlineMessage() error {
	// TODO 踢出訊息

	// 關閉長連線
	c.close()
	return nil
}

func (c *Client) writeBinaryMsg(b []byte) error {
	if c.closed.Load() {
		return nil
	}

	c.w.Lock()
	defer c.w.Unlock()

	return c.conn.WriteMessage(MessageBinary, b)
}

func (c *Client) writeStringMsg(s string) error {
	if c.closed.Load() {
		return nil
	}

	c.w.Lock()
	defer c.w.Unlock()

	return c.conn.WriteMessage(MessageText, []byte(s))
}

func (c *Client) writePongMsg() error {
	if c.closed.Load() {
		return nil
	}

	c.w.Lock()
	defer c.w.Unlock()

	err := c.conn.SetWriteDeadline(writeWait)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return c.conn.WriteMessage(PongMessage, nil)
}
