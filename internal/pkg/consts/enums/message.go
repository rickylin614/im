package enums

type MessageStatus int

const (
	Normal MessageStatus = iota
	Withdrawn
	Deleted
	Hidden
)
