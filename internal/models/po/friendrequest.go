package po

import (
	"time"

	"im/internal/pkg/consts/enums"
)

type FriendRequests struct {
	ID            string                `gorm:"primarykey;column:id"`
	SenderID      string                `gorm:"column:sender_id"`
	SenderName    string                `gorm:"column:sender_name"`
	ReceiverID    string                `gorm:"column:receiver_id"`
	ReceiverName  string                `gorm:"column:receiver_name"`
	RequestStatus enums.FriendReqStatus `gorm:"column:request_status"`
	CreatedAt     time.Time             `gorm:"column:created_at"`
	UpdatedAt     time.Time             `gorm:"column:updated_at"`
}

func (*FriendRequests) TableName() string {
	return "friend_requests"
}
