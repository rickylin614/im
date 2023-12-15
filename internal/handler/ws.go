package handler

import (
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
	// h.in.WsManager.
}
