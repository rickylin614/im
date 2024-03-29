package request

import (
	"{{ .ProjectName }}/internal/models/po"

	"gorm.io/gorm"
)

type {{ .FileName }}Get struct{}

type {{ .FileName }}GetList struct {
	po.Page `gorm:"-"`
}

func (list {{ .FileName }}GetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type {{ .FileName }}Create struct{}

type {{ .FileName }}Update struct{}

type {{ .FileName }}Delete  struct {
	ID string `json:"id"`
}