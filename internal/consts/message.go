package consts

type MessageStatus int

const (
	Normal MessageStatus = iota
	Withdrawn
	Deleted
	Hidden
)
