package request

import (
	"im/internal/models/po"

	"gorm.io/gorm"
)

type GroupsGet struct {
}

type GroupsGetList struct {
	po.Page `gorm:"-"`
}

func (list GroupsGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type GroupsCreate struct {
	GroupName   string `json:"group_name" binding:"required"` // 群組名稱
	Description string `json:"description"`                   // 描述
}

type GroupsUpdate struct{}

type GroupsDelete struct {
	ID string `json:"id"`
}
