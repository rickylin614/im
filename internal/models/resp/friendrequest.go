package resp

type FriendRequestsGet struct{}

type FriendRequestsGetList struct {
	Page PageResponse        `json:"page,omitempty"`
	Data []FriendRequestsGet `json:"data"`
}
