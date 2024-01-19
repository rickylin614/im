package service

import (
	"im/internal/models/po"
	"im/internal/models/request"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IExampleService interface {
	Get(ctx *gin.Context, cond *request.ExampleGet) (*po.Example, error)
	GetList(ctx *gin.Context, cond *request.ExampleGetList) (*po.PageResult[*po.Example], error)
	Create(ctx *gin.Context, cond *request.ExampleCreate) (id any, err error)
	Update(ctx *gin.Context, cond *request.ExampleUpdate) (err error)
	Delete(ctx *gin.Context, cond *request.ExampleDelete) (err error)
}

func NewExampleService(in DigIn) IExampleService {
	return exampleService{in: in}
}

type exampleService struct {
	in DigIn
}

func (s exampleService) Get(ctx *gin.Context, cond *request.ExampleGet) (*po.Example, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Get(db, cond)
}

func (s exampleService) GetList(ctx *gin.Context, cond *request.ExampleGetList) (*po.PageResult[*po.Example], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.GetList(db, cond)
}

func (s exampleService) Create(ctx *gin.Context, cond *request.ExampleCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &po.Example{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.ExampleRepo.Create(db, insertData)
}

func (s exampleService) Update(ctx *gin.Context, cond *request.ExampleUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &po.Example{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.ExampleRepo.Update(db, updateData)
}

func (s exampleService) Delete(ctx *gin.Context, cond *request.ExampleDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Delete(db, cond.Id)
}
