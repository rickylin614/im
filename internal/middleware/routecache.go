package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"

	"im/internal/models/req"
	"im/internal/pkg/consts"
	"im/internal/util/ctxs"
	"im/internal/util/errs"
)

type CacheMiddleware struct {
	in    digIn
	group singleflight.Group
}

func (m *CacheMiddleware) RouteCacheMiddleware(ctx *gin.Context) {

	key := consts.ROUTE_CACHE_KEY + ctx.Request.Method + ":" + ctx.FullPath()

	// 確認緩存
	cache, err := m.in.Service.RouteCacheSrv.Get(ctx, &req.RouteCacheGet{
		RouteCacheKey: key,
	})

	// 有緩存，返回緩存資料
	if err == nil && len(cache) > 0 {
		ctx.String(http.StatusOK, cache)
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
		if customTTL, exists := ctx.Get(consts.ROUTE_TTL_KEY); exists {
			if t, ok := customTTL.(time.Duration); ok {
				ttl = t
			}
		}

		// 取得response資料
		responseData := wrappedWriter.body.String()

		// 保存緩存資料
		err := m.in.Service.RouteCacheSrv.Set(ctx, &req.RouteCacheSet{
			RouteCacheKey:  key,
			RouteCacheData: responseData,
			TTL:            ttl,
		})

		if err != nil {
			m.in.Logger.Error(ctx, err)
		}

		// 获取response数据
		return responseData, nil
	})

	if shared {
		ctx.Abort()
	}

	if v, ok := data.(string); err == nil && ok {
		ctx.String(http.StatusOK, v)
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
