package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type usersHandler struct {
	in digIn
}

// Get
// @Summary Get
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
// @Summary GetList
// @Tags users
// @Param body body request.UsersGetList true "param"
// @Success 200 {object} response.APIResponse[response.UsersGetList]
// @Router /users [get]
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
// @Summary Create
// @Tags users
// @Param body body request.UsersCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /users [post]
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
// @Summary Update
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

// Delete
// @Summary Delete
// @Tags users
// @Param body body request.UsersDelete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /users [delete]
func (h usersHandler) Delete(ctx *gin.Context) {
	req := &request.UsersDelete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.UsersSrv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
