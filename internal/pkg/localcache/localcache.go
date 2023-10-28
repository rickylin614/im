package localcache

import (
	"im/internal/pkg/config"

	"github.com/coocood/freecache"
	"go.uber.org/dig"
)

type digIn struct {
	dig.In

	Config *config.Config
}

func NewLocalCache(in digIn) *freecache.Cache {
	// TODO 待決定是否要抽象
	return freecache.NewCache(in.Config.CacheConfig.CacheSize)
}
