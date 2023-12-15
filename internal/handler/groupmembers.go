package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type groupMembersHandler struct {
	in webDigIn
}

// GetList
// @Summary GetList
// @Tags groupMembers
// @Param request query request.GroupMembersGetList true "param"
// @Success 200 {object} response.APIResponse[response.GroupMembersGetList.Data]
// @Router /group-members/:id [get]
func (h groupMembersHandler) GetList(ctx *gin.Context) {
	req := &request.GroupMembersGetList{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.GroupMembersSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.GroupMembersGetList{Data: make([]response.GroupMembersGet, 0)}
	if err := copier.Copy(result.Data, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result.Data)
}

// Create
// @Summary Create
// @Tags groupMembers
// @Param request body request.GroupMembersCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-members [post]
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
// @Param request body request.GroupMembersUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-members [put]
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
// @Param request body request.GroupMembersDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /group-members [delete]
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
