package po

// type Users struct {
// 	Id           string `gorm:"column:id"`
// 	Username     string `gorm:"column:username"`
// 	Email        string `gorm:"column:email"`
// 	PasswordHash string `gorm:"column:password_hash"`
// 	CreatedAt    int64  `gorm:"column:created_at"`
// 	UpdatedAt    int64  `gorm:"column:updated_at"`
// }

// type UserProfile struct {
// 	UserID      uint `gorm:"primaryKey"`
// 	AvatarURL   string
// 	DisplayName string
// 	Bio         string
// 	Location    string
// 	Website     string
// 	CreatedAt   time.Time `gorm:"not null"`
// 	UpdatedAt   time.Time `gorm:"not null"`
// }

// type Friend struct {
// 	UserID1          uint                    `gorm:"primaryKey"`
// 	UserID2          uint                    `gorm:"primaryKey"`
// 	FriendshipStatus consts.FriendshipStatus `gorm:"not null"`
// 	CreatedAt        time.Time               `gorm:"not null"`
// 	UpdatedAt        time.Time               `gorm:"not null"`
// }

// type FriendRequests struct {
// 	ID            uint      `gorm:"primaryKey"`
// 	SenderID      uint      `gorm:"not null"`
// 	ReceiverID    uint      `gorm:"not null"`
// 	RequestStatus string    `gorm:"not null"`
// 	CreatedAt     time.Time `gorm:"not null"`
// 	UpdatedAt     time.Time `gorm:"not null"`
// 	// Sender        User      `gorm:"foreignKey:SenderID"`
// 	// Receiver      User      `gorm:"foreignKey:ReceiverID"`
// }

// type Group struct {
// 	ID           uint   `gorm:"primaryKey"`
// 	Name         string `gorm:"unique;not null"`
// 	Description  string
// 	GroupOwnerID uint      `gorm:"not null"`
// 	CreatedAt    time.Time `gorm:"not null"`
// 	UpdatedAt    time.Time `gorm:"not null"`
// 	// GroupOwner   User      `gorm:"foreignKey:GroupOwnerID"`
// }

// type GroupMember struct {
// 	GroupID  uint      `gorm:"primaryKey"`
// 	UserID   uint      `gorm:"primaryKey"`
// 	Role     string    `gorm:"not null"`
// 	JoinedAt time.Time `gorm:"not null"`
// }

// type GroupInvitation struct {
// 	ID               uint      `gorm:"primaryKey"`
// 	GroupID          uint      `gorm:"not null"`
// 	InviterID        uint      `gorm:"not null"`
// 	InviteeID        uint      `gorm:"not null"`
// 	InvitationStatus string    `gorm:"not null"`
// 	CreatedAt        time.Time `gorm:"not null"`
// 	UpdatedAt        time.Time `gorm:"not null"`
// }

// type GroupRequest struct {
// 	ID            uint      `gorm:"primaryKey"`
// 	GroupID       uint      `gorm:"not null"`
// 	RequesterID   uint      `gorm:"not null"`
// 	RequestStatus string    `gorm:"not null"`
// 	CreatedAt     time.Time `gorm:"not null"`
// 	UpdatedAt     time.Time `gorm:"not null"`
// }

// type Message struct {
// 	ID         uint                 `gorm:"primaryKey"`
// 	SenderID   uint                 `gorm:"not null"`
// 	ReceiverID uint                 `gorm:"not null"`
// 	MsgContent    string               `gorm:"not null"`
// 	Timestamp  time.Time            `gorm:"not null"`
// 	Status     consts.MessageStatus `gorm:"not null"` // 狀態: 正常, 收回, 刪除, 隱藏
// }
