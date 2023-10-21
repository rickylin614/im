package resp

type GroupsGet struct {
	ID          string `json:"id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
}

type GroupsGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []GroupsGet  `json:"data"`
}
