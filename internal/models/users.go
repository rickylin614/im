package models

import (
	"im/internal/consts"
	"time"
)

type Users struct {
	ID           string            `gorm:"column:id"`            // 用戶唯一標識符
	Username     string            `gorm:"column:username"`      // 用戶名稱
	Nickname     string            `gorm:"column:nickname"`      // 用戶暱稱
	PasswordHash string            `gorm:"column:password_hash"` // 密碼哈希值
	Email        string            `gorm:"column:email"`         // 電子郵件地址
	PhoneNumber  string            `gorm:"column:phone_number"`  // 手機號碼
	CreatedAt    time.Time         `gorm:"column:created_at"`    // 創建時間
	UpdatedAt    time.Time         `gorm:"column:updated_at"`    // 更新時間
	Status       consts.UserStatus `gorm:"column:status"`        // 用戶狀態
}

func (*Users) TableName() string {
	return "users"
}
