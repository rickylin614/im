package resp

import "im/internal/models"

type ExampleGet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleGetList struct {
	Page models.Page      `json:"page,omitempty"`
	Data []models.Example `json:"data"`
}

type ExamplePost struct {
	Id uint `json:"id"`
}
