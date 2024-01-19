package po

import "gorm.io/gorm"

type Page struct {
	Index     int    `gorm:"-"  json:"index" form:"index" example:"1""`     // 頁碼
	Size      int    `gorm:"-"  json:"size" form:"size" example:"20"`       // 筆數
	TotalPage int    `gorm:"-"  json:"-"`                                   // 總頁數
	Total     int64  `gorm:"-"  json:"-"`                                   // 總筆數
	Order     string `gorm:"-"  json:"order" form:"order" example:"id asc"` // 排序
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
	if p.Index <= 0 {
		p.Index = 1
	}
	if p.Size <= 0 {
		p.Size = 1
	}
	p.TotalPage = (int(p.Total) + p.Size - 1) / p.Size
}

func (p *PageResult[T]) PagerCond(db *gorm.DB) *gorm.DB {
	p.setTotalPage()
	order := p.getOrder()
	offset := p.getOffset()
	limit := p.getLimit()
	return db.Order(order).Offset(offset).Limit(limit)
}
