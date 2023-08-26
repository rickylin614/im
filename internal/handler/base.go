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
// @Description check server exist
// @Tags member
// @ID get-members
// @Accept  json
// @Produce  json
// @Param body body dto.QueryMemberCond true "request param"
// @Success 200 {object} dto.StandardResponse[[]dto.Member]
// @Security ApiKeyAuth
// @Router /member/list [get]
func (b baseHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (b baseHandler) Metrics() gin.HandlerFunc {
	promHandler := promhttp.Handler()
	return func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	}
}
