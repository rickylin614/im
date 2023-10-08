package req

import (
	"im/internal/consts"
	"im/internal/models"
	"time"

	"gorm.io/gorm"
)

type FriendRequestsGet struct {
	ID                 string                   `gorm:"column:id"`
	SenderID           string                   `gorm:"column:sender_id"`
	ReceiverID         string                   `gorm:"column:receiver_id"`
	RequestStatus      consts.FriendReqStatus   `gorm:"column:request_status"`
	RequestStatusConds []consts.FriendReqStatus `gorm:"-"`
	CreatedAt          time.Time                `gorm:"column:created_at"`
	UpdatedAt          time.Time                `gorm:"column:updated_at"`
}

func (f FriendRequestsGet) Scope(db *gorm.DB) *gorm.DB {
	if len(f.RequestStatusConds) > 0 {
		db.Where("request_status IN ?", f.RequestStatusConds)
	}
	return db
}

type FriendRequestsGetList struct {
	UserId      string `json:"-"`                         // 用戶id
	IsSender    bool   `json:"is_sender" example:"false"` // true: 請求列表 false: 被請求列表
	models.Page `gorm:"-"`
}

func (cond FriendRequestsGetList) Scope(db *gorm.DB) *gorm.DB {
	if cond.IsSender {
		db = db.Where("sender_id = ?", cond.UserId)
	} else {
		db = db.Where("receiver_id = ?", cond.UserId)
	}
	return db
}

type FriendRequestsCreate struct {
	UserName string `json:"user_name" binding:"required" example:"user"` // 對象用戶username
}

type FriendRequestsUpdate struct {
	ID            string `json:"id"` // 請求單ID
	RequestStatus string `json:"request_status" binding:"required,oneof=accepted rejected" enums:"accepted,rejected"`
}

type FriendRequestsDelete struct {
	ID string `json:"id"` // 請求單ID
}
