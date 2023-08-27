package repository

import (
	"im/internal/pkg/mdb"

	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

func NewRepository(in digIn) *Repository {
	return &Repository{in: in,
		ExampleRepo: NewExampleRepository(in),
	}
}

type Repository struct {
	in digIn

	ExampleRepo IExampleRepository
}

// digIn repository require indendency
type digIn struct {
	dig.In

	Db  mdb.Client
	Rdb redis.UniversalClient
}
