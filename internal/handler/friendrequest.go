package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type FriendRequestsHandler struct {
	in digIn
}

// GetList
// @Summary 獲取指定用戶ID收到的好友請求列表
// @Tags FriendRequests
// @Param body body request.FriendRequestsGetList true "param"
// @Success 200 {object} response.APIResponse[response.FriendRequestsGetList]
// @Router /friend-requests [get]
func (h FriendRequestsHandler) GetList(ctx *gin.Context) {
	req := &request.FriendRequestsGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.FriendRequestsrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.FriendRequestsGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary 向指定用戶ID發送好友請求
// @Tags FriendRequests
// @Param body body request.FriendRequestsCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend-requests [post]
func (h FriendRequestsHandler) Create(ctx *gin.Context) {
	req := &request.FriendRequestsCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.FriendRequestsrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary 指定用戶ID接受或拒絕來自requester-id的好友請求
// @Tags FriendRequests
// @Param body body request.FriendRequestsUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend-requests [put]
func (h FriendRequestsHandler) Update(ctx *gin.Context) {
	req := &request.FriendRequestsUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.FriendRequestsrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary 指定用戶ID刪除來自requester-id的好友請求
// @Tags FriendRequests
// @Param body body request.FriendRequestsDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend-requests [delete]
func (h FriendRequestsHandler) Delete(ctx *gin.Context) {
	req := &request.FriendRequestsDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.FriendRequestsrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
