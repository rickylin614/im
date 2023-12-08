package rcache

import (
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
)

func NewRocksCache(rdb redis.UniversalClient) *rockscache.Client {
	rcOptions := rockscache.NewDefaultOptions()
	rcOptions.StrongConsistency = true
	rcOptions.RandomExpireAdjustment = 0.2
	rcOptions.Delay = time.Second

	return rockscache.NewClient(rdb, rcOptions)
}
