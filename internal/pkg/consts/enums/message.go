package enums

type MessageStatus int

const (
	Normal MessageStatus = iota
	Withdrawn
	Deleted
	Hidden
)

type MessageType int

const (
	SingleChatType MessageType = iota + 1
	GroupChatType
	SuperGroupChatType
	NotificationChatType
)
