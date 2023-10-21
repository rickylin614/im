package models

type Groups struct {
	ID string `gorm:"column:id"`
}

func (*Groups) TableName() string {
	return "groups"
}