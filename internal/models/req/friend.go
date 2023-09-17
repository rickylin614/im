package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type FriendGet struct{}

type FriendGetList struct {
	models.Page `gorm:"-"`
}

func (list FriendGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type FriendCreate struct{}

type FriendUpdate struct{}

type FriendDelete  struct {
	ID string `json:"id"`
}