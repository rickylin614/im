package req

import "im/internal/models"

type {{ .FileName }}Get struct{}

type {{ .FileName }}GetList struct {
	Page models.Page `json:"page,omitempty"`
}

func (list {{ .FileName }}GetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type {{ .FileName }}Post struct{}

type {{ .FileName }}Put struct{}

type {{ .FileName }}Delete struct{}
