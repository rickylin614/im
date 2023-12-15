package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type WsRouter struct {
	in wsDigIn
}

func (r WsRouter) SetRouter(router *gin.Engine) {
	// 註冊pprof
	pprof.Register(router, "/debug/pprof")
	router.Use(cors.Default())
	// check mode
	if r.in.Config.GinConfig.DebugMode {
		gin.SetMode(gin.DebugMode)
		router.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO define recovery middleware
	router.Use(
		gin.Recovery(),
	)

	router.GET("/connect", r.in.Handler.WsHandler.Connect)
}
