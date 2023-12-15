package main

import (
	"im/internal/provider"
	"im/internal/server"

	_ "go.uber.org/automaxprocs"
)

func main() {
	// initial di
	container := provider.New()

	if err := container.Invoke(server.RunWs()); err != nil {
		panic(err)
	}
}
