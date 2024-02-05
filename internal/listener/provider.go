package listener

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/dig"

	"im/internal/manager/msggateway"
	"im/internal/pkg/signalctx"
	"im/internal/service"
)

func NewListener(in digIn) digOut {
	return digOut{
		MessageListener:     NewMessageListener(in),
		MessageSaveListener: NewMessageSaveListener(in),
	}
}

type digIn struct {
	Ctx        *signalctx.Context
	Publisher  message.Publisher
	Subscriber message.Subscriber
	WsManager  msggateway.IWsManager

	Service service.Service
}

type digOut struct {
	dig.Out

	MessageListener     IListener `name:"messageListener"`
	MessageSaveListener IListener `name:"messageSaveListener"`
}
