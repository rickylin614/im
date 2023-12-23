package handler

import (
	"im/internal/manager/msggateway"
	"im/internal/util/ctxs"
	"time"

	"github.com/gin-gonic/gin"
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
	longConn := msggateway.NewGWebSocket(1,
		time.Duration(h.in.Config.WsConfig.HandshakeTimeoutSec*int(time.Second)),
		h.in.Config.WsConfig.WriteBufferSize)

	// 升級協定
	err := longConn.GenerateLongConn(ctx.Writer, ctx.Request)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	// 會員資料
	client := h.in.WsManager.NewClient(ctx, longConn, true, false, "")

	// 註冊
	err = h.in.WsManager.Register(client)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
}
