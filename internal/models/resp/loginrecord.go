package resp

type LoginRecordGet struct{}

type LoginRecordGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []LoginRecordGet `json:"data"`
}
