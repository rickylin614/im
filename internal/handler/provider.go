package handler

import (
	"go.uber.org/dig"

	"im/internal/service"
)

// NewHandler
//
// param: in 依賴
// return: handler 所有
func NewHandler(in digIn) *Handler {
	return &Handler{in: in,
		BaseHandler:    &baseHandler{in},
		ExampleHandler: &exampleHandler{in},
		UsersHandler:   &usersHandler{in: in},
	}
}

type Handler struct {
	in digIn

	BaseHandler    *baseHandler
	ExampleHandler *exampleHandler
	UsersHandler   *usersHandler
}

type digIn struct {
	dig.In

	Service *service.Service
}
