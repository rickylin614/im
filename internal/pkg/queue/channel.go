package queue

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func newChannelPubSub(in digIn) (pub *gochannel.GoChannel, sub *gochannel.GoChannel) {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{
			OutputChannelBuffer:            1024,
			Persistent:                     false,
			BlockPublishUntilSubscriberAck: false,
		},
		newWatermillZap(in))
	return pubSub, pubSub
}
