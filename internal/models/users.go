package models

import (
	"im/internal/consts"
	"time"
)

type Users struct {
	ID           string            `gorm:"column:id" json:"i"`            // 用戶唯一標識符
	Username     string            `gorm:"column:username" json:"u"`      // 用戶名稱
	Nickname     string            `gorm:"column:nickname" json:"n"`      // 用戶暱稱
	PasswordHash string            `gorm:"column:password_hash" json:"p"` // 密碼哈希值
	Email        string            `gorm:"column:email" json:"e"`         // 電子郵件地址
	PhoneNumber  string            `gorm:"column:phone_number" json:"ph"` // 手機號碼
	CreatedAt    time.Time         `gorm:"column:created_at" json:"ct"`   // 創建時間
	UpdatedAt    time.Time         `gorm:"column:updated_at" json:"ut"`   // 更新時間
	Status       consts.UserStatus `gorm:"column:status" json:"s"`        // 用戶狀態
}

func (*Users) TableName() string {
	return "users"
}
