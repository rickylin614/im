package request

import (
	"gorm.io/gorm"

	"im/internal/pkg/consts/enums"
)

type GroupMembersGet struct {
	UserId string // 用戶ID
}

type GroupMembersGetList struct {
	Id            string           `uri:"id"`               // 群組ID
	Role          *enums.GroupRole `form:"role"`            // 角色
	StatusInGroup *string          `form:"status_in_group"` // 群組內狀態
}

func (list GroupMembersGetList) Scope(db *gorm.DB) *gorm.DB {
	db.Where("group_id", list.Id)
	if list.Role != nil {
		db = db.Where("role", list.Role)
	}
	if list.StatusInGroup != nil {
		db = db.Where("status_in_group", list.StatusInGroup)
	}
	return db
}

type GroupMembersCreate struct{}

type GroupMembersUpdate struct{}

type GroupMembersDelete struct {
	ID string `json:"id"`
}
