package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type UsersGet struct{}

type UsersGetList struct {
	models.Page `gorm:"-"`
}

func (list UsersGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type UsersCreate struct{}

type UsersUpdate struct{}

type UsersDelete  struct {
	ID string `json:"id"`
}