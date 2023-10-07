package req

import (
	"im/internal/models"
	"time"

	"gorm.io/gorm"
)

type FriendGet struct {
	PUserID   string    `json:"-" gorm:"column:p_user_id"`           // 主要用户的ID
	FUserID   string    `json:"f_user_id" gorm:"column:f_user_id"`   // 好友的用户ID
	Status    string    `json:"status" gorm:"column:status"`         // 好友关系的状态（例如：pending，accepted，rejected）
	Mute      bool      `json:"mute" gorm:"column:mute"`             // 是否将此好友静音
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 好友关系创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 好友关系更新时间
}

type FriendGetList struct {
	models.Page `gorm:"-"`
}

func (list FriendGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type FriendCreate struct{}

type FriendUpdate struct{}

type FriendDelete struct {
	ID string `json:"id"`
}
