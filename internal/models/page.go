package models

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Page struct {
	PageIndex int    `gorm:"-"` // 頁碼
	Size      int    `gorm:"-"` // 筆數
	TotalPage int    `gorm:"-"` // 總頁數
	Total     int    `gorm:"-"` // 總筆數
	Order     string `gorm:"-"` // 排序
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
	Data  []T `gorm:"-"`
}

func (p *PageResult[T]) GetIndex() int {
	return p.PageIndex
}

func (p *PageResult[T]) GetLimit() int {
	return p.Size
}

func (p *PageResult[T]) GetOffset() int {
	return (p.PageIndex - 1) * p.Size
}

func (p *PageResult[T]) GetOrder() string {
	return p.Order
}

func (p *PageResult[T]) SetTotal(count int64) {
	p.Total = int(count)
	p.TotalPage = (int(count) + p.Size - 1) / p.Size
}

func (p *PageResult[T]) GetTableName() string {
	tType := reflect.TypeOf(p.Data).Elem()
	t := reflect.New(tType).Elem()

	str := T.TableName("")
	fmt.Print(str)

	tm := t.MethodByName("TableName")
	if !tm.IsValid() {
		panic(errors.New("type does not have a TableName method"))
	}

	result := tm.Call(nil)
	if len(result) == 0 {
		panic(errors.New("TableName method did not return a value"))
	}

	return result[0].Interface().(string)
}

func (p *PageResult[T]) BeforeFind(db *gorm.DB) (err error) {
	if p.GetIndex() > 0 && p.GetLimit() > 0 {
		var count int64
		db = db.Table(p.GetTableName())
		err := db.Count(&count).Error
		if err != nil {
			return err
		}
		db = db.Offset(p.GetOffset()).Limit(p.GetLimit())
		p.SetTotal(count)
	}
	if p.GetOrder() != "" {
		db = db.Order(p.GetOrder())
	}
	return nil
}

//result := &po.PageResult[*po.Member]{
//	Page: cond.GetPager(),
//	Data: make([]*po.Member, 0),
//}
//if err := m.in.GetDB(ctx).Model(result).Find(&result.Data, cond).Error; err != nil {
//	return nil, nil, errs.UnknownError().Log(ctx, err)
//}
