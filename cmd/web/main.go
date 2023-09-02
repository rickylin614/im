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

// @host      127.0.0.1:9000
// @BasePath  /im

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// initial di
	container := provider.New()

	if err := container.Invoke(server.Run(server.WEB)); err != nil {
		panic(err)
	}
}
