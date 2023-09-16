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

	priGroup := router.Group("/im/")
	priGroup.Use(r.in.Middle.Auth.IsLogin)
	r.setAuthRouter(priGroup)
}

// 不需要登入的API
func (r WebRouter) setPublicRouter(router *gin.RouterGroup) {
	router.GET("/ping", r.in.Handler.BaseHandler.Ping)
	router.GET("/metrics", r.in.Handler.BaseHandler.Metrics())

	// example
	router.GET("/example/:id", r.in.Handler.ExampleHandler.Get)
	router.GET("/example", r.in.Handler.ExampleHandler.GetList)
	router.POST("/example", r.in.Handler.ExampleHandler.Create)
	router.PUT("/example", r.in.Handler.ExampleHandler.Update)
	router.DELETE("/example", r.in.Handler.ExampleHandler.Delete)

	// 使用者相關
	router.POST("/users/register", r.in.Handler.UsersHandler.Create) // 註冊新用戶
	router.POST("/users/login", r.in.Handler.UsersHandler.Login)     // 用戶登錄並返回授權令牌

}

// setAuthRouter 需要登入的API
func (r WebRouter) setAuthRouter(router *gin.RouterGroup) {
	// router.DELETE("/api/users/{id}", r.in.Handler.UsersHandler.Delete)                        // 刪除指定ID的用戶
	router.POST("/users/logout", r.in.Handler.UsersHandler.Logout)                        // 用戶登出
	router.GET("/users/:id", r.in.Handler.UsersHandler.Get)                               // 獲取指定ID的用戶詳細信息
	router.PUT("/users", r.in.Handler.UsersHandler.Update)                                // 更新指定ID的用戶信息
	router.GET("/users/search", r.in.Handler.UsersHandler.GetList)                        // 根據查詢條件搜索用戶
	router.GET("/users/{id}/online-status", r.in.Handler.UsersHandler.GetOnlineStatus)    // 獲取指定用戶ID的在線狀態
	router.PUT("/users/{id}/online-status", r.in.Handler.UsersHandler.UpdateOnlineStatus) // 更新指定用戶ID的在線狀態

}
