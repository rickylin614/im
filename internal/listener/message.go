package listener

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/vmihailenco/msgpack"

	"im/internal/models/po"
	"im/internal/pkg/consts/topic"
)

type IMessageListener interface {
}

type MessageListener struct {
	in digIn
}

// NewMessageListener 訊息監聽器
//
// param: in 依賴
// return: IMessageListener 訊息監聽器接口
func NewMessageListener(in digIn) IMessageListener {
	return MessageListener{in: in}
}

// Start 啟動排程
//
// param: worker 啟動監聽的數量
func (m MessageListener) Start(workerNum int) {
	msg, err := m.in.Subscriber.Subscribe(context.Background(), topic.MSG)
	if err != nil {
		slog.Error("MessageListener Start Error", "error", err.Error())
	}

	for i := 0; i < workerNum; i++ {
		go func() {
			for {
				select {
				case <-m.in.Ctx.Done():
					return
				case recMsg := <-msg:
					processMsg(m.in, recMsg)
				}
			}
		}()
	}
}

// processMsg
//
// param: in 依賴
// param: msg 訂閱的訊息
func processMsg(in digIn, msg *message.Message) {
	var err error
	defer func() {
		if err == nil {
			msg.Ack()
		}
	}()

	msgModel := &po.Message{}
	if err = msgpack.Unmarshal(msg.Payload, msgModel); err != nil {
		return
	}

	// 丟進ws處理器
	ctx := context.Background()
	err = in.WsManager.SendMessage(ctx, msgModel)

	return
}