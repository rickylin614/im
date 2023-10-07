package models

import (
	"im/internal/consts"
	"time"
)

type FriendRequests struct {
	ID            string                 `gorm:"column:id"`
	SenderID      string                 `gorm:"column:sender_id"`
	ReceiverID    string                 `gorm:"column:receiver_id"`
	RequestStatus consts.FriendReqStatus `gorm:"column:request_status"`
	CreatedAt     time.Time              `gorm:"column:created_at"`
	UpdatedAt     time.Time              `gorm:"column:updated_at"`
}

func (*FriendRequests) TableName() string {
	return "friend_requests"
}
