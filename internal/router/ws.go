package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type WsRouter struct {
	in wsDigIn
}

func (r WsRouter) SetRouter(router *gin.Engine) {
	// 註冊pprof
	pprof.Register(router, "/debug/pprof")
	router.Use(r.in.Middleware.Cors.New())
	// check mode
	if r.in.Config.GinConfig.DebugMode {
		gin.SetMode(gin.DebugMode)
		router.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO define recovery Middleware
	router.Use(
		gin.Recovery(),
	)

	router.GET("/connect", r.in.Handler.WsHandler.Connect)
	router.GET("/online", r.in.Handler.WsHandler.CheckOnlineMembers)
	router.GET("/connect2", r.in.Handler.WsHandler.WsTest)
}
