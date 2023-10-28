package middleware

import (
	"im/internal/pkg/consts"
	"im/internal/util/ctxs"
	"im/internal/util/errs"

	"github.com/gin-gonic/gin"
	"github.com/throttled/throttled/v2"
	goredisstore "github.com/throttled/throttled/v2/store/goredisstore.v9"
	"github.com/throttled/throttled/v2/store/memstore"
)

type RateLimitMiddleware struct {
	in digIn
}

func (m *RateLimitMiddleware) RateLimitMiddleware() gin.HandlerFunc {
	var store throttled.GCRAStoreCtx
	var err error
	if m.in.Config.RateConfig.UseMemoryStore {
		store, err = memstore.NewCtx(65536)
		if err != nil {
			panic(err)
		}
	} else {
		rs, err := goredisstore.NewCtx(m.in.Rdb, "")
		if err != nil {
			panic(err)
		}
		store = rs
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerSec(m.in.Config.RateConfig.Rate),
		MaxBurst: m.in.Config.RateConfig.Burst,
	}

	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)

	rateLimitMiddleware := func(ctx *gin.Context) {
		key := consts.RATE_LIMIT_KEY + ctx.Request.Method + ":" + ctx.FullPath()
		// quantity 表示消耗配額, 如果限制每分鐘100次。消耗配額填50，則每分鐘只能通過該請求2次
		quantity := 1
		ok, _, err := rateLimiter.RateLimitCtx(ctx, key, quantity)
		if err != nil {
			m.in.Logger.Error(ctx, err)
			ctxs.SetError(ctx, errs.CommonUnknownError)
			return
		}
		if !ok {
			ctxs.SetError(ctx, errs.RequestFrequentOperationError)
			return
		}

		ctx.Next()
	}

	return rateLimitMiddleware
}