package sqldb

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type MockDB struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (g *MockDB) Session(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return g.db.WithContext(context.Background())
	}
	return g.db.WithContext(ctx)
}

func (g *MockDB) Begin(ctx context.Context) *gorm.DB {
	tx := g.db.WithContext(ctx).Clauses(dbresolver.Write).Begin()
	return tx
}

func (g *MockDB) GetMock() sqlmock.Sqlmock {
	return g.mock
}

func NewMockDB() Client {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	return &MockDB{mock: mock, db: gormDB}
}
