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
}

// 不需要登入的API
func (r WebRouter) setPublicRouter(router *gin.RouterGroup) {
	router.GET("/ping", r.in.Handler.BaseHandler.Ping)
	router.GET("/metrics", r.in.Handler.BaseHandler.Metrics())
}

// setAuthRouter 需要登入的API
func (r WebRouter) setAuthRouter(router *gin.RouterGroup) {

}
