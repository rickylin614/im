package request

import (
	"im/internal/models/po"

	"gorm.io/gorm"
)

type MessageGet struct{}

type MessageGetList struct {
	po.Page `gorm:"-"`
}

func (list MessageGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type MessageCreate struct{}

type MessageUpdate struct{}

type MessageDelete struct {
	ID string `json:"id"`
}
