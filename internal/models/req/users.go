package req

import (
	"im/internal/models"

	"gorm.io/gorm"
)

type UsersLogin struct {
	Username string `json:"username" binding:"required,alphanum,min=6"`
	Password string `json:"password" binding:"required,alphanum,min=6"`
}

type UsersGet struct {
	Username    string `json:"username" binding:"required,alphanum,min=6"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,alphanum,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,len=10"`
}

type UsersGetList struct {
	models.Page `gorm:"-"`
}

func (list UsersGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type UsersCreate struct {
	Username    string `json:"username" binding:"required,alphanum,min=6"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,alphanum,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,len=10"`
}

type UsersUpdate struct{}

type UsersDelete struct {
	ID string `json:"id"`
}
