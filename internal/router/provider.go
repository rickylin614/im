package router

import (
	"im/internal/pkg/config"

	"go.uber.org/dig"
)

func NewRouter(in digIn) Router {
	return Router{
		WebRouter: &WebRouter{in: in},
	}
}

type Router struct {
	dig.Out

	WebRouter *WebRouter
}

type digIn struct {
	dig.In

	Config *config.Config
}
