package handler

import (
	request "im/internal/models/req"
	"im/internal/service"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	in digIn

	Service *service.Service
}

func (h ExampleHandler) ExampleGet(ctx *gin.Context) {
	req := &request.ExampleGet{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		// TODO Err set
		return
	}
	// if h.in.Service

}
