package handler

import (
	"im/internal/models/request"
	"im/internal/models/response"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type messageHandler struct {
	in webDigIn
}

// Get
// @Summary Get
// @Tags message
// @Param request query request.MessageGet true "param"
// @Success 200 {object} response.APIResponse[response.MessageGet]
// @Router /message/:id [get]
func (h messageHandler) Get(ctx *gin.Context) {
	req := &request.MessageGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.MessageSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.MessageGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags message
// @Param request query request.MessageGetList true "param"
// @Success 200 {object} response.APIResponse[response.MessageGetList]
// @Router /message [get]
func (h messageHandler) GetList(ctx *gin.Context) {
	req := &request.MessageGetList{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.MessageSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.MessageGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Update
// @Summary Update
// @Tags message
// @Param request body request.MessageUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /message [put]
func (h messageHandler) Update(ctx *gin.Context) {
	req := &request.MessageUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.MessageSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags message
// @Param request body request.MessageDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /message [delete]
func (h messageHandler) Delete(ctx *gin.Context) {
	req := &request.MessageDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.MessageSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
