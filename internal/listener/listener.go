package listener

import (
	"context"
	"log/slog"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type IListener interface {
	Start(workerNum int)
	getTopic() string
	processMsg(in digIn, msg *message.Message)
}

type msgHandler interface {
	processMsg(in digIn, msg *message.Message)
}

type Base struct {
	in    digIn
	topic string
	msgHandler
}

// NewBase 監聽器基底
//
// param: in 依賴
// param: topic 監聽的對象主題
// return: Base
func NewBase(in digIn, topic string) Base {
	return Base{in: in, topic: topic}
}

func (l *Base) Start(workerNum int) {
	var msg <-chan *message.Message
	var err error
	for {
		msg, err = l.in.Subscriber.Subscribe(context.Background(), l.getTopic())
		if err != nil {
			slog.Error("ListenerBase Start Error", "error", err.Error(), "topic", l.getTopic())
			time.Sleep(time.Second * 30)
			continue
		}
		break
	}

	for i := 0; i < workerNum; i++ {
		go func() {
			l.in.Ctx.RunFunc(func() {
				select {
				case <-l.in.Ctx.Done():
					return
				case recMsg := <-msg:
					l.processMsg(l.in, recMsg)
				}
			})
		}()
	}
}

func (l *Base) getTopic() string {
	return l.topic
}
