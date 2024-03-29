package cache

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"im/internal/util/errs"

	"github.com/coocood/freecache"
	"github.com/dtm-labs/rockscache"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

func GetCache[T any](ctx context.Context,
	c *freecache.Cache,
	rdb redis.UniversalClient,
	lock *singleflight.Group,
	key string,
	fn func() (T, error),
	ttl time.Duration,
) (result T, err error) {
	cacheData, err := c.Get([]byte(key))
	if err == nil {
		if err := json.Unmarshal(cacheData, &result); err != nil {
			return result, nil
		}
	}

	// use singleflight block other req and get redis
	data, err, _ := lock.Do(key, func() (any, error) {
		// 從redis取得值
		if data := rdb.Get(ctx, key); data.Err() != nil {
			if !errors.Is(data.Err(), redis.Nil) {
				slog.ErrorContext(ctx, err.Error())
			}
			return "", data.Err()
		} else {
			c.Set([]byte(key), []byte(data.Val()), 2)
			return data.Bytes()
		}
	})

	if v, ok := data.([]byte); err == nil && ok {
		if err := json.Unmarshal(v, result); err == nil {
			return result, nil
		}
	}

	data, err, _ = lock.Do(key, func() (any, error) {
		// TODO 增加集群鎖防止髒讀
		r, err := fn()
		if err != nil {
			return nil, err
		}
		if b, err := json.Marshal(r); err == nil {
			rdb.Set(ctx, key, b, ttl)
		} else {
			slog.ErrorContext(ctx, err.Error())
		}
		return r, nil
	})

	if err != nil {
		return
	}
	if v, ok := data.(T); ok {
		result = v
	}
	return
}

func GetCache2[T any](ctx context.Context,
	c *freecache.Cache,
	rcClient *rockscache.Client,
	key string,
	fn func() (T, error),
	ttl time.Duration,
) (result T, err error) {
	cacheData, err := c.Get([]byte(key))
	if err == nil {
		if err := json.Unmarshal(cacheData, &result); err != nil {
			return result, nil
		}
	}
	writer := false

	v, err := rcClient.Fetch2(ctx, key, ttl, func() (string, error) {
		result, err = fn()
		if err != nil {
			return "", err
		}
		bs, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(bs), err
	})

	if err != nil {
		return
	}
	if writer {
		return
	}

	if v == "" {
		err = errs.RequestNoData
		return
	}

	err = json.Unmarshal([]byte(v), &result)
	return
}
