package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"

	"im/internal/models"
	"im/internal/pkg/consts/rediskey"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
)

type CacheMiddleware struct {
	in    digIn
	group singleflight.Group
}

// RouteCacheMiddleware 路由緩存
//
//	參數限定只能有Uri。避免過多差異性。
//	使用gin.ShouldBindUri
func (m *CacheMiddleware) RouteCacheMiddleware(ctx *gin.Context) {
	key := rediskey.ROUTE_CACHE_KEY + ctx.Request.Method + ":" + ctx.FullPath()

	// 確認緩存
	cache, err := m.in.Service.RouteCacheSrv.Get(ctx, &models.RouteCacheGet{
		RouteCacheKey: key,
	})

	// 有緩存，返回緩存資料
	if err == nil && len(cache.Bytes()) > 0 {
		ctx.Data(cache.Status, "application/json; charset=utf-8", cache.Body)
		ctx.Abort()
		return
	}

	// 使用singleflight避免缓存穿透
	data, err, shared := m.group.Do(key, func() (interface{}, error) {
		// 創建一個用於存儲響應數據的buffer
		buff := new(bytes.Buffer)
		// 包裝原始的ResponseWriter
		wrappedWriter := &responseBodyWriter{body: buff, ResponseWriter: ctx.Writer}
		// 替換原始的ResponseWriter
		ctx.Writer = wrappedWriter
		ctx.Next()

		// 取得Route TTL時間
		ttl := 5 * time.Minute
		if customTTL, exists := ctx.Get(rediskey.ROUTE_TTL_KEY); exists {
			if t, ok := customTTL.(time.Duration); ok {
				ttl = t
			}
		}

		// 取得response資料
		responseData := wrappedWriter.body.Bytes()
		status := wrappedWriter.Status()

		routeCache := models.NewRouteCache(status, responseData)
		// 保存緩存資料
		err := m.in.Service.RouteCacheSrv.Set(ctx, &models.RouteCacheSet{
			RouteCacheKey:  key,
			RouteCacheData: routeCache,
			TTL:            ttl,
		})

		if err != nil {
			m.in.Logger.Error(ctx, err)
		}

		// 获取response数据
		return routeCache, nil
	})

	if v, ok := data.(*models.RouteCache); err == nil && ok {
		if shared {
			ctx.Data(v.Status, "application/json; charset=utf-8", v.Body)
			ctx.Abort()
		}
	} else {
		if err != nil {
			m.in.Logger.Error(ctx, err)
		}
		ctxs.SetError(ctx, errs.CommonUnknownError)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseBodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
