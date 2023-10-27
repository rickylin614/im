package middleware

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/service"
)

func NewMiddleware(in digIn) *Middleware {
	return &Middleware{
		Auth:  authMiddleware{in: in},
		Cache: CacheMiddleware{in: in},
	}
}

type Middleware struct {
	Auth  authMiddleware
	Cache CacheMiddleware
}

type digIn struct {
	dig.In

	Logger  logger.Logger
	Service *service.Service
}
