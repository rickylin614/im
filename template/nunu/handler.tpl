package handler

import (
	"{{ .ProjectName }}/internal/models/request"
	"{{ .ProjectName }}/internal/models/response"
	"{{ .ProjectName }}/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type {{ .FileNameTitleLower }}Handler struct {
	in digIn
}

// Get
// @Summary Get
// @Tags {{ .FileNameTitleLower }}
// @Param request query request.{{ .FileName }}Get true "param"
// @Success 200 {object} response.APIResponse[response.{{ .FileName }}Get]
// @Router /{{ .FileNameKebabCase }}/:id [get]
func (h {{ .FileNameTitleLower }}Handler) Get(ctx *gin.Context) {
	req := &request.{{ .FileName }}Get{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.{{ .FileName }}Srv.Get(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	if data == nil {
		ctxs.SetResp(ctx, data)
	}
	result := &response.{{ .FileName }}Get{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// GetList
// @Summary GetList
// @Tags {{ .FileNameTitleLower }}
// @Param request query request.{{ .FileName }}GetList true "param"
// @Success 200 {object} response.APIResponse[response.{{ .FileName }}GetList]
// @Router /{{ .FileNameKebabCase }} [get]
func (h {{ .FileNameTitleLower }}Handler) GetList(ctx *gin.Context) {
	req := &request.{{ .FileName }}GetList{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	data, err := h.in.Service.{{ .FileName }}Srv.GetList(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	result := &response.{{ .FileName }}GetList{}
	if err := copier.Copy(result, data); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, result)
}

// Create
// @Summary Create
// @Tags {{ .FileNameTitleLower }}
// @Param request body request.{{ .FileName }}Create true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /{{ .FileNameKebabCase }} [post]
func (h {{ .FileNameTitleLower }}Handler) Create(ctx *gin.Context) {
	req := &request.{{ .FileName }}Create{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	id, err := h.in.Service.{{ .FileName }}Srv.Create(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetResp(ctx, id)
}

// Update
// @Summary Update
// @Tags {{ .FileNameTitleLower }}
// @Param request body request.{{ .FileName }}Update true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /{{ .FileNameKebabCase }} [put]
func (h {{ .FileNameTitleLower }}Handler) Update(ctx *gin.Context) {
	req := &request.{{ .FileName }}Update{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}

	err := h.in.Service.{{ .FileName }}Srv.Update(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}

// Delete
// @Summary Delete
// @Tags {{ .FileNameTitleLower }}
// @Param request body request.{{ .FileName }}Delete true "param"
// @Success 200 {object} response.APIResponse[string]
// @Router /{{ .FileNameKebabCase }} [delete]
func (h {{ .FileNameTitleLower }}Handler) Delete(ctx *gin.Context) {
	req := &request.{{ .FileName }}Delete{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	err := h.in.Service.{{ .FileName }}Srv.Delete(ctx, req)
	if err != nil {
		ctxs.SetError(ctx, err)
		return
	}
	ctxs.SetSuccessResp(ctx)
}
