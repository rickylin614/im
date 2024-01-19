package po

import "time"

type GroupMembers struct {
	GroupID           string     `gorm:"primarykey;column:group_id"` // 唯一标识群组的ID
	UserID            string     `gorm:"primarykey;column:user_id"`  // 唯一标识用户的ID
	UserName          string     `gorm:"column:user_name"`           // 用户的Name
	Role              string     `gorm:"column:role"`                // 成员的角色,可以是owner、admin或member
	JoinedAt          time.Time  `gorm:"column:joined_at"`           // 成员加入群组的时间
	NicknameInGroup   *string    `gorm:"column:nickname_in_group"`   // 成员在该群组中的昵称或显示名称
	LastSeen          *time.Time `gorm:"column:last_seen"`           // 记录成员最后一次在群组中活动的时间
	MuteNotifications *bool      `gorm:"column:mute_notifications"`  // 标记成员是否静音了群组通知
	CustomPermissions *string    `gorm:"column:custom_permissions"`  // 特定成员的自定义权限设置
	MessageReadUpTo   *int64     `gorm:"column:message_read_up_to"`  // 标记成员已读消息的最后位置
	StatusInGroup     *string    `gorm:"column:status_in_group"`     // 成员在群组中的状态，如活跃、闲置或离线等
	InvitedByUserID   *string    `gorm:"column:invited_by_user_id"`  // 记录哪个用户邀请该成员加入群组
	InvitedByCode     *string    `gorm:"column:invited_by_code"`     // 记录哪个用户邀请碼進入群組
}

func (*GroupMembers) TableName() string {
	return "group_members"
}
