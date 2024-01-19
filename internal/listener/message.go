package listener

import "context"

type IMessageListener interface {
}

type MessageListener struct {
	in digIn
}

func NewMessageListener(in digIn) IMessageListener {
	return MessageListener{in: in}
}

func (m MessageListener) Start(worker int) {

	m.in.Subscriber.Subscribe(context.Background(), "message")

	for i := 0; i < worker; i++ {
		go func() {

			for {
				select {
				case <-m.in.Ctx.Done():
					return

				}
			}
		}()
	}
}
