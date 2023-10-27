package service

import (
	"github.com/gin-gonic/gin"

	"im/internal/models/req"
)

type IRouteCacheService interface {
	Get(ctx *gin.Context, cond *req.RouteCacheGet) (string, error)
	Set(ctx *gin.Context, cond *req.RouteCacheSet) error
}

func NewRouteCacheService(in digIn) IRouteCacheService {
	return routeCacheService{in: in}
}

type routeCacheService struct {
	in digIn
}

func (s routeCacheService) Get(ctx *gin.Context, cond *req.RouteCacheGet) (string, error) {
	return s.in.Repository.RouteCacheRepo.Get(ctx, cond)
}

func (s routeCacheService) Set(ctx *gin.Context, cond *req.RouteCacheSet) error {
	return s.in.Repository.RouteCacheRepo.Set(ctx, cond)
}
