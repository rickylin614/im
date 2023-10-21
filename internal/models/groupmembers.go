package models

type GroupMembers struct {
	ID string `gorm:"column:id"`
}

func (*GroupMembers) TableName() string {
	return "group_members"
}