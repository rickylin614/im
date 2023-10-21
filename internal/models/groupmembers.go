package models

import "time"

type GroupMembers struct {
	GroupID  string    `gorm:"primarykey;column:group_id"`
	UserID   string    `gorm:"primarykey;column:user_id"`
	Role     string    `gorm:"column:role"`
	JoinedAt time.Time `gorm:"column:joined_at"`
}

func (*GroupMembers) TableName() string {
	return "group_members"
}
