package service

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"
	"im/internal/repository"
)

func NewService(in digIn) *Service {
	return &Service{in: in,
		ExampleSrv: NewExampleService(in),
		UsersSrv:   NewUsersService(in),
	}
}

type Service struct {
	in digIn

	ExampleSrv IExampleService
	UsersSrv   IUsersService
}

// digIn repository require independence
type digIn struct {
	dig.In

	Repository *repository.Repository
	Logger     *logger.Logger
	DB         sqldb.Client
}
