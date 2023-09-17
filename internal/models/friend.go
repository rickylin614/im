package models

type Friend struct {
	ID string `gorm:"column:id"`
}

func (*Friend) TableName() string {
	return "friend"
}