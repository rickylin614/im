package middleware

import (
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
	data, err, _ := m.group.Do(key, func() (interface{}, error) {
		// 沒有緩存時執行的原始方法
		capture := NewResponseCapture(ctx.Writer)
		ctx.Writer = capture
		ctx.Next()

		// 取得Route TTL時間
		ttl := 5 * time.Minute
		if customTTL, exists := ctx.Get(consts.ROUTE_TTL_KEY); exists {
			if t, ok := customTTL.(time.Duration); ok {
				ttl = t
			}
		}

		// 取得response資料
		responseData := string(capture.capturedData)

		// 保存緩存資料
		m.in.Service.RouteCacheSrv.Set(ctx, &req.RouteCacheSet{
			RouteCacheKey:  key,
			RouteCacheData: responseData,
			TTL:            ttl,
		})

		// 获取response数据
		return responseData, nil
	})

	if v, ok := data.(string); err != nil && ok {
		ctx.String(http.StatusOK, v)
	} else {
		m.in.Logger.Error(ctx, err)
		ctxs.SetError(ctx, errs.CommonUnknownError)
	}
}

// ResponseCapture 自定義Write 用於抓取response資料
type ResponseCapture struct {
	gin.ResponseWriter
	capturedData []byte
}

func (r *ResponseCapture) Write(b []byte) (int, error) {
	r.capturedData = append(r.capturedData, b...)
	return r.ResponseWriter.Write(b)
}

func NewResponseCapture(w gin.ResponseWriter) *ResponseCapture {
	return &ResponseCapture{
		ResponseWriter: w,
		capturedData:   make([]byte, 0),
	}
}
