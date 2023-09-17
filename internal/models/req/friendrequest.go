package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type FriendRequestsGet struct{}

type FriendRequestsGetList struct {
	models.Page `gorm:"-"`
}

func (list FriendRequestsGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type FriendRequestsCreate struct{}

type FriendRequestsUpdate struct{}

type FriendRequestsDelete struct {
	ID string `json:"id"`
}
