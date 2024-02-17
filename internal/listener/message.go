package listener

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/vmihailenco/msgpack/v5"

	"im/internal/models/po"
	"im/internal/pkg/consts/topic"
	"im/internal/util/uuid"
)

type MessageListener struct {
	Base
}

// NewMessageListener 訊息監聽器
//
// param: in 依賴
// return: IMessageListener 訊息監聽器接口
func NewMessageListener(in digIn) IListener {
	m := &MessageListener{Base: NewBase(in, topic.MSG)}
	m.msgHandler = m
	return m
}

// processMsg
//
// param: in 依賴
// param: msg 訂閱的訊息
func (m *MessageListener) processMsg(in digIn, msg *message.Message) {
	var err error
	defer func() {
		if err == nil {
			msg.Ack()
		}
	}()

	msgModel := &po.Message{}
	if err = msgpack.Unmarshal(msg.Payload, msgModel); err != nil {
		slog.Error("processMsg Unmarshal", "error", err)
		return
	}

	// 丟進ws處理器
	ctx := context.Background()
	err = in.WsManager.SendMessage2Client(ctx, msgModel)
	if err != nil {
		slog.Error("processMsg send message", "error", err)
		return
	}

	// 送到queue保存資料進repo
	err = in.Publisher.Publish(topic.MSG_SAVE, message.NewMessage(uuid.New(), msg.Payload))
	if err != nil {
		slog.Error("processMsg Publisher.Publish", "error", err)
		return
	}

	return
}
