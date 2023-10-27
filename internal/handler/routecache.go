package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type routeCacheHandler struct {
	in digIn
}

// Get
// @Summary Get
// @Tags routeCache
// @Param body body request.RouteCacheGet true "param"
// @Success 200 {object} response.APIResponse[response.RouteCacheGet]
// @Router /routeCache/:id [get]
func (h routeCacheHandler) Get(ctx *gin.Context) {
	req := &request.RouteCacheGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.RouteCacheSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.RouteCacheGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags routeCache
// @Param body body request.RouteCacheGetList true "param"
// @Success 200 {object} response.APIResponse[response.RouteCacheGetList]
// @Router /routeCache [get]
func (h routeCacheHandler) GetList(ctx *gin.Context) {
	req := &request.RouteCacheGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.RouteCacheSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.RouteCacheGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary Create
// @Tags routeCache
// @Param body body request.RouteCacheCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /routeCache [post]
func (h routeCacheHandler) Create(ctx *gin.Context) {
	req := &request.RouteCacheCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.RouteCacheSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary Update
// @Tags routeCache
// @Param body body request.RouteCacheUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /routeCache [put]
func (h routeCacheHandler) Update(ctx *gin.Context) {
	req := &request.RouteCacheUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.RouteCacheSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags routeCache
// @Param body body request.RouteCacheDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /routeCache [delete]
func (h routeCacheHandler) Delete(ctx *gin.Context) {
	req := &request.RouteCacheDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.RouteCacheSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
