package models

type FriendRequests struct {
	ID string `gorm:"column:id"`
}

func (*FriendRequests) TableName() string {
	return "friend_request"
}
