package redis

import (
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"

	"im/internal/pkg/config"
)

type digIn struct {
	dig.In

	Config *config.Config
}

// NewRedis
func NewRedis(in digIn) redis.UniversalClient {
	var rdb redis.UniversalClient
	addrs := strings.Split(in.Config.RedisConfig.Address, ",")
	if len(addrs) == 1 {
		rdb = redis.NewClient(&redis.Options{
			Addr:            in.Config.RedisConfig.Address,
			Password:        in.Config.RedisConfig.Password,
			PoolSize:        in.Config.RedisConfig.MaxActive,
			MaxIdleConns:    in.Config.RedisConfig.MaxIdle,
			DialTimeout:     time.Duration(in.Config.RedisConfig.ConnectTimeout) * time.Second,
			ReadTimeout:     time.Duration(in.Config.RedisConfig.ReadTimeout) * time.Second,
			WriteTimeout:    time.Duration(in.Config.RedisConfig.WriteTimeout) * time.Second,
			ConnMaxIdleTime: time.Duration(in.Config.RedisConfig.IdleTimeout) * time.Second,
		})
	} else {
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:           addrs,
			Password:        in.Config.RedisConfig.Password,
			PoolSize:        in.Config.RedisConfig.MaxActive,
			MaxIdleConns:    in.Config.RedisConfig.MaxIdle,
			DialTimeout:     time.Duration(in.Config.RedisConfig.ConnectTimeout) * time.Second,
			ReadTimeout:     time.Duration(in.Config.RedisConfig.ReadTimeout) * time.Second,
			WriteTimeout:    time.Duration(in.Config.RedisConfig.WriteTimeout) * time.Second,
			ConnMaxIdleTime: time.Duration(in.Config.RedisConfig.IdleTimeout) * time.Second,
		})
	}

	return rdb
}

func NewRedisLock(in redis.UniversalClient) *redsync.Redsync {
	pool := goredis.NewPool(in)
	return redsync.New(pool)
}
