package queue

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func newChannelPubSub(in digIn) *gochannel.GoChannel {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{
			OutputChannelBuffer: 1024,
		},
		newWatermillZap(in))
	return pubSub
}
