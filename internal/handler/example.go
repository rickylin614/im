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

// @Summary Get
// @Tags example
// @Param body body request.ExampleGet true "param"
// @Success 200 {object} response.APIResponse[response.ExampleGet]
// @Router /example/:id [get]
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

// @Summary GetList
// @Tags example
// @Param body body request.ExampleGetList true "param"
// @Success 200 {object} response.APIResponse[response.ExampleGetList]
// @Router /example [get]
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

// @Summary Create
// @Tags example
// @Param body body request.ExampleCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /example [post]
func (h ExampleHandler) Create(ctx *gin.Context) {
	req := &request.ExampleCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.ExampleSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// @Summary Update
// @Tags example
// @Param body body request.ExampleUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /example [put]
func (h ExampleHandler) Update(ctx *gin.Context) {
	req := &request.ExampleUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.ExampleSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// @Summary Delete
// @Tags example
// @Param body body request.ExampleDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /example [delete]
func (h ExampleHandler) Delete(ctx *gin.Context) {
	req := &request.ExampleDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.ExampleSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
