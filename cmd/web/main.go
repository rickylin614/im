package main

import (
	"im/internal/provider"
	"im/internal/server"
)

func main() {
	// initial di
	container := provider.New()

	if err := container.Invoke(server.Run(server.WEB)); err != nil {
		panic(err)
	}
}
