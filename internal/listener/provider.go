package listener

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/dig"

	"im/internal/manager/msggateway"
	"im/internal/pkg/config"
	"im/internal/pkg/signalctx"
	"im/internal/service"
)

func NewListener(in digIn) DigOut {
	return DigOut{
		MessageListener:     NewMessageListener(in),
		MessageSaveListener: NewMessageSaveListener(in),
	}
}

type digIn struct {
	dig.In

	Ctx    *signalctx.Context
	Config *config.Config

	Publisher  message.Publisher
	Subscriber message.Subscriber
	WsManager  msggateway.IWsManager

	Service *service.Service
}

type DigOut struct {
	dig.Out

	MessageListener     IListener `name:"messageListener"`
	MessageSaveListener IListener `name:"messageSaveListener"`
}
