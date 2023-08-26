package mdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"moul.io/zapgorm2"
)

const DB_CTX_KEY = "DB_GORM_CTX"

type Client interface {
	NewDB(ctx *gin.Context) *gorm.DB
	GetDB(ctx *gin.Context) *gorm.DB
	SetDB(ctx *gin.Context, db *gorm.DB) error
	GetMock() sqlmock.Sqlmock
}

// GormDB 實現 DBInterface 接口
type GormDB struct {
	in digIn
	db *gorm.DB
}

// GetDB 返回 *gorm.DB
func (g *GormDB) GetDB(ctx *gin.Context) *gorm.DB {
	newDb := g.db.WithContext(context.Background())
	if ctx == nil {
		return newDb
	}

	if val, ok := ctx.Get(DB_CTX_KEY); !ok {
		ctx.Set(DB_CTX_KEY, newDb)
		return newDb
	} else {
		if v, typeOK := val.(*gorm.DB); typeOK {
			return v
		} else {
			ctx.Set(DB_CTX_KEY, newDb)
			return newDb
		}
	}
}

func (g *GormDB) NewDB(ctx *gin.Context) *gorm.DB {
	newDb := g.db.WithContext(context.Background())
	if ctx == nil {
		return newDb
	}
	ctx.Set(DB_CTX_KEY, newDb)
	return newDb
}

func (g *GormDB) SetDB(ctx *gin.Context, db *gorm.DB) error {
	if ctx == nil || db == nil {
		return errors.New("SetDb not fuond context")
	}
	ctx.Set(DB_CTX_KEY, db)
	return nil
}

// GetDB 返回 *gorm.DB
func (g *GormDB) GetMock() sqlmock.Sqlmock {
	return nil
}

func newDB(in digIn) Client {
	dbSetting := [4]string{
		in.Config.MySQLConfig.Username,
		in.Config.MySQLConfig.Password,
		in.Config.MySQLConfig.Master,
		in.Config.MySQLConfig.Database,
	}
	masterDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC", dbSetting)
	slaveDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC", dbSetting)

	logger := zapgorm2.New(in.Log.GetLogger())
	logger.SetAsDefault()

	var err error
	db, err := gorm.Open(mysql.Open(masterDB),
		&gorm.Config{
			Logger:                 logger,
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", masterDB, err))
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(masterDB)},
		Replicas: []gorm.Dialector{mysql.Open(slaveDB)},
	}))

	// 註冊頁碼PreQuery
	AddPreQueryCallback(db)

	if in.Config.MySQLConfig.LogMode {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	in.Log.Info(context.Background(), fmt.Sprintf("Database [%s] Connect success", in.Config.MySQLConfig.Database))

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(in.Config.MySQLConfig.MaxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(in.Config.MySQLConfig.MaxOpen)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(in.Config.MySQLConfig.ConnMaxLifeSec) * time.Second)

	return &GormDB{in: in, db: db}
}
