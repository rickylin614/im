package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type baseHandler struct {
	in digIn
}

// Ping Healthcheck
//
// param: ctx is gin param
func (b baseHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (b baseHandler) Metrics() gin.HandlerFunc {
	promHandler := promhttp.Handler()
	return func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	}
}
