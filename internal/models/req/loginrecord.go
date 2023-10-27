package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type LoginRecordGet struct{}

type LoginRecordGetList struct {
	models.Page `gorm:"-"`
}

func (list LoginRecordGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type LoginRecordCreate struct{}

type LoginRecordUpdate struct{}

type LoginRecordDelete struct {
	ID string `json:"id"`
}
