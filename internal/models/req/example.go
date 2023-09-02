package req

import "im/internal/models"

type ExampleGet struct {
	Id          uint   `json:"-" uri:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleGetList struct {
	models.Page `gorm:"-"`
}

type ExampleCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleUpdate struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleDelete struct {
	Id string `json:"id,omitempty"`
}
