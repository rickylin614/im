package provider

import (
	"sync"

	"go.uber.org/dig"

	"im/internal/handler"
	"im/internal/listener"
	"im/internal/manager/msggateway"
	"im/internal/middleware"
	"im/internal/pkg/config"
	"im/internal/pkg/localcache"
	"im/internal/pkg/logger"
	"im/internal/pkg/mongo"
	"im/internal/pkg/prom"
	"im/internal/pkg/queue"
	"im/internal/pkg/rcache"
	"im/internal/pkg/redis"
	"im/internal/pkg/signalctx"
	"im/internal/pkg/sqldb"
	"im/internal/repository"
	"im/internal/repository/sql"
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
		// 基礎套件
		if err := container.Provide(signalctx.NewContext); err != nil {
			panic(err)
		}
		if err := container.Provide(logger.NewLogger); err != nil {
			panic(err)
		}
		if err := container.Provide(config.NewConfig); err != nil {
			panic(err)
		}
		if err := container.Provide(localcache.NewLocalCache); err != nil {
			panic(err)
		}
		if err := container.Provide(prom.NewPromManager); err != nil {
			panic(err)
		}
		// localcache
		// 需要連線配置的套件
		if err := container.Provide(sqldb.NewDB); err != nil {
			panic(err)
		}
		if err := container.Provide(redis.NewRedis); err != nil {
			panic(err)
		}
		if err := container.Provide(rcache.NewRocksCache); err != nil {
			panic(err)
		}
		if err := container.Provide(mongo.NewMongoDB); err != nil {
			panic(err)
		}
		if err := container.Provide(queue.NewQueue); err != nil {
			panic(err)
		}
		// Web業務配置
		if err := container.Provide(router.NewRouter); err != nil {
			panic(err)
		}
		if err := container.Provide(router.NewWsRouter); err != nil {
			panic(err)
		}
		if err := container.Provide(middleware.NewMiddleware); err != nil {
			panic(err)
		}
		if err := container.Provide(handler.NewWebHandler); err != nil {
			panic(err)
		}
		if err := container.Provide(handler.NewWebSocketHandler); err != nil {
			panic(err)
		}
		if err := container.Provide(msggateway.NewWsManger); err != nil {
			panic(err)
		}
		if err := container.Provide(listener.NewListener); err != nil {
			panic(err)
		}
		if err := container.Provide(service.NewService); err != nil {
			panic(err)
		}
		if err := container.Provide(sql.NewSqlEmbedFile); err != nil {
			panic(err)
		}
		if err := container.Provide(repository.NewRepository); err != nil {
			panic(err)
		}
		// 啟動程序
		if err := container.Provide(server.NewServerController); err != nil {
			panic(err)
		}
	})

	return container
}
