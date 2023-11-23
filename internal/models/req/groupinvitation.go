package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type GroupInvitationGet struct{}

type GroupInvitationGetList struct {
	models.Page `gorm:"-"`
}

func (list GroupInvitationGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type GroupInvitationCreate struct{}

type GroupInvitationUpdate struct{}

type GroupInvitationDelete struct {
	ID string `json:"id"`
}
