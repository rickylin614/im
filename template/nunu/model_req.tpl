package req

import "im/internal/models"

type {{ .FileName }}Get struct{}

type {{ .FileName }}GetList struct {
	Page models.Page `json:"page,omitempty"`
}

type {{ .FileName }}Post struct{}

type {{ .FileName }}Put struct{}

type {{ .FileName }}Delete struct{}
