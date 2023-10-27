package middleware

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/service"
)

func NewMiddleware(in digIn) *Middleware {
	return &Middleware{Auth: authMiddleware{in: in}}
}

type Middleware struct {
	Auth authMiddleware
}

type digIn struct {
	dig.In

	Logger  logger.Logger
	Service *service.Service
}
