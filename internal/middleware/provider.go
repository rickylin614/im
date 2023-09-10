package middleware

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/service"
)

func NewHandler(in digIn) *Middleware {
	return &Middleware{}
}

type Middleware struct {
}

type digIn struct {
	dig.In

	Logger  *logger.Logger
	Service *service.Service
}
