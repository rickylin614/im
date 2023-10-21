package resp

type GroupsGet struct{}

type GroupsGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []GroupsGet `json:"data"`
}
