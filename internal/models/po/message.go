package po

import (
	"time"

	"im/internal/pkg/consts/enums"
)

type Message struct {
	ID         string              `json:"ID" bson:"ID"`                 // 訊息的唯一識別碼
	Sender     string              `json:"sender" bson:"sender"`         // 發送者的識別碼
	Recipient  string              `json:"recipient" bson:"recipient"`   // 接收者的識別碼
	MsgContent []byte              `json:"msgContent" bson:"msgContent"` // 訊息內容
	CreatedAt  time.Time           `json:"createdAt" bson:"createdAt"`   // 訊息的時間創建戳記
	UpdatedAt  time.Time           `json:"updatedAt" bson:"updatedAt"`   // 訊息的時間修改戳記
	Status     enums.MessageStatus `json:"status" bson:"status"`         // 訊息狀態
	MsgType    enums.MessageType   `json:"msgType" bson:"msgType"`       // 訊息種類
}

func (*Message) TableName() string {
	return "message"
}
