// 紀錄IM相關資料tables
package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

type UserProfile struct {
	UserID      uint `gorm:"primaryKey"`
	AvatarURL   string
	DisplayName string
	Bio         string
	Location    string
	Website     string
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
}

type Friend struct {
	UserID1          uint      `gorm:"primaryKey"`
	UserID2          uint      `gorm:"primaryKey"`
	FriendshipStatus string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	User1            User      `gorm:"foreignKey:UserID1"`
	User2            User      `gorm:"foreignKey:UserID2"`
}

type FriendRequest struct {
	ID            uint      `gorm:"primaryKey"`
	SenderID      uint      `gorm:"not null"`
	ReceiverID    uint      `gorm:"not null"`
	RequestStatus string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
	Sender        User      `gorm:"foreignKey:SenderID"`
	Receiver      User      `gorm:"foreignKey:ReceiverID"`
}

type Group struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"unique;not null"`
	Description  string
	GroupOwnerID uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
	GroupOwner   User      `gorm:"foreignKey:GroupOwnerID"`
}

type GroupMember struct {
	GroupID  uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"primaryKey"`
	Role     string    `gorm:"not null"`
	JoinedAt time.Time `gorm:"not null"`
	Group    Group     `gorm:"foreignKey:GroupID"`
	User     User      `gorm:"foreignKey:UserID"`
}

type GroupInvitation struct {
	ID               uint      `gorm:"primaryKey"`
	GroupID          uint      `gorm:"not null"`
	InviterID        uint      `gorm:"not null"`
	InviteeID        uint      `gorm:"not null"`
	InvitationStatus string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	Group            Group     `gorm:"foreignKey:GroupID"`
	Inviter          User      `gorm:"foreignKey:InviterID"`
	Invitee          User      `gorm:"foreignKey:InviteeID"`
}

type GroupRequest struct {
	ID            uint      `gorm:"primaryKey"`
	GroupID       uint      `gorm:"not null"`
	RequesterID   uint      `gorm:"not null"`
	RequestStatus string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
	Group         Group     `gorm:"foreignKey:GroupID"`
	Requester     User      `gorm:"foreignKey:RequesterID"`
}

type PrivateMessage struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"`
	ReceiverID uint      `gorm:"not null"`
	Content    string    `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null"`
	Sender     User      `gorm:"foreignKey:SenderID"`
	Receiver   User      `gorm:"foreignKey:ReceiverID"`
}

type GroupMessage struct {
	ID        uint      `gorm:"primaryKey"`
	GroupID   uint      `gorm:"not null"`
	SenderID  uint      `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
	Group     Group     `gorm:"foreignKey:GroupID"`
	Sender    User      `gorm:"foreignKey:SenderID"`
}
