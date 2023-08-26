package provider

import (
	"sync"

	"go.uber.org/dig"

	"im/internal/handler"
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/mdb"
	"im/internal/pkg/mongo"
	"im/internal/pkg/redis"
	"im/internal/repository"
	"im/internal/router"
	"im/internal/server"
	"im/internal/service"
)

var (
	once sync.Once
)

func New() *dig.Container {
	var container *dig.Container
	once.Do(func() {
		container = dig.New()
		if err := container.Provide(logger.NewLogger); err != nil {
			panic(err)
		}
		if err := container.Provide(config.NewConfig); err != nil {
			panic(err)
		}
		if err := container.Provide(mdb.NewDB); err != nil {
			panic(err)
		}
		if err := container.Provide(redis.NewRedis); err != nil {
			panic(err)
		}
		if err := container.Provide(mongo.NewMongoDB); err != nil {
			panic(err)
		}
		if err := container.Provide(router.NewRouter); err != nil {
			panic(err)
		}
		if err := container.Provide(handler.NewHandler); err != nil {
			panic(err)
		}
		if err := container.Provide(service.NewService); err != nil {
			panic(err)
		}
		if err := container.Provide(repository.NewRepository); err != nil {
			panic(err)
		}
		if err := container.Provide(server.NewServerController); err != nil {
			panic(err)
		}
	})

	return container
}
