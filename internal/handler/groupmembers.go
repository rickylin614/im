package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type groupMembersHandler struct {
	in digIn
}

// Get
// @Summary Get
// @Tags groupMembers
// @Param body body request.GroupMembersGet true "param"
// @Success 200 {object} response.APIResponse[response.GroupMembersGet]
// @Router /groupMembers/:id [get]
func (h groupMembersHandler) Get(ctx *gin.Context) {
	req := &request.GroupMembersGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupMembersSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.GroupMembersGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags groupMembers
// @Param body body request.GroupMembersGetList true "param"
// @Success 200 {object} response.APIResponse[response.GroupMembersGetList]
// @Router /groupMembers [get]
func (h groupMembersHandler) GetList(ctx *gin.Context) {
	req := &request.GroupMembersGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupMembersSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.GroupMembersGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary Create
// @Tags groupMembers
// @Param body body request.GroupMembersCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groupMembers [post]
func (h groupMembersHandler) Create(ctx *gin.Context) {
	req := &request.GroupMembersCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.GroupMembersSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary Update
// @Tags groupMembers
// @Param body body request.GroupMembersUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groupMembers [put]
func (h groupMembersHandler) Update(ctx *gin.Context) {
	req := &request.GroupMembersUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.GroupMembersSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags groupMembers
// @Param body body request.GroupMembersDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /groupMembers [delete]
func (h groupMembersHandler) Delete(ctx *gin.Context) {
	req := &request.GroupMembersDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.GroupMembersSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
