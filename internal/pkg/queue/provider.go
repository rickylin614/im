package queue

import (
	"go.uber.org/dig"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
)

func NewQueue(in digIn) digOut {
	p, s := NewPubSub(in)
	return digOut{
		Publisher:  p,
		Subscriber: s,
	}
}

type digIn struct {
	dig.In

	Logger logger.Logger
	Config *config.Config
	Rdb    redis.UniversalClient `optional:"true"` // 標記為可選依賴
}

type digOut struct {
	dig.Out

	Publisher  message.Publisher
	Subscriber message.Subscriber
}

func NewPubSub(in digIn) (message.Publisher, message.Subscriber) {
	switch in.Config.QueueConfig.Mode {
	case "gochannel":
		return newChannelPubSub(in)
	case "redis":
		return newRedisPublic(in), newRedisSubscriber(in)
	case "kafka":
		return newKafkaPublic(in), newKafkaSubscriber(in)
	default:
		return newChannelPubSub(in)
	}
}
