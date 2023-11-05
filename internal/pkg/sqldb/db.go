package sqldb

import (
	"context"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"moul.io/zapgorm2"
)

type Client interface {
	Session(ctx context.Context) *gorm.DB
	Begin(ctx context.Context) *gorm.DB
	GetMock() sqlmock.Sqlmock
}

// GormDB 實現 DBInterface 接口
type GormDB struct {
	in digIn
	db *gorm.DB
}

// GetDB 返回 *gorm.DB
func (g *GormDB) Session(ctx context.Context) *gorm.DB {
	newDb := g.db.WithContext(ctx)
	return newDb
}

func (g *GormDB) Begin(ctx context.Context) *gorm.DB {
	tx := g.db.WithContext(ctx).Clauses(dbresolver.Write).Begin()
	return tx
}

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
	masterDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC", dbSetting[0], dbSetting[1], dbSetting[2], dbSetting[3])
	slaveDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC", dbSetting[0], dbSetting[1], dbSetting[2], dbSetting[3])

	gormConfig := &gorm.Config{
		//Logger:                 logger,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if l, ok := in.Log.GetLogger().(*zap.Logger); ok {
		logger := zapgorm2.New(l)
		logger.SetAsDefault()
		gormConfig.Logger = logger
	}

	var err error
	db, err := gorm.Open(mysql.Open(masterDB), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", masterDB, err))
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(masterDB)},
		Replicas: []gorm.Dialector{mysql.Open(slaveDB)},
	}))

	// 註冊頁碼PreQuery
	// AddPreQueryCallback(db)

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
