package resp

type UsersGet struct{}

type UsersGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []UsersGet `json:"data"`
}
