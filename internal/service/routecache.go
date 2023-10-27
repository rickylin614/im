package service

import (
	"im/internal/models"
	"im/internal/models/req"
	"im/internal/util/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IRouteCacheService interface {
	Get(ctx *gin.Context, cond *req.RouteCacheGet) (*models.RouteCache, error)
	GetList(ctx *gin.Context, cond *req.RouteCacheGetList) (*models.PageResult[*models.RouteCache], error)
	Create(ctx *gin.Context, cond *req.RouteCacheCreate) (id any, err error)
	Update(ctx *gin.Context, cond *req.RouteCacheUpdate) (err error)
	Delete(ctx *gin.Context, cond *req.RouteCacheDelete) (err error)
}

func NewRouteCacheService(in digIn) IRouteCacheService {
	return routeCacheService{in: in}
}

type routeCacheService struct {
	in digIn
}

func (s routeCacheService) Get(ctx *gin.Context, cond *req.RouteCacheGet) (*models.RouteCache, error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.RouteCacheRepo.Get(db, cond)
}

func (s routeCacheService) GetList(ctx *gin.Context, cond *req.RouteCacheGetList) (*models.PageResult[*models.RouteCache], error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.RouteCacheRepo.GetList(db, cond)
}

func (s routeCacheService) Create(ctx *gin.Context, cond *req.RouteCacheCreate) (id any, err error) {
	db := s.in.DB.Session(ctx)
	insertData := &models.RouteCache{ID: uuid.New()}
	if err := copier.Copy(insertData, cond); err != nil {
		return nil, err
	}
	return s.in.Repository.RouteCacheRepo.Create(db, insertData)
}

func (s routeCacheService) Update(ctx *gin.Context, cond *req.RouteCacheUpdate) (err error) {
	db := s.in.DB.Session(ctx)
	updateData := &models.RouteCache{}
	if err := copier.Copy(updateData, cond); err != nil {
		return err
	}
	return s.in.Repository.RouteCacheRepo.Update(db, updateData)
}

func (s routeCacheService) Delete(ctx *gin.Context, cond *req.RouteCacheDelete) (err error) {
	db := s.in.DB.Session(ctx)
	return s.in.Repository.RouteCacheRepo.Delete(db, cond.ID)
}
