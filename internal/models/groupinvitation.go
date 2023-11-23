package models

type GroupInvitation struct {
	ID string `gorm:"column:id"`
}

func (*GroupInvitation) TableName() string {
	return "group_invitations"
}
