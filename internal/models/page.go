package models

import (
	"gorm.io/gorm"
)

type Page struct {
	PageIndex int    `gorm:"-"`                  // 頁碼
	Size      int    `gorm:"-"`                  // 筆數
	TotalPage int    `gorm:"-"`                  // 總頁數
	Total     int64  `gorm:"-"`                  // 總筆數
	Order     string `gorm:"-" example:"id asc"` // 排序
}

func (p *Page) GetPager() *Page {
	return p
}

type ITable interface {
	TableName() string
}

// PageResult 帶有Page的實體
type PageResult[T ITable] struct {
	*Page `gorm:"-"`
	Data  []T
}

func (p *PageResult[T]) getLimit() int {
	return p.Size
}

func (p *PageResult[T]) getOffset() int {
	return (p.PageIndex - 1) * p.Size
}

func (p *PageResult[T]) getOrder() string {
	return p.Order
}

func (p *PageResult[T]) setTotalPage() {
	p.TotalPage = (int(p.Total) + p.Size - 1) / p.Size
}

func (p *PageResult[T]) PagerCond() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.PageIndex <= 0 {
			p.PageIndex = 1
		}
		if p.Size <= 0 {
			p.Size = 1
		}

		return db.Order(p.getOrder()).Offset(p.getOffset()).Limit(p.getLimit())
	}
}

//result := &po.PageResult[*po.Member]{
//	Page: cond.GetPager(),
//	Data: make([]*po.Member, 0),
//}
//if err := m.in.GetDB(ctx).Model(result).Find(&result.Data, cond).Error; err != nil {
//	return nil, nil, errs.UnknownError().Log(ctx, err)
//}
