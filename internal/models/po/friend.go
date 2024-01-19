package po

import (
	"time"

	"im/internal/pkg/consts/enums"
)

type Friend struct {
	ID        string             `gorm:"primarykey;id"`
	PUserID   string             `gorm:"primarykey;p_user_id"`
	PUserName string             `gorm:"p_user_name"`
	FUserID   string             `gorm:"primarykey;f_user_id"`
	FUserName string             `gorm:"f_user_name"`
	MessageId string             `gorm:"message_id"`
	Status    enums.FriendStatus `gorm:"status"`
	Mute      bool               `gorm:"mute"`
	CreatedAt time.Time          `gorm:"created_at"`
	UpdatedAt time.Time          `gorm:"updated_at"`
}

func (*Friend) TableName() string {
	return "friends"
}
