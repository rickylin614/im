package req

import "im/internal/models"

type ExampleGet struct {
}

type ExampleGetList struct {
	models.Page `gorm:"-"`
}

type ExamplePost struct{}

type ExamplePut struct{}

type ExampleDelete struct {
	Id string `json:"id,omitempty"`
}
