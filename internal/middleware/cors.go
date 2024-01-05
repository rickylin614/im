package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type corsMiddleware struct {
}

func (m corsMiddleware) New() gin.HandlerFunc {

	conf := cors.DefaultConfig()
	conf.AddAllowHeaders("Authorization")
	conf.AllowWebSockets = true
	conf.AllowBrowserExtensions = true
	conf.AllowAllOrigins = true

	return cors.New(conf)
}
