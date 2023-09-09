package models

import (
	"time"
)

type Example struct {
	ID          string    `gorm:"primarykey;column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (*Example) TableName() string {
	return "example"
}
