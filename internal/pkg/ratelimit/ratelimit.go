package ratelimit

// import (
// 	"context"
// 	"im/internal/pkg/config"
// 	"net/http"

// 	"github.com/redis/go-redis/v9"
// 	"github.com/throttled/throttled/v2"
// 	"github.com/throttled/throttled/v2/store/goredisstore"
// 	"go.uber.org/dig"
// )

// // type Config struct {
// // 	Rate           int                   // 请求速率
// // 	Burst          int                   // 请求突发数
// // 	StoreSize      int                   // 为内存存储定义大小
// // 	RedisClient    redis.UniversalClient // Redis客户端
// // 	UseMemoryStore bool                  // 使用内存存储还是Redis存储
// // }

// type digIn struct {
// 	dig.In

// 	Config *config.Config
// 	Rdb    redis.UniversalClient
// }

// func NewRateLimiter(in digIn) (http.Handler, error) {
// 	var store throttled.GCRAStore
// 	if in.Config.RateConfig.UseMemoryStore {
// 		store = throttled.NewMemStore(in.Config.RateConfig.StoreSize)
// 	} else {
// 		rs, err := goredisstore.NewCtx(context.Background(), "")
// 		if err != nil {
// 			return nil, err
// 		}
// 		store = rs
// 	}

// 	quota := throttled.RateQuota{throttled.PerSec(in.Config.RateConfig.Rate), in.Config.RateConfig.Burst}
// 	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return throttled.HTTPRateLimiter{
// 		RateLimiter: rateLimiter,
// 		VaryBy:      &throttled.VaryBy{Path: true},
// 	}, nil
// }
