package queue

import (
	"github.com/IBM/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

// newKafkaConfig 暫時只有queue初始化用到, 不額外產生物件
func newKafkaConfig(in digIn) *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	return config
}

func newKafkaPublic(in digIn) message.Publisher {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:               in.Config.KafkaConfig.Brokers,
			Marshaler:             kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: newKafkaConfig(in),
		},
		newWatermillZap(in),
	)
	if err != nil {
		panic(err)
	}

	return publisher
}

func newKafkaSubscriber(in digIn) message.Subscriber {
	saramaSubscriberConfig := newKafkaConfig(in)
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               in.Config.KafkaConfig.Brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         in.Config.KafkaConfig.GroupID,
		},
		newWatermillZap(in),
	)

	if err != nil {
		panic(err)
	}

	return subscriber
}
