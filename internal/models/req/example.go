package req

import "im/internal/models"

type ExampleGet struct {
	Id          uint   `json:"id"`
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
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleDelete struct {
	Id string `json:"id,omitempty"`
}
