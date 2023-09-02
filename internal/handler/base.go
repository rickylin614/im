package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type baseHandler struct {
	in digIn
}

// Ping
// @Summary Health check
// @Tags public
// @Success 200 {string} string
// @Router /ping [get]
func (b baseHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Metrics
// @Summary Metrics
// @Tags public
// @Success 200 {string} string
// @Router /metrics [get]
func (b baseHandler) Metrics() gin.HandlerFunc {
	promHandler := promhttp.Handler()
	return func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	}
}
