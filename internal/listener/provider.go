package listener

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/dig"

	"im/internal/pkg/signalctx"
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
}

type digOut struct {
	dig.Out

	MessageListener IMessageListener
}
