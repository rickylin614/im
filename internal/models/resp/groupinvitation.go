package resp

type GroupInvitationGet struct{}

type GroupInvitationGetList struct {
	Page PageResponse         `json:"page,omitempty"`
	Data []GroupInvitationGet `json:"data"`
}
