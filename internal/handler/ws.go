package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"

	"im/internal/manager/msggateway"
	"im/internal/util/ctxs"
)

type wsHandler struct {
	in wsDigIn
}

// Connect establishes a WebSocket connection
// @Summary Establish WebSocket connection
// @Tags ws
// @Success 101 {object} object "WebSocket Protocol Handshake"
// @Failure 400 {object} object "Invalid request format"
// @Router /connect [get]
func (h wsHandler) Connect(ctx *gin.Context) {
	// check login
	if err := ctxs.CheckLoginByParam(ctx, h.in.Service.UsersSrv); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	// 創建長連線物件
	gSocket := msggateway.NewClientConn(1,
		time.Duration(h.in.Config.WsConfig.HandshakeTimeoutSec*int(time.Second)),
		h.in.Config.WsConfig.WriteBufferSize)

	// 升級協定 upgrade http to ws
	err := gSocket.GenerateConnection(ctx.Writer, ctx.Request)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	// 會員資料 member info
	client := h.in.WsManager.NewClient(ctx, gSocket, true, false, "")

	// 註冊 (create message listener.go here)
	err = h.in.WsManager.Register(client)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
}

func (h wsHandler) WsTest(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error("upgrade fail", err)
		return
	}
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error(string(msg), err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, append(msg, " - from server"...))
		if err != nil {
			slog.Error(string(msg), err)
			break
		}
	}

}
