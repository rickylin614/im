package listener

import (
	"context"
	"fmt"
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
	for i := 0; i < workerNum; i++ {
		go func(work int) {
			msg, err := l.in.Subscriber.Subscribe(context.Background(), l.getTopic())
			if err != nil {
				slog.Error("create worker fail", "topic", l.topic, "workerNumber", work)
				return
			}

			defer func() {
				fmt.Println("listen end")
			}()

			l.in.Ctx.RunFunc(func() {
				select {
				case <-l.in.Ctx.Done():
					return
				case recMsg := <-msg:
					if err := l.processMsg(l.in, recMsg); err == nil {
						recMsg.Ack()
					} else {
						// 避免服務阻塞 延遲丟出失敗的請求
						go func() {
							time.Sleep(time.Second * 10)
							recMsg.Nack()
						}()
					}
				}
			})
		}(i)
	}
}

func (l *Base) getTopic() string {
	return l.topic
}
