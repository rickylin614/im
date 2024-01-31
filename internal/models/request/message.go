package request

import (
	"time"

	"im/internal/models/po"
	"im/internal/pkg/consts/enums"

	"gorm.io/gorm"
)

type MessageGet struct{}

type MessageGetList struct {
	po.Page `gorm:"-"`
}

func (list MessageGetList) Scope(db *gorm.DB) *gorm.DB {
	// TODO write where condition
	return db
}

type MessageCreate struct {
	ID         string
	Sender     string
	Recipient  string
	MsgContent []byte
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     enums.MessageStatus
	MsgType    enums.MessageType
}

type MessageUpdate struct{}

type MessageDelete struct {
	ID string `json:"id"`
}
