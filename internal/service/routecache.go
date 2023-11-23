package service

import (
	"github.com/gin-gonic/gin"

	"im/internal/models"
)

type IRouteCacheService interface {
	Get(ctx *gin.Context, cond *models.RouteCacheGet) (*models.RouteCache, error)
	Set(ctx *gin.Context, cond *models.RouteCacheSet) error
}

func NewRouteCacheService(in DigIn) IRouteCacheService {
	return routeCacheService{in: in}
}

type routeCacheService struct {
	in DigIn
}

func (s routeCacheService) Get(ctx *gin.Context, cond *models.RouteCacheGet) (*models.RouteCache, error) {
	return s.in.Repository.CacheRepo.GetRouteCache(ctx, cond)
}

func (s routeCacheService) Set(ctx *gin.Context, cond *models.RouteCacheSet) error {
	return s.in.Repository.CacheRepo.SetRouteCache(ctx, cond)
}
