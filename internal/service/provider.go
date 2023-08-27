package service

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/pkg/mdb"
	"im/internal/repository"
)

func NewService(in digIn) *Service {
	return &Service{in: in,
		ExampleSrv: NewExampleService(in),
	}
}

type Service struct {
	in digIn

	ExampleSrv IExampleService
}

// digIn repository require independence
type digIn struct {
	dig.In

	Repository *repository.Repository
	Logger     *logger.Logger
	DB         mdb.Client
}
