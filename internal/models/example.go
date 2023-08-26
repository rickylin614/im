package models

type Example struct{}

func (*Example) TableName() string {
	return "example"
}
