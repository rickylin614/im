package repository

import (
	"im/internal/pkg/logger"
	"im/internal/pkg/sqldb"

	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

func NewRepository(in digIn) *Repository {
	return &Repository{in: in,
		ExampleRepo:     NewExampleRepository(in),
		UsersRepo:       NewUsersRepository(in),
		LoginRecordRepo: NewLoginRecordRepository(in),
	}
}

type Repository struct {
	in digIn

	ExampleRepo     IExampleRepository
	UsersRepo       IUsersRepository
	LoginRecordRepo ILoginRecordRepository
}

// digIn repository require indendency
type digIn struct {
	dig.In

	Logger *logger.Logger
	Db     sqldb.Client
	Rdb    redis.UniversalClient
}
