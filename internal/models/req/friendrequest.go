package req

import (
	"im/internal/consts"
	"im/internal/models"
	"time"

	"gorm.io/gorm"
)

type FriendRequestsGet struct {
	ID            string                 `gorm:"id"`
	SenderID      string                 `gorm:"sender_id"`
	ReceiverID    string                 `gorm:"receiver_id"`
	RequestStatus consts.FriendReqStatus `gorm:"request_status"`
	CreatedAt     time.Time              `gorm:"created_at"`
	UpdatedAt     time.Time              `gorm:"updated_at"`
}

type FriendRequestsGetList struct {
	UserId      string // 用戶id
	models.Page `gorm:"-"`
}

func (list FriendRequestsGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type FriendRequestsCreate struct {
	UserName string `json:"user_name" binding:"required" example:"user"` // 對象用戶username
}

type FriendRequestsUpdate struct{}

type FriendRequestsDelete struct {
	ID string `json:"id"`
}
