package resp

import "time"

type GroupMembersGet struct {
	GroupID           string     `json:"group_id,omitempty"`
	UserID            string     `json:"user_id,omitempty"`
	UserName          string     `json:"user_name,omitempty"`
	Role              string     `json:"role,omitempty"`
	JoinedAt          time.Time  `json:"joined_at"`
	NicknameInGroup   *string    `json:"nickname_in_group,omitempty"`
	LastSeen          *time.Time `json:"last_seen,omitempty"`
	MuteNotifications *bool      `json:"mute_notifications,omitempty"`
	CustomPermissions *string    `json:"custom_permissions,omitempty"`
	MessageReadUpTo   *int64     `json:"message_read_up_to,omitempty"`
	StatusInGroup     *string    `json:"status_in_group,omitempty"`
	InvitedByUserID   *string    `json:"invited_by_user_id,omitempty"`
	InvitedByCode     *string    `json:"invited_by_code,omitempty"`
}

type GroupMembersGetList struct {
	Page PageResponse      `json:"page,omitempty"`
	Data []GroupMembersGet `json:"data"`
}
