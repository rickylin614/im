package consts

type FriendshipStatus string

const (
	FriendshipStatusPending  FriendshipStatus = "pending"
	FriendshipStatusAccepted                  = "accepted"
	FriendshipStatusRejected                  = "rejected"
	FriendshipStatusBlocked                   = "blocked"
)
