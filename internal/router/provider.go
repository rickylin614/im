package router

import (
	"im/internal/handler"
	"im/internal/middleware"
	"im/internal/pkg/config"

	"go.uber.org/dig"
)

func NewRouter(in webDigIn) *WebRouter {
	return &WebRouter{in: in}
}

func NewWsRouter(in wsDigIn) *WsRouter {
	return &WsRouter{in: in}
}

type webDigIn struct {
	dig.In

	Config  *config.Config
	Handler *handler.WebHandler
	Middle  *middleware.Middleware
}

type wsDigIn struct {
	dig.In

	Config     *config.Config
	Handler    *handler.WebSocketHandler
	Middleware *middleware.Middleware
}
