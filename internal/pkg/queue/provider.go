package queue

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/dig"

	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
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

func kafkaExample() {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func channelExample() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false))

	message, err := pubSub.Subscribe(context.Background(), "example.Topic")
	if err != nil {
		panic(err)
	}

	go process(message)
	publishMessages(pubSub)
}

func redisExample() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client: client,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	subscriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client: client,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Microsecond * 250)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		time.Sleep(time.Second * 2)
		msg.Ack()
	}
}
