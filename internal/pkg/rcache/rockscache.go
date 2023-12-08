package rcache

import (
	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
)

func NewRocksCache(rdb redis.UniversalClient) *rockscache.Client {
	rcOptions := rockscache.NewDefaultOptions()
	rcOptions.StrongConsistency = true
	rcOptions.RandomExpireAdjustment = 0.2

	return rockscache.NewClient(rdb, rcOptions)
}
