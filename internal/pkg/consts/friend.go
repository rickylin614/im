package consts

type FriendReqStatus string

const (
	FriendReqStatusPending  FriendReqStatus = "pending"
	FriendReqStatusAccepted FriendReqStatus = "accepted"
	FriendReqStatusRejected FriendReqStatus = "rejected"
)

type FriendStatus string

const (
	FriendStatusActive  FriendStatus = "active"
	FriendStatusBlocked FriendStatus = "blocked"
)
