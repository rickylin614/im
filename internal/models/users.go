package models

type Users struct {
	ID           string `gorm:"column:id"`
	Username     string `gorm:"column:username"`
	Nickname     string `gorm:"column:nickname"`
	PasswordHash string `gorm:"column:password_hash"`
	Email        string `gorm:"column:email"`
	PhoneNumber  string `gorm:"column:phone_number"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`
}

func (*Users) TableName() string {
	return "users"
}
