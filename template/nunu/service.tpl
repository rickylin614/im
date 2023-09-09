package service

import (
	"context"
	"{{ .ProjectName }}/internal/models"
	"{{ .ProjectName }}/internal/models/req"
	"{{ .ProjectName }}/internal/util/uuid"

	"github.com/jinzhu/copier"
)

type I{{ .FileName }}Service interface {
	Get(ctx context.Context, cond *req.{{ .FileName }}Get) (*models.{{ .FileName }}, error)
	GetList(ctx context.Context, cond *req.{{ .FileName }}GetList) (*models.PageResult[*models.{{ .FileName }}], error)
	Create(ctx context.Context, cond *req.{{ .FileName }}Create) (id any, err error)
	Update(ctx context.Context, cond *req.{{ .FileName }}Update) (err error)
	Delete(ctx context.Context, cond *req.{{ .FileName }}Delete) (err error)
}

func New{{ .FileName }}Service(in digIn) I{{ .FileName }}Service {
	return {{ .FileNameTitleLower }}Service{in: in}
}

type {{ .FileNameTitleLower }}Service struct {
	in digIn
}

func (s {{ .FileNameTitleLower }}Service) Get(ctx context.Context, cond *req.{{ .FileName }}Get) (*models.{{ .FileName }}, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.{{ .FileName }}Repo.Get(db, cond)
}

func (s {{ .FileNameTitleLower }}Service) GetList(ctx context.Context, cond *req.{{ .FileName }}GetList) (*models.PageResult[*models.{{ .FileName }}], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.{{ .FileName }}Repo.GetList(db, cond)
}

func (s {{ .FileNameTitleLower }}Service) Create(ctx context.Context, cond *req.{{ .FileName }}Create) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.{{ .FileName }}{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.{{ .FileName }}Repo.Create(db, insertData)
}

func (s {{ .FileNameTitleLower }}Service) Update(ctx context.Context, cond *req.{{ .FileName }}Update) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.{{ .FileName }}{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.{{ .FileName }}Repo.Update(db, updateData)
}

func (s {{ .FileNameTitleLower }}Service) Delete(ctx context.Context, cond *req.{{ .FileName }}Delete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.{{ .FileName }}Repo.Delete(db, cond.ID)
}
