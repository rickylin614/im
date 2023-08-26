package service

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/repository"
)

func NewService(in digIn) *Service {
	return &Service{in: in}
}

type Service struct {
	in digIn
}

// digIn repository require independence
type digIn struct {
	dig.In

	Repository *repository.Repository
	Logger     *logger.Logger
}
