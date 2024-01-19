package po

import (
	"time"

	"im/internal/pkg/consts/enums"
)

type GroupInvitation struct {
	ID               string                      `gorm:"primarykey;column:id"` // 此邀請標示
	GroupID          string                      // 群组的唯一标识符
	InviterID        string                      // 邀请者的唯一标识符
	InviteeID        string                      // 被邀请者的唯一标识符
	InvitationStatus enums.GroupInvitationStatus // 邀请的状态
	CreatedAt        time.Time                   // 记录创建邀请的时间戳
	UpdatedAt        time.Time                   // 记录邀请最后更新的时间戳
}

func (*GroupInvitation) TableName() string {
	return "group_invitations"
}
