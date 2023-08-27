package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/pkg/ctxs"
	"im/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ExampleHandler struct {
	in digIn

	Service *service.Service
}

func (h ExampleHandler) Get(ctx *gin.Context) {
	req := &request.ExampleGet{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.ExampleSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.ExampleGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

func (h ExampleHandler) GetList(ctx *gin.Context) {
	req := &request.ExampleGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.ExampleSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.ExampleGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}
