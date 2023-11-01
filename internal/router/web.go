package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type WebRouter struct {
	in digIn
}

func (r WebRouter) SetRouter(router *gin.Engine) {
	//	註冊pprof
	pprof.Register(router, "/debug/pprof")
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

	pubGroup := router.Group("/im/")
	pubGroup.Use(r.in.Middle.Rate.RateLimitMiddleware())
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
	// Users
	router.POST("/users/logout", r.in.Handler.UsersHandler.Logout)                        // 用戶登出
	router.GET("/users/:id", r.in.Handler.UsersHandler.Get)                               // 獲取指定ID的用戶詳細信息
	router.PUT("/users", r.in.Handler.UsersHandler.Update)                                // 更新指定ID的用戶信息
	router.GET("/users/search", r.in.Handler.UsersHandler.GetList)                        // 根據查詢條件搜索用戶
	router.GET("/users/{id}/online-status", r.in.Handler.UsersHandler.GetOnlineStatus)    // 獲取指定用戶ID的在線狀態
	router.PUT("/users/{id}/online-status", r.in.Handler.UsersHandler.UpdateOnlineStatus) // 更新指定用戶ID的在線狀態

	// friend
	router.GET("/friend", r.in.Handler.FriendHandler.GetFriends)                // 獲取用戶的好友列表
	router.PUT("/friend", r.in.Handler.FriendHandler.UpdateFriendStatus)        // 指定用戶ID封鎖或取消封鎖指定好友ID
	router.DELETE("/friend", r.in.Handler.FriendHandler.DeleteFriend)           // 刪除與指定用戶的好友關係
	router.GET("/friend/blocked", r.in.Handler.FriendHandler.GetBlockedFriends) // 獲取指定用戶ID的已封鎖好友列表
	router.GET("/friend/mutual", r.in.Handler.FriendHandler.GetMutualFriends)   // 獲取指定用戶ID與另一指定用戶ID的共同好友列表

	// TODO nunu router @Router 部分改為-做為切割版本
	// friend-requests
	router.GET("/friend-requests", r.in.Handler.FriendRequestsHandler.GetList) // 獲取指定用戶ID收到的好友請求列表
	router.POST("/friend-requests", r.in.Handler.FriendRequestsHandler.Create) // 向指定用戶ID發送好友請求
	router.PUT("/friend-requests", r.in.Handler.FriendRequestsHandler.Update)  // 指定用戶ID接受或拒絕來自requester-id的好友請求

	router.GET("/group/:id", r.in.Middle.Cache.RouteCacheMiddleware, r.in.Handler.GroupsHandler.Get)
	router.GET("/group", r.in.Handler.GroupsHandler.GetList)
	router.POST("/group", r.in.Handler.GroupsHandler.Create)
	router.PUT("/group", r.in.Handler.GroupsHandler.Update)
}
