package resp

import (
	"im/internal/consts"
	"time"
)

type FriendRequestsGet struct {
	ID            string                 `gorm:"primarykey;column:id"`
	SenderID      string                 `gorm:"column:sender_id"`
	SenderName    string                 `gorm:"column:sender_name"`
	ReceiverID    string                 `gorm:"column:receiver_id"`
	ReceiverName  string                 `gorm:"column:receiver_name"`
	RequestStatus consts.FriendReqStatus `gorm:"column:request_status"`
	CreatedAt     time.Time              `gorm:"column:created_at"`
}

type FriendRequestsGetList struct {
	Page PageResponse        `json:"page,omitempty"`
	Data []FriendRequestsGet `json:"data"`
}
