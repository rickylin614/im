package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type RouteCacheGet struct{}

type RouteCacheGetList struct {
	models.Page `gorm:"-"`
}

func (list RouteCacheGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type RouteCacheCreate struct{}

type RouteCacheUpdate struct{}

type RouteCacheDelete struct {
	ID string `json:"id"`
}
