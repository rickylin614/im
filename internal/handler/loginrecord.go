package handler

import (
	"im/internal/models/request"
	"im/internal/models/response"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type loginRecordHandler struct {
	in webDigIn
}

// GetList
// @Summary 取得登入記錄
// @Tags loginRecord
// @Param request body request.LoginRecordGetList true "param"
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
