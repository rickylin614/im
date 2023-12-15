package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type friendHandler struct {
	in webDigIn
}

// GetFriends
// @Summary 獲取用戶的好友列表
// @Tags friend
// @Param request body request.FriendGetList true "param"
// @Success 200 {object} response.APIResponse[response.FriendGetList]
// @Router /friend [get]
func (h friendHandler) GetFriends(ctx *gin.Context) {
	req := &request.FriendGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.FriendSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.FriendGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// UpdateFriendStatus
// @Summary 更新與指定用戶的好友關係（接受/拒絕/阻止）
// @Tags friend
// @Param request body request.FriendUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend [put]
func (h friendHandler) UpdateFriendStatus(ctx *gin.Context) {
	req := &request.FriendUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.FriendSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// DeleteFriend
// @Summary 刪除與指定用戶的好友關係
// @Tags friend
// @Param request body request.FriendDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend [delete]
func (h friendHandler) DeleteFriend(ctx *gin.Context) {
	req := &request.FriendDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.FriendSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// GetBlockedFriends
// @Summary 獲取指定用戶ID的已封鎖好友列表
// @Tags friend
// @Param request body request.FriendGetList true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend/blocked [get]
func (h friendHandler) GetBlockedFriends(ctx *gin.Context) {
	req := &request.FriendGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.FriendSrv.GetBlackList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.FriendGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetMutualFriends
// @Summary 獲取指定用戶ID與另一指定用戶ID的共同好友列表
// @Tags friend
// @Param request body request.FriendMutualGet true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /friend/mutual [get]
func (h friendHandler) GetMutualFriends(ctx *gin.Context) {
	req := &request.FriendMutualGet{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.FriendSrv.GetMutualList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.FriendMutualList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}
