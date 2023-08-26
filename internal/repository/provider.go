package repository

import (
	"go.uber.org/dig"
)

// digIn repository require indendency
type digIn struct {
	dig.In

	//SqlDB    mysqlcli.DBInterface
	//RDB      redis.UniversalClient
	//Mongo    *mongocli.Mongo
	//SqlFiles *sql.SqlFiles
}

//func (in digIn) GetDB(ctx context.Context) *gorm.DB {
//	return in.SqlDB.GetDB(ctx)
//}

type Repository struct {
}

func NewRepository(in digIn) *Repository {
	return &Repository{}
}
