package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IExampleService interface {
	Get(ctx *gin.Context, cond *req.ExampleGet) (*models.Example, error)
	GetList(ctx *gin.Context, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error)
	Create(ctx *gin.Context, cond *req.ExampleCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.ExampleUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.ExampleDelete) (err error)
}

func NewExampleService(in digIn) IExampleService {
	return exampleService{in: in}
}

type exampleService struct {
	in digIn
}

func (s exampleService) Get(ctx *gin.Context, cond *req.ExampleGet) (*models.Example, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Get(db, cond)
}

func (s exampleService) GetList(ctx *gin.Context, cond *req.ExampleGetList) (*models.PageResult[*models.Example], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.GetList(db, cond)
}

func (s exampleService) Create(ctx *gin.Context, cond *req.ExampleCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.Example{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.ExampleRepo.Create(db, insertData)
}

func (s exampleService) Update(ctx *gin.Context, cond *req.ExampleUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.Example{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.ExampleRepo.Update(db, updateData)
}

func (s exampleService) Delete(ctx *gin.Context, cond *req.ExampleDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.ExampleRepo.Delete(db, cond.Id)
}
