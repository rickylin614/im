package service

import (
	"context"
	"im/internal/models"
	"im/internal/models/req"

	"github.com/jinzhu/copier"
)

type IExampleService interface {
	Get(ctx context.Context, cond *req.ExampleGet) (*models.Example, error)
	GetList(ctx context.Context, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error)
	Create(ctx context.Context, cond *req.ExamplePost) (err error)
	Update(ctx context.Context, cond *req.ExamplePut) (err error)
	Delete(ctx context.Context, cond *req.ExampleDelete) (err error)
}

func NewExampleService(in digIn) IExampleService {
	return ExampleService{in: in}
}

type ExampleService struct {
	in digIn
}

func (s ExampleService) Get(ctx context.Context, cond *req.ExampleGet) (*models.Example, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Get(db, cond)
}

func (s ExampleService) GetList(ctx context.Context, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.GetList(db, cond)
}

func (s ExampleService) Create(ctx context.Context, cond *req.ExamplePost) (err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Example{}
	if err := copier.Copy(insertData, cond); err != nil {
		return err
	}
	return s.in.Repository.ExampleRepo.Create(db, insertData)
}

func (s ExampleService) Update(ctx context.Context, cond *req.ExamplePut) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Example{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.ExampleRepo.Update(db, updateData)
}

func (s ExampleService) Delete(ctx context.Context, cond *req.ExampleDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Delete(db, cond.Id)
}
