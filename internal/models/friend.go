package models

import "time"

type Friend struct {
	ID        string    `gorm:"primarykey;id"`
	PUserID   string    `gorm:"p_user_id"`
	FUserID   string    `gorm:"f_user_id"`
	Status    string    `gorm:"status"`
	Mute      bool      `gorm:"mute"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (*Friend) TableName() string {
	return "friends"
}
