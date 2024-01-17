package request

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type GroupInvitationGet struct {
	ID               string `gorm:"primarykey;column:id"` // 此邀請標示
	GroupID          string // 群组的唯一标识符
	InviterID        string // 邀请者的唯一标识符
	InviteeID        string // 被邀请者的唯一标识符
	InvitationStatus string // 邀请的状态
}

type GroupInvitationGetList struct {
	models.Page `gorm:"-"`
}

func (list GroupInvitationGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type GroupInvitationCreate struct {
	GroupId   string // 群組ID
	InviteeId string // 被邀請者ID
}

type GroupInvitationUpdate struct{}

type GroupInvitationDelete struct {
	ID string `json:"id"`
}
