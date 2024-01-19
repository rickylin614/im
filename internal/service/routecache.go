package service

import (
	"github.com/gin-gonic/gin"

	"im/internal/models/po"
)

type IRouteCacheService interface {
	Get(ctx *gin.Context, cond *po.RouteCacheGet) (*po.RouteCache, error)
	Set(ctx *gin.Context, cond *po.RouteCacheSet) error
}

func NewRouteCacheService(in DigIn) IRouteCacheService {
	return routeCacheService{in: in}
}

type routeCacheService struct {
	in DigIn
}

func (s routeCacheService) Get(ctx *gin.Context, cond *po.RouteCacheGet) (*po.RouteCache, error) {
	return s.in.Repository.CacheRepo.GetRouteCache(ctx, cond)
}

func (s routeCacheService) Set(ctx *gin.Context, cond *po.RouteCacheSet) error {
	return s.in.Repository.CacheRepo.SetRouteCache(ctx, cond)
}
