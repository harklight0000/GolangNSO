package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	config2 "nso/config"
	. "nso/logging"
	"time"
)

const connectTimeOut = 20 * time.Second

func InitMongoDB() *mongo.Database {
	config := config2.GetAppConfig()
	ctx, _ := context.WithTimeout(context.Background(), connectTimeOut)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDBURL))
	if err != nil {
		Logger.Info("Error when connect to mongodb", zap.Error(err))
	}
	return client.Database(config.DbName)
}
