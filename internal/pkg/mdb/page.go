package mdb

import (
	"context"

	"gorm.io/gorm"

	"im/internal/pkg/consts"
)

type Paged interface {
	GetIndex() int
	GetLimit() int
	GetOffset() int
	GetOrder() string
	SetTotal(int64)
	GetTableName() string
}

// AddPreQueryCallback 新增查詢前動作, 在有pager的資料查詢時，先查詢count，並且總數寫回count
func AddPreQueryCallback(db *gorm.DB) {
	db.Callback().Query().Before("gorm:query").Register("pre_page_callback", func(db *gorm.DB) {
		if _, ok := db.Statement.Context.Value(consts.COUNTING).(bool); ok {
			return
		}

		pager, ok := db.Statement.Model.(Paged)
		if ok && pager != nil && pager.GetIndex() > 0 && pager.GetLimit() > 0 {
			var count int64
			db = db.Table(pager.GetTableName())
			ctx := context.WithValue(db.Statement.Context, consts.COUNTING, true) // avoid loop
			err := db.Session(&gorm.Session{}).WithContext(ctx).Count(&count).Error
			if err != nil {
				db.AddError(err)
				return
			}

			db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit()) // 設定指定頁
			pager.SetTotal(count)                                     // 將查詢資訊寫回pager

		}
		if ok && pager != nil && len(pager.GetOrder()) > 0 {
			db = db.Order(pager.GetOrder())
		}
	})
}
