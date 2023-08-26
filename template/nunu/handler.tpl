package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"{{ .ProjectName }}/service/internal/errs"
	"{{ .ProjectName }}/service/internal/model/bo"
	"{{ .ProjectName }}/service/internal/model/dto"
	"{{ .ProjectName }}/service/internal/util/ctxs"
)

type {{ .FileNameTitleLower }}Handler struct {
	in digIn
}

// @Summary Get{{ .FileName }}
// @Description This endpoint receives and sends back an {{ .FileName }} struct
// @Tags {{ .FileNameTitleLower }}
// @ID get-{{ .FileNameTitleLower }}
// @Accept  json
// @Produce  json
// @Param body body dto.Query{{ .FileName }}Cond true "request param"
// @Success 200 {object} dto.StandardResponse[dto.{{ .FileName }}]
// @Security ApiKeyAuth
// @Router /{{ .FileNameTitleLower }} [get]
func (h *{{ .FileNameTitleLower }}Handler) Get{{ .FileName }}ByID(ctx *gin.Context) {
	req := &dto.Query{{ .FileName }}Cond{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, errs.RequestParamParseFailed, err)
		return
	}

	cond := &bo.Query{{ .FileName }}Cond{}
	if err := copier.Copy(cond, req); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, err)
		return
	}

	{{ .FileNameTitleLower }}, err := h.in.Module.{{ .FileName }}Module.Get{{ .FileName }}ByID(ctx, cond)
	if err != nil {
		ctxs.SetError(ctx, errs.CommonUnknownError, err)
		return
	}

	result := &dto.{{ .FileName }}{}
	if err := copier.Copy(result, {{ .FileNameTitleLower }}); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, err)
		return
	}

	ctxs.SetResp(ctx, result)
}

// @Summary Get{{ .FileName }}List
// @Description This endpoint receives and sends back an {{ .FileName }} struct
// @Tags {{ .FileNameTitleLower }}
// @ID get-{{ .FileNameTitleLower }}-list
// @Accept  json
// @Produce  json
// @Param body body dto.Query{{ .FileName }}Cond true "request param"
// @Success 200 {object} dto.StandardResponse[[]dto.{{ .FileName }}]
// @Security ApiKeyAuth
// @Router /{{ .FileNameTitleLower }}/list [get]
func (h *{{ .FileNameTitleLower }}Handler) Get{{ .FileName }}(ctx *gin.Context) {
	req := &dto.Query{{ .FileName }}Cond{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctxs.SetError(ctx, errs.RequestParamParseFailed, nil)
		return
	}

	cond := &bo.Query{{ .FileName }}Cond{}
	if err := copier.Copy(cond, req); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, nil)
	}

	{{ .FileNameTitleLower }}, boPager, err := h.in.Module.{{ .FileName }}Module.Get{{ .FileName }}(ctx, cond)
	if err != nil {
		ctxs.SetError(ctx, errs.CommonUnknownError, err)
		return
	}

	{{ .FileNameTitleLower }}s := make([]*dto.{{ .FileName }}, 0)
	if err := copier.Copy(&{{ .FileNameTitleLower }}s, {{ .FileNameTitleLower }}); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, err)
		return
	}
	page := &dto.PageResult{}
	if err := copier.Copy(page, boPager); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, err)
		return
	}
	result := dto.{{ .FileName }}Response{
		{{ .FileName }}: {{ .FileNameTitleLower }}s, 
		Page: page,
	}

	ctxs.SetResp(ctx, result)
}

// @Summary Create{{ .FileName }}
// @Description This endpoint receives and sends back an {{ .FileName }} struct
// @Tags {{ .FileNameTitleLower }}
// @ID create-{{ .FileNameTitleLower }}
// @Accept  json
// @Produce  json
// @Param body body dto.Create{{ .FileName }}Cond true "request param"
// @Success 200 {object} dto.StandardResponse[string]	"0"
// @Security ApiKeyAuth
// @Router /{{ .FileNameTitleLower }} [post]
func (h *{{ .FileNameTitleLower }}Handler) Create{{ .FileName }}(ctx *gin.Context) {
	req := &dto.Create{{ .FileName }}Cond{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, errs.RequestParamParseFailed, nil)
		return
	}

	cond := &bo.Create{{ .FileName }}Cond{}
	if err := copier.Copy(cond, req); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, nil)
	}

	_, err := h.in.Module.{{ .FileName }}Module.Create{{ .FileName }}(ctx, cond)
	if err != nil {
		ctxs.SetError(ctx, errs.CommonUnknownError, err)
		return
	}

	ctxs.SetResp(ctx, nil)
}

// @Summary Update{{ .FileName }}
// @Description This endpoint receives and sends back an {{ .FileName }} struct
// @Tags {{ .FileNameTitleLower }}
// @ID update-{{ .FileNameTitleLower }}
// @Accept  json
// @Produce  json
// @Param body body dto.Update{{ .FileName }}Cond true "request param"
// @Success 200 {string} dto.StandardResponse[string] "0"
// @Security ApiKeyAuth
// @Router /{{ .FileNameTitleLower }} [put]
func (h *{{ .FileNameTitleLower }}Handler) Update{{ .FileName }}(ctx *gin.Context) {
	req := &dto.Update{{ .FileName }}Cond{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, errs.RequestParamParseFailed, nil)
		return
	}

	cond := &bo.Update{{ .FileName }}Cond{}
	if err := copier.Copy(cond, req); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, nil)
	}

	_, err := h.in.Module.{{ .FileName }}Module.Update{{ .FileName }}(ctx, cond)
	if err != nil {
		ctxs.SetError(ctx, errs.CommonUnknownError, err)
		return
	}

	ctxs.SetResp(ctx, nil)
}

// @Summary Delete{{ .FileName }}
// @Description This endpoint receives and sends back an {{ .FileName }} struct
// @Tags {{ .FileNameTitleLower }}
// @ID delete-{{ .FileNameTitleLower }}
// @Accept  json
// @Produce  json
// @Param body body dto.Delete{{ .FileName }}Cond true "request param"
// @Success 200 {object} dto.StandardResponse[string] "0"
// @Security ApiKeyAuth
// @Router /{{ .FileNameTitleLower }} [delete]
func (h *{{ .FileNameTitleLower }}Handler) Delete{{ .FileName }}(ctx *gin.Context) {
	req := &dto.Delete{{ .FileName }}Cond{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctxs.SetError(ctx, errs.RequestParamParseFailed, nil)
		return
	}

	cond := &bo.Delete{{ .FileName }}Cond{}
	if err := copier.Copy(cond, req); err != nil {
		ctxs.SetError(ctx, errs.CommonParseError, nil)
	}

	err := h.in.Module.{{ .FileName }}Module.Delete{{ .FileName }}(ctx, cond)
	if err != nil {
		ctxs.SetError(ctx, errs.CommonUnknownError, err)
		return
	}

	ctxs.SetResp(ctx, nil)
}
