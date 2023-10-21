package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type groupsHandler struct {
	in digIn
}

// Get
// @Summary Get
// @Tags groups
// @Param body body request.GroupsGet true "param"
// @Success 200 {object} response.APIResponse[response.GroupsGet]
// @Router /groups/:id [get]
func (h groupsHandler) Get(ctx *gin.Context) {
	req := &request.GroupsGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupsSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.GroupsGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags groups
// @Param body body request.GroupsGetList true "param"
// @Success 200 {object} response.APIResponse[response.GroupsGetList]
// @Router /groups [get]
func (h groupsHandler) GetList(ctx *gin.Context) {
	req := &request.GroupsGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupsSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.GroupsGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary Create
// @Tags groups
// @Param body body request.GroupsCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groups [post]
func (h groupsHandler) Create(ctx *gin.Context) {
	req := &request.GroupsCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.GroupsSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary Update
// @Tags groups
// @Param body body request.GroupsUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groups [put]
func (h groupsHandler) Update(ctx *gin.Context) {
	req := &request.GroupsUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.GroupsSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags groups
// @Param body body request.GroupsDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groups [delete]
func (h groupsHandler) Delete(ctx *gin.Context) {
	req := &request.GroupsDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.GroupsSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
