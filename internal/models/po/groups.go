package po

import "time"

type Groups struct {
	ID           string    `gorm:"primarykey;column:id"`
	GroupName    string    `gorm:"not null;column:group_name;unique"`
	Description  string    `gorm:"type:text;column:description"`
	GroupOwnerID string    `gorm:"not null;column:group_owner_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (*Groups) TableName() string {
	return "groups"
}
