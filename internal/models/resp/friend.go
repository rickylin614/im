package resp

import (
	"im/internal/pkg/consts/enums"
)

type FriendGet struct {
	ID        string             `json:"id"`
	PUserID   string             `json:"p_user_id"`
	PUserName string             `json:"p_user_name"`
	FUserID   string             `json:"f_user_id"`
	FUserName string             `json:"f_user_name"`
	Status    enums.FriendStatus `json:"status"`
	Mute      bool               `json:"mute"`
}

type FriendMutualGet struct {
	PUserID   string `json:"user_id"`
	PUserName string `json:"user_name"`
}

type FriendGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []FriendGet  `json:"data"`
}

type FriendMutualList struct {
	Page PageResponse      `json:"page,omitempty"`
	Data []FriendMutualGet `json:"data"`
}
