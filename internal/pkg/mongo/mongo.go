package mongo

import (
	"context"
	"strings"
	"time"

	"github.com/rickylin614/common/cmongo"
	"go.uber.org/dig"

	"go.mongodb.org/mongo-driver/mongo/options"

	"im/internal/pkg/config"
)

type DigIn struct {
	dig.In
	Config *config.Config
}

func NewMongoDB(in DigIn) cmongo.Client {
	if !in.Config.MongoConfig.Enable {
		return nil
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoConfig := &in.Config.MongoConfig

	credential := options.Credential{
		Username: mongoConfig.User,
		Password: mongoConfig.Password,
	}

	hosts := strings.Split(mongoConfig.Host, ",")
	connectOptions := options.Client().
		SetHosts(hosts).
		SetAuth(credential)

	client := cmongo.NewMongoDB()
	client, err := client.Connect(ctx, mongoConfig.DB, connectOptions)
	if err != nil {
		panic(err)
	}

	return client
}
