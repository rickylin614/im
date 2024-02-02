package handler

import (
	"go.uber.org/dig"

	"im/internal/manager/msggateway"
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/prom"
	"im/internal/service"
)

// NewWebHandler
//
// param: in 依賴
// return: handler 所有
func NewWebHandler(in webDigIn) *WebHandler {
	return &WebHandler{in: in,
		BaseHandler:            &baseHandler{in},
		ExampleHandler:         &exampleHandler{in},
		UsersHandler:           &usersHandler{in: in},
		LoginRecordHandler:     &loginRecordHandler{in: in},
		FriendHandler:          &friendHandler{in: in},
		FriendRequestsHandler:  &FriendRequestsHandler{in: in},
		GroupsHandler:          &groupsHandler{in: in},
		GroupMembersHandler:    &groupMembersHandler{in: in},
		GroupInvitationHandler: &groupInvitationHandler{in: in},
		MessageHandler:         &messageHandler{in: in},
	}
}

type WebHandler struct {
	in webDigIn

	Logger                 logger.Logger
	BaseHandler            *baseHandler
	ExampleHandler         *exampleHandler
	UsersHandler           *usersHandler
	LoginRecordHandler     *loginRecordHandler
	FriendHandler          *friendHandler
	FriendRequestsHandler  *FriendRequestsHandler
	GroupsHandler          *groupsHandler
	GroupMembersHandler    *groupMembersHandler
	GroupInvitationHandler *groupInvitationHandler
	MessageHandler         *messageHandler
}

type webDigIn struct {
	dig.In

	Service *service.Service
	Prom    *prom.Manager
}

// NewWebSocketHandler
//
// param: in 依賴
// return: WebSocketHandler ws用的處理器
func NewWebSocketHandler(in wsDigIn) *WebSocketHandler {
	return &WebSocketHandler{
		WsHandler: &wsHandler{in: in},
	}
}

type WebSocketHandler struct {
	WsHandler *wsHandler
}

type wsDigIn struct {
	dig.In

	WsManager msggateway.IWsManager
	Service   *service.Service
	Config    *config.Config
}
