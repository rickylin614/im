package router

import (
	"github.com/gin-gonic/gin"
)

type WebRouter struct {
	in digIn
}

func (r WebRouter) SetRouter(router *gin.Engine) {
	// set middleware
	if r.in.Config.GinConfig.DebugMode {
		gin.SetMode(gin.DebugMode)
	}
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	pubGroup := router.Group("/im/")
	r.setPublicRouter(pubGroup)

	priGroup := router.Group("/auth/")
	r.setAuthRouter(priGroup)

	// router.Use(
	// r.in.Middleware.RecoveryMiddleware.Recovery,
	// r.in.Middleware.TraceMiddleware.WithTraceLogger,
	// )
	// r.setBaseRoute(router)
	// router.Use(r.in.Middleware.ResponseMiddleware.ResponseFormat)

	// set router
	// baseGroup := router.Group("central-platform")
	// r.setPublicRouter(baseGroup)
	// r.setAuthRouter(baseGroup)
}

// 不需要登入的API
func (r WebRouter) setPublicRouter(router *gin.RouterGroup) {

}

// setAuthRouter 需要登入的API
func (r WebRouter) setAuthRouter(router *gin.RouterGroup) {

}
