package response

type MessageGet struct{}

type MessageGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []MessageGet `json:"data"`
}
