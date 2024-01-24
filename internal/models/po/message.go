package po

import (
	"time"

	"im/internal/pkg/consts/enums"
)

type Message struct {
	ID        string              // 訊息的唯一識別碼
	Sender    string              // 發送者的識別碼
	Recipient string              // 接收者的識別碼
	Content   string              // 訊息內容
	CreatedAt time.Time           // 訊息的時間創建戳記
	UpdatedAt time.Time           // 訊息的時間修改戳記
	Status    enums.MessageStatus // 訊息狀態
}

func (*Message) TableName() string {
	return "message"
}
