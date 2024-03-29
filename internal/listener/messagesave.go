package listener

import (
	"log/slog"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/vmihailenco/msgpack/v5"

	"im/internal/models/po"
	"im/internal/pkg/consts/topic"
	"im/internal/util/ctxs"
)

type MessageSaveListener struct {
	Base
}

// NewMessageSaveListener
//
// param: in
// return: IListener
func NewMessageSaveListener(in digIn) IListener {
	m := &MessageSaveListener{Base: NewBase(in, topic.MSG_SAVE)}
	m.msgHandler = m
	return m
}

// processMsg
//
// param: in 依賴
// param: msg 訂閱的訊息
func (m MessageSaveListener) processMsg(in digIn, msg *message.Message) (err error) {
	msgModel := &po.Message{}
	if err = msgpack.Unmarshal(msg.Payload, msgModel); err != nil {
		slog.Error("messageSave processMsg Unmarshal", "error", err)
		return
	}

	if _, err = in.Service.MessageSrv.Create(ctxs.Background(), msgModel); err != nil {
		slog.Error("messageSave processMsg create", "error", err)
	}
	return
}
