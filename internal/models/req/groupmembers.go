package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type GroupMembersGet struct{}

type GroupMembersGetList struct {
	models.Page `gorm:"-"`
}

func (list GroupMembersGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type GroupMembersCreate struct{}

type GroupMembersUpdate struct{}

type GroupMembersDelete  struct {
	ID string `json:"id"`
}