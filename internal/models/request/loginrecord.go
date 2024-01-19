package request

import (
	"im/internal/models/po"

	"gorm.io/gorm"
)

type LoginRecordGet struct{}

type LoginRecordGetList struct {
	po.Page `gorm:"-"`
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
