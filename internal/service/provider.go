package service

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/repository"
)

// digIn repository require independence
type digIn struct {
	dig.In

	Repository *repository.Repository
	Logger     *logger.Logger
}

type Service struct {
}

func NewService(in digIn) *Service {
	return &Service{}
}
