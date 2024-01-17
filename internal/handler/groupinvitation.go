package handler

import (
	"im/internal/models/request"
	"im/internal/models/response"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type groupInvitationHandler struct {
	in webDigIn
}

// Get
// @Summary Get
// @Tags groupInvitation
// @Param request body request.GroupInvitationGet true "param"
// @Success 200 {object} response.APIResponse[response.GroupInvitationGet]
// @Router /group-invitation/:id [get]
func (h groupInvitationHandler) Get(ctx *gin.Context) {
	req := &request.GroupInvitationGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupInvitationSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.GroupInvitationGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags groupInvitation
// @Param request body request.GroupInvitationGetList true "param"
// @Success 200 {object} response.APIResponse[response.GroupInvitationGetList]
// @Router /group-invitation [get]
func (h groupInvitationHandler) GetList(ctx *gin.Context) {
	req := &request.GroupInvitationGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupInvitationSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.GroupInvitationGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary 新建邀請碼
// @Tags groupInvitation
// @Param request body request.GroupInvitationCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-invitation [post]
func (h groupInvitationHandler) Create(ctx *gin.Context) {
	req := &request.GroupInvitationCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.GroupInvitationSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary Update
// @Tags groupInvitation
// @Param request body request.GroupInvitationUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-invitation [put]
func (h groupInvitationHandler) Update(ctx *gin.Context) {
	req := &request.GroupInvitationUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.GroupInvitationSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags groupInvitation
// @Param request body request.GroupInvitationDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-invitation [delete]
func (h groupInvitationHandler) Delete(ctx *gin.Context) {
	req := &request.GroupInvitationDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.GroupInvitationSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
