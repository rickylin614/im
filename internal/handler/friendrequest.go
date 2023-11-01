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
// @Summary 好友請求列表
// @Tags FriendRequests
// @Param body body request.FriendRequestsGetList true "param"
// @Success 200 {object} response.APIResponse[response.FriendRequestsGetList]
// @Router /friend-requests [get]
func (h FriendRequestsHandler) GetList(ctx *gin.Context) {
	req := &request.FriendRequestsGetList{UserId: ctxs.GetUserInfo(ctx).ID}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.FriendRequestSrv.GetList(ctx, req)
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
// @Failure 400 {object} response.APIResponse[string] "无效的用户"
// @Failure 401 {object} response.APIResponse[string] "Unauthorized"
// @Failure 409 {object} response.APIResponse[string] "Friend request already exists between these users."
// @Failure 500 {object} response.APIResponse[string] "未知错误"
// @Router /friend-requests [post]
func (h FriendRequestsHandler) Create(ctx *gin.Context) {
	req := &request.FriendRequestsCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.FriendRequestSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary 接受或拒絕來自id的好友請求
// @Tags FriendRequests
// @Param body body request.FriendRequestsUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Failure 400 {object} response.APIResponse[string] "无效的ID"
// @Failure 401 {object} response.APIResponse[string] "Unauthorized"
// @Failure 500 {object} response.APIResponse[string] "未知错误"
// @Router /friend-requests [put]
func (h FriendRequestsHandler) Update(ctx *gin.Context) {
	req := &request.FriendRequestsUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.FriendRequestSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
