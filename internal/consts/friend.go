package consts

type FriendReqStatus string

const (
	FriendReqStatusPending  FriendReqStatus = "pending"
	FriendReqStatusAccepted                 = "accepted"
	FriendReqStatusRejected                 = "rejected"
)

type FriendStatus string

const (
	FriendStatusActive  FriendStatus = "active"
	FriendStatusBlocked              = "blocked"
)
