package queue

import (
	"go.uber.org/dig"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
)

func NewQueue(in digIn) digOut {
	return digOut{
		Publisher:  NewPublisher(in),
		Subscriber: NewSubscriber(in),
	}
}

type digIn struct {
	dig.In

	Logger logger.Logger
	Config *config.Config
	rdb    redis.UniversalClient `optional:"true"` // 標記為可選依賴
}

type digOut struct {
	dig.Out

	Publisher  message.Publisher
	Subscriber message.Subscriber
}

func NewPublisher(in digIn) message.Publisher {
	switch in.Config.QueueConfig.Mode {
	case "gochannel":
		return newChannelPubSub(in)
	case "redis":
		return newRedisPublic(in)
	case "kafka":
		return newKafkaPublic(in)
	default:
		return newChannelPubSub(in)
	}
}

func NewSubscriber(in digIn) message.Subscriber {
	switch in.Config.QueueConfig.Mode {
	case "gochannel":
		return newChannelPubSub(in)
	case "redis":
		return newRedisSubscriber(in)
	case "kafka":
		return newKafkaSubscriber(in)
	default:
		return newChannelPubSub(in)
	}
}
