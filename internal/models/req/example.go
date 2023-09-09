package req

import (
	"gorm.io/gorm"

	"im/internal/models"
)

type ExampleGet struct {
	Id          string `json:"-" uri:"id" binding:"required" example:"1"` // ID
	Name        string `json:"name" example:"名字"`                         // 範例名
	Description string `json:"description" example:"取得描述"`                // 範例描述
}

type ExampleGetList struct {
	models.Page `gorm:"-"`
	Name        string   `json:"name" example:"名字"`                                     // 範例名
	Id          []string `json:"id" gorm:"id" binding:"required" example:"1,2,3,4,5,6"` // 範例ID列表
}

func (list ExampleGetList) Scope(db *gorm.DB) *gorm.DB {
	db = db.Where("id IN ?", list.Id)
	return db
}

type ExampleCreate struct {
	Name        string `json:"name" example:"小明"`             // 創建範例名
	Description string `json:"description"  example:"一個武林高手"` //創建範例描述
}

type ExampleUpdate struct {
	Id          string `json:"id,omitempty" binding:"required" example:"1"` // 修改範例ID
	Name        string `json:"name" example:"小明的孩子"`                        // 修改範例名
	Description string `json:"description" example:"繼承小明武功但沒天分的孩子"`         // 修改範例描述
}

type ExampleDelete struct {
	Id string `json:"id,omitempty"  example:"1"` // 刪除範例ID
}
