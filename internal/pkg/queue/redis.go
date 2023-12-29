package queue

import (
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
)

func newRedisPublic(in digIn) message.Publisher {
	client := in.Rdb

	public, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client: client,
		},
		newWatermillZap(in),
	)
	if err != nil {
		panic(err)
	}

	return public
}

func newRedisSubscriber(in digIn) message.Subscriber {
	client := in.Rdb

	public, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client: client,
		},
		newWatermillZap(in),
	)
	if err != nil {
		panic(err)
	}

	return public
}
