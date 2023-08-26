package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.uber.org/dig"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"im/internal/pkg/config"
)

type DigIn struct {
	dig.In
	Config *config.Config
}

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(in DigIn) *Mongo {
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

	client, err := mongo.Connect(ctx, connectOptions)
	if err != nil {
		panic(fmt.Sprintf("Init mongo error: %s", err))
	}

	database := client.Database(mongoConfig.DB)

	return &Mongo{
		client:   client,
		database: database,
	}
}

func (m *Mongo) GetDB() *mongo.Database {
	return m.database
}
