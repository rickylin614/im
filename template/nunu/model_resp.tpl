package resp

import "im/internal/models"

type {{ .FileName }}Get struct{}

type {{ .FileName }}GetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []{{ .FileName }}Get `json:"data"`
}
