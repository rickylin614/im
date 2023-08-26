package resp

import "im/internal/models"

type ExampleGet struct{}

type ExampleGetList struct {
	Page models.Page      `json:"page,omitempty"`
	Data []models.Example `json:"data"`
}

type ExamplePost struct{}

type ExamplePut struct{}

type ExampleDelete struct{}
