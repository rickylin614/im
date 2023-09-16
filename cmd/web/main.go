package main

import (
	"im/internal/provider"
	"im/internal/server"
)

// @title           Im
// @version         1.0
// @description     This is a project im.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support

// @host      localhost:9000
// @BasePath  /im

// 定義api登入條件
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token

// @schemes http

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// initial di
	container := provider.New()

	if err := container.Invoke(server.Run(server.WEB)); err != nil {
		panic(err)
	}
}
