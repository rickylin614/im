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
	return &Handler{}
}

type Handler struct {
}

type digIn struct {
	dig.In

	Service *service.Service
}
