package handler

import (
	"go.uber.org/dig"

	"im/internal/pkg/logger"
	"im/internal/service"
)

// NewHandler
//
// param: in 依賴
// return: handler 所有
func NewHandler(in digIn) *Handler {
	return &Handler{in: in,
		BaseHandler:           &baseHandler{in},
		ExampleHandler:        &exampleHandler{in},
		UsersHandler:          &usersHandler{in: in},
		LoginRecordHandler:    &loginRecordHandler{in: in},
		FriendHandler:         &friendHandler{in: in},
		FriendRequestsHandler: &FriendRequestsHandler{in: in},
	}
}

type Handler struct {
	in digIn

	Logger                *logger.Logger
	BaseHandler           *baseHandler
	ExampleHandler        *exampleHandler
	UsersHandler          *usersHandler
	LoginRecordHandler    *loginRecordHandler
	FriendHandler         *friendHandler
	FriendRequestsHandler *FriendRequestsHandler
}

type digIn struct {
	dig.In

	Service *service.Service
}
