package service

import (
	"{{ .ProjectName }}/internal/models/po"
	"{{ .ProjectName }}/internal/models/request"
	"{{ .ProjectName }}/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type I{{ .FileName }}Service interface {
	Get(ctx *gin.Context, cond *request.{{ .FileName }}Get) (*po.{{ .FileName }}, error)
	GetList(ctx *gin.Context, cond *request.{{ .FileName }}GetList) (*po.PageResult[*po.{{ .FileName }}], error)
	Create(ctx *gin.Context, cond *request.{{ .FileName }}Create) (id any, err error)
	Update(ctx *gin.Context, cond *request.{{ .FileName }}Update) (err error)
	Delete(ctx *gin.Context, cond *request.{{ .FileName }}Delete) (err error)
}

func New{{ .FileName }}Service(in DigIn) I{{ .FileName }}Service {
	return {{ .FileNameTitleLower }}Service{In: in}
}

type {{ .FileNameTitleLower }}Service struct {
	In DigIn
}

func (s {{ .FileNameTitleLower }}Service) Get(ctx *gin.Context, cond *request.{{ .FileName }}Get) (*po.{{ .FileName }}, error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.{{ .FileName }}Repo.Get(db, cond)
}

func (s {{ .FileNameTitleLower }}Service) GetList(ctx *gin.Context, cond *request.{{ .FileName }}GetList) (*po.PageResult[*po.{{ .FileName }}], error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.{{ .FileName }}Repo.GetList(db, cond)
}

func (s {{ .FileNameTitleLower }}Service) Create(ctx *gin.Context, cond *request.{{ .FileName }}Create) (id any, err error) {
	db := s.In.DB.Session(ctx)
	insertData := &po.{{ .FileName }}{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.In.Repository.{{ .FileName }}Repo.Create(db, insertData)
}

func (s {{ .FileNameTitleLower }}Service) Update(ctx *gin.Context, cond *request.{{ .FileName }}Update) (err error) {
	db := s.In.DB.Session(ctx)
	updateData := &po.{{ .FileName }}{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.In.Repository.{{ .FileName }}Repo.Update(db, updateData)
}

func (s {{ .FileNameTitleLower }}Service) Delete(ctx *gin.Context, cond *request.{{ .FileName }}Delete) (err error) {
	db := s.In.DB.Session(ctx)
	return s.In.Repository.{{ .FileName }}Repo.Delete(db, cond.ID)
}
