package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type GroupsGet struct{}

type GroupsGetList struct {
	models.Page `gorm:"-"`
}

func (list GroupsGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type GroupsCreate struct{}

type GroupsUpdate struct{}

type GroupsDelete  struct {
	ID string `json:"id"`
}