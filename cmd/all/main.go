package main

import (
	"im/internal/provider"
	"im/internal/server"
)

func main() {
	// initial di
	container := provider.New()

	if err := container.Invoke(server.RunAll()); err != nil {
		panic(err)
	}
}
