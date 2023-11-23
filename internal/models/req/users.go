package req

import (
	"im/internal/models"
	"im/internal/pkg/consts"

	"gorm.io/gorm"
)

type UsersLogin struct {
	Username string `json:"username" binding:"required,alphanum,min=6"` // 使用者名稱
	Password string `json:"password" binding:"required,alphanum,min=6"` // 密碼
}

type UsersGet struct {
	ID          string            `json:"-"`
	Username    string            `json:"username"`
	Nickname    string            `json:"nickname"`
	Email       string            `json:"email" `
	Password    string            `json:"password"`
	PhoneNumber string            `json:"phone_number"`
	Status      consts.UserStatus `json:"status"` // 用戶狀態
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

type UsersUpdate struct {
	ID          string `json:"id" binding:"required,uuid"`
	Username    string `json:"username" binding:"required,alphanum,min=6"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,len=10"`
}

type UsersDelete struct {
	ID string `json:"id"`
}
