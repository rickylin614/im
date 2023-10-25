package models

import (
	"time"

	"im/internal/pkg/consts"
)

type LoginRecord struct {
	ID         int64             `gorm:"column:id"`
	Name       string            `gorm:"column:name"`
	UserID     string            `gorm:"column:user_id"`
	UserAgent  string            `gorm:"column:user_agent"`
	Ip         string            `gorm:"column:ip"`
	RemoteIp   string            `gorm:"column:remote_ip"`
	LoginState consts.LoginState `gorm:"column:login_state"`
	CreatedAt  time.Time         `gorm:"column:created_at"`
}

func (*LoginRecord) TableName() string {
	return "login_record"
}
