package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"

	"im/internal/models"
	"im/internal/util/errs"
)

type IRouteCacheRepository interface {
	Get(ctx context.Context, cond *models.RouteCacheGet) (*models.RouteCache, error)
	Set(ctx context.Context, cond *models.RouteCacheSet) error
}

func NewRouteCacheRepository(in digIn) IRouteCacheRepository {
	return &routeCacheRepository{in: in, group: singleflight.Group{}}
}

type routeCacheRepository struct {
	in    digIn
	group singleflight.Group
	count int
}

func (r *routeCacheRepository) Get(ctx context.Context, cond *models.RouteCacheGet) (*models.RouteCache, error) {
	result := &models.RouteCache{}
	// 設定三秒以內的吃內存
	cacheData, err := r.in.Cache.Get([]byte(cond.RouteCacheKey))
	if err == nil {
		return result.Set(cacheData), nil
	}

	// 使用 singleflight 確保只有一個 goroutine 呼叫此 function 對於相同的一個 key
	data, err, _ := r.group.Do(cond.RouteCacheKey, func() (interface{}, error) {
		fmt.Println(r.count)
		r.count++
		if data := r.in.Rdb.Get(ctx, cond.RouteCacheKey); data.Err() != nil {
			if !errors.Is(data.Err(), redis.Nil) {
				r.in.Logger.Error(ctx, err)
			}
			return "", data.Err()
		} else {
			r.in.Cache.Set([]byte(cond.RouteCacheKey), []byte(data.Val()), 3)
			return data.Bytes()
		}
	})
	if err != nil {
		return nil, err
	}
	if v, ok := data.([]byte); ok {
		return result.Set(v), nil
	}
	return nil, errs.CommonUnknownError
}

func (r *routeCacheRepository) Set(ctx context.Context, cond *models.RouteCacheSet) error {
	r.in.Cache.Set([]byte(cond.RouteCacheKey), cond.RouteCacheData.Bytes(), 3)
	return r.in.Rdb.Set(ctx, cond.RouteCacheKey, cond.RouteCacheData.String(), cond.TTL).Err()
}
