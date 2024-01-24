package listener

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/dig"

	"im/internal/pkg/signalctx"
	"im/internal/service"
)

func NewListener(in digIn) digOut {
	return digOut{
		MessageListener: NewMessageListener(in),
	}
}

type digIn struct {
	Ctx        *signalctx.Context
	Publisher  message.Publisher
	Subscriber message.Subscriber

	Service service.Service
}

type digOut struct {
	dig.Out

	MessageListener IMessageListener
}
