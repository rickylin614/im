package middleware

import (
	"go.uber.org/dig"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/service"

	"github.com/redis/go-redis/v9"
)

func NewMiddleware(in digIn) *Middleware {
	return &Middleware{
		Auth:  authMiddleware{in: in},
		Cache: CacheMiddleware{in: in},
		Rate:  RateLimitMiddleware{in: in},
	}
}

type Middleware struct {
	Auth  authMiddleware
	Cache CacheMiddleware
	Rate  RateLimitMiddleware
}

type digIn struct {
	dig.In

	Logger  logger.Logger
	Service *service.Service
	Config  *config.Config
	Rdb     redis.UniversalClient
}
