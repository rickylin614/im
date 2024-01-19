package request

import (
	"im/internal/models/po"
	"im/internal/pkg/consts/enums"

	"time"

	"gorm.io/gorm"
)

type FriendGet struct {
	PUserID   string             `json:"-" gorm:"column:p_user_id"`                          // 主要用户的ID
	FUserID   string             `json:"f_user_id" gorm:"column:f_user_id"`                  // 好友的用户ID
	Status    enums.FriendStatus `json:"status" gorm:"column:status" enums:"active,blocked"` // 好友关系的状态（例如：active / ）
	Mute      bool               `json:"mute" gorm:"column:mute"`                            // 是否将此好友静音
	CreatedAt time.Time          `json:"created_at" gorm:"column:created_at"`                // 好友关系创建时间
	UpdatedAt time.Time          `json:"updated_at" gorm:"column:updated_at"`                // 好友关系更新时间
}

type FriendMutualGet struct {
	UserID  string `json:"-"`                             // 用户ID
	TUserId string `json:"t_user_id"  binding:"required"` // 對象用户ID
	po.Page `gorm:"-"`
}

type FriendGetList struct {
	PUserID string             `json:"-"` // 主要用户的ID
	Status  enums.FriendStatus `json:"-"`
	po.Page `gorm:"-"`
}

func (list FriendGetList) Scope(db *gorm.DB) *gorm.DB {
	db.Where("p_user_id", list.PUserID)
	db.Where("status", list.Status)
	return db
}

type FriendUpdate struct {
	FUserID string             `json:"f_user_id" binding:"required"` // 好友的用户ID
	Status  enums.FriendStatus `json:"status" `
}

type FriendDelete struct {
	ID string `json:"id"`
}
