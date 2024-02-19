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
	processMsg(in digIn, msg *message.Message) (err error)
}

type msgHandler interface {
	processMsg(in digIn, msg *message.Message) (err error)
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
	// gochannel 強制只能使用一個監聽, 避免資源競爭
	if l.in.Config.QueueConfig.Mode == "gochannel" {
		workerNum = 1
	}

	for i := 0; i < workerNum; i++ {
		go func(work int) {
			msgQueue, err := l.in.Subscriber.Subscribe(context.Background(), l.getTopic())
			if err != nil {
				slog.Error("create worker fail", "topic", l.topic, "workerNumber", work)
				return
			}

			l.in.Ctx.RunFunc(func() {
				select {
				case <-l.in.Ctx.Done():
					return
				case recMsg := <-msgQueue:
					if err := l.processMsg(l.in, recMsg); err == nil {
						recMsg.Ack()
					} else {
						// 避免服務阻塞 延遲丟出失敗的請求
						go func(msg <-chan *message.Message, recMsg *message.Message) {
							time.Sleep(time.Second * 10)
							recMsg.Nack()
						}(msgQueue, recMsg)
						msgQueue, _ = l.in.Subscriber.Subscribe(context.Background(), l.getTopic())
					}
				}
			})
		}(i)
	}
}

func (l *Base) getTopic() string {
	return l.topic
}
