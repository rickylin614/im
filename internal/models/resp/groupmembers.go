package resp

type GroupMembersGet struct{}

type GroupMembersGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []GroupMembersGet `json:"data"`
}
