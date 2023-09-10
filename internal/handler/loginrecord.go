package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type loginRecordHandler struct {
	in digIn
}

// GetList
// @Summary GetList
// @Tags loginRecord
// @Param body body request.LoginRecordGetList true "param"
// @Success 200 {object} response.APIResponse[response.LoginRecordGetList]
// @Router /loginRecord [get]
func (h loginRecordHandler) GetList(ctx *gin.Context) {
	req := &request.LoginRecordGetList{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.LoginRecordSrv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.LoginRecordGetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary Create
// @Tags loginRecord
// @Param body body request.LoginRecordCreate true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /loginRecord [post]
func (h loginRecordHandler) Create(ctx *gin.Context) {
	req := &request.LoginRecordCreate{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.LoginRecordSrv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}
