package resp

type FriendGet struct{}

type FriendGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []FriendGet `json:"data"`
}
