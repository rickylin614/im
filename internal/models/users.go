package models

type Users struct {
	ID string `gorm:"column:id"`
}

func (*Users) TableName() string {
	return "users"
}