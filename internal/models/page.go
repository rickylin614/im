package models

import (
	"gorm.io/gorm"
)

type Page struct {
	Index     int    `gorm:"-" json:"index"`                  // 頁碼
	Size      int    `gorm:"-" json:"size"`                   // 筆數
	TotalPage int    `gorm:"-" json:"-"`                      // 總頁數
	Total     int64  `gorm:"-" json:"-"`                      // 總筆數
	Order     string `gorm:"-" example:"id asc" json:"order"` // 排序
}

func (p *Page) GetPager() *Page {
	return p
}

type ITable interface {
	TableName() string
}

// PageResult 帶有Page的實體
type PageResult[T ITable] struct {
	*Page `json:"page" gorm:"-"`
	Data  []T `json:"data"  gorm:"-"`
}

func (p *PageResult[T]) getLimit() int {
	return p.Size
}

func (p *PageResult[T]) getOffset() int {
	return (p.Index - 1) * p.Size
}

func (p *PageResult[T]) getOrder() string {
	return p.Order
}

func (p *PageResult[T]) setTotalPage() {
	p.TotalPage = (int(p.Total) + p.Size - 1) / p.Size
}

func (p *PageResult[T]) PagerCond() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Index <= 0 {
			p.Index = 1
		}
		if p.Size <= 0 {
			p.Size = 1
		}
		p.setTotalPage()

		return db.Order(p.getOrder()).Offset(p.getOffset()).Limit(p.getLimit())
	}
}
