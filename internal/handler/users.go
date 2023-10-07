package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"
	"im/internal/util/errs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type usersHandler struct {
	in digIn
}

// Login
// @Summary 用戶登錄並返回授權令牌
// @Tags users
// @Param body body request.UsersLogin true "param"
// @Success 200 {object} response.APIResponse[response.UsersLogin]
// @Router /users/login [post]
func (h usersHandler) Login(ctx *gin.Context) {
	req := &request.UsersLogin{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	token, err := h.in.Service.UsersSrv.Login(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, response.UsersLogin{
		Token:    token,
		Username: req.Username,
	})
}

// Logout
// @Summary 用戶登出
// @Tags users
// @Success 200 {object} response.APIResponse[string]
// @Security ApiKeyAuth
// @Router /users/logout [post]
func (h usersHandler) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if len(token) == 0 {
		ctxs.SetError(ctx, errs.RequestTokenError)
	}
	h.in.Service.UsersSrv.Logout(ctx, token)
	ctxs.SetSuccessResp(ctx)
}

// Get
// @Summary 取得用戶訊息
// @Tags users
// @Param body body request.UsersGet true "param"
// @Success 200 {object} response.APIResponse[response.UsersGet]
// @Router /users/:id [get]
func (h usersHandler) Get(ctx *gin.Context) {
	req := &request.UsersGet{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.UsersSrv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.UsersGet{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary 用戶清單
// @Tags users
// @Param body body request.UsersGetList true "param"
// @Success 200 {object} response.APIResponse[response.UsersGetList]
// @Router /users/search [get]
func (h usersHandler) GetList(ctx *gin.Context) {
	req := &request.UsersGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.UsersSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.UsersGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary 用戶註冊
// @Tags users
// @Param body body request.UsersCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /users/register [post]
func (h usersHandler) Create(ctx *gin.Context) {
	req := &request.UsersCreate{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.UsersSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary 用戶訊息修改
// @Tags users
// @Param body body request.UsersUpdate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /users [put]
func (h usersHandler) Update(ctx *gin.Context) {
	req := &request.UsersUpdate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.UsersSrv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// GetOnlineStatus
// @Summary 獲取指定用戶ID的在線狀態
// @Tags users
// @Router /users/{id}/online-status [get]
func (h usersHandler) GetOnlineStatus(ctx *gin.Context) {
	// TODO
}

// UpdateOnlineStatus
// @Summary 更新指定用戶ID的在線狀態
// @Tags users
// @Router /users/{id}/online-status [put]
func (h usersHandler) UpdateOnlineStatus(ctx *gin.Context) {
	// TODO
}
