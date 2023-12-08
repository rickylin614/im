package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"

	"im/internal/models"
	"im/internal/util/errs"
)

//go:generate mockery --name ICacheRepository --structname MockCacheRepository --filename mock_cache.go --output mock_repository --outpkg mock_repository --with-expecter
type ICacheRepository interface {
	GetRouteCache(ctx context.Context, cond *models.RouteCacheGet) (*models.RouteCache, error)
	SetRouteCache(ctx context.Context, cond *models.RouteCacheSet) error
	GetCache(ctx context.Context, key string) ([]byte, error)
	SetCache(ctx context.Context, key string, value []byte, ttl time.Duration) error
	DelCache(ctx context.Context, key string) error
}

func NewCacheRepository(in digIn) ICacheRepository {
	return &cacheRepository{in: in, group: singleflight.Group{}}
}

type cacheRepository struct {
	in    digIn
	group singleflight.Group
}

func (r *cacheRepository) GetRouteCache(ctx context.Context, cond *models.RouteCacheGet) (*models.RouteCache, error) {
	result := &models.RouteCache{}
	// 設定三秒以內的吃內存
	cacheData, err := r.in.Cache.Get([]byte(cond.RouteCacheKey))
	if err == nil {
		return result.Set(cacheData), nil
	}

	// 使用 singleflight 確保只有一個 goroutine 呼叫此 function 對於相同的一個 key
	data, err, _ := r.group.Do(cond.RouteCacheKey, func() (interface{}, error) {
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

func (r *cacheRepository) SetRouteCache(ctx context.Context, cond *models.RouteCacheSet) error {
	r.in.Cache.Set([]byte(cond.RouteCacheKey), cond.RouteCacheData.Bytes(), 3)
	return r.in.Rdb.Set(ctx, cond.RouteCacheKey, cond.RouteCacheData.String(), cond.TTL).Err()
}

func (r *cacheRepository) GetCache(ctx context.Context, key string) ([]byte, error) {
	cacheData, err := r.in.Cache.Get([]byte(key))
	if err == nil {
		return cacheData, nil
	}

	// 使用 singleflight 確保只有一個 goroutine 呼叫此 function 對於相同的一個 key
	data, err, _ := r.group.Do(key, func() (interface{}, error) {
		// 從redis取得值
		if data := r.in.Rdb.Get(ctx, key); data.Err() != nil {
			if !errors.Is(data.Err(), redis.Nil) {
				r.in.Logger.Error(ctx, err)
			}
			return "", data.Err()
		} else {
			r.in.Cache.Set([]byte(key), []byte(data.Val()), 3)
			return data.Bytes()
		}
	})
	if err != nil {
		return nil, err
	}
	if v, ok := data.([]byte); ok {
		return v, nil
	}
	return nil, errs.CommonUnknownError
}

func (r *cacheRepository) SetCache(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	r.in.Cache.Set([]byte(key), value, 1)
	return r.in.Rdb.Set(ctx, key, value, ttl).Err()
}

func (r *cacheRepository) DelCache(ctx context.Context, key string) error {
	return r.in.Rdb.Del(ctx, key).Err()
}
