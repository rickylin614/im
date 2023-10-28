package repository

import (
	"context"

	"golang.org/x/sync/singleflight"

	"im/internal/models/req"
)

type IRouteCacheRepository interface {
	Get(ctx context.Context, cond *req.RouteCacheGet) (string, error)
	Set(ctx context.Context, cond *req.RouteCacheSet) error
}

func NewRouteCacheRepository(in digIn) IRouteCacheRepository {
	return routeCacheRepository{in: in, group: singleflight.Group{}}
}

type routeCacheRepository struct {
	in    digIn
	group singleflight.Group
}

func (r routeCacheRepository) Get(ctx context.Context, cond *req.RouteCacheGet) (string, error) {
	// 使用 singleflight 確保只有一個 goroutine 呼叫此 function 對於相同的一個 key
	data, err, _ := r.group.Do(cond.RouteCacheKey, func() (interface{}, error) {
		if data := r.in.Rdb.Get(ctx, cond.RouteCacheKey); data.Err() != nil {
			return "", data.Err()
		} else {
			return data.String(), nil
		}
	})
	if err != nil {
		return "", err
	}
	return data.(string), nil
}

func (r routeCacheRepository) Set(ctx context.Context, cond *req.RouteCacheSet) error {
	return r.in.Rdb.Set(ctx, cond.RouteCacheKey, cond.RouteCacheData, cond.TTL).Err()
}