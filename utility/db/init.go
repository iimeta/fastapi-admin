package db

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client          *mongo.Client
	DefaultDatabase string
)

func init() {

	ctx := gctx.New()
	var err error

	uri, err := config.Get(ctx, "mongodb.uri")
	if err != nil {
		logger.Error(ctx, err)
	}

	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri.String())); err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprint("MongoDB", err))
	}

	logger.Info(ctx, "MongoDB Successfully connected and pinged.")

	database, err := config.Get(ctx, "mongodb.database")
	if err != nil {
		logger.Error(ctx, err)
	}

	DefaultDatabase = database.String()

	unique := new(bool)
	*unique = true

	if _, err = client.Database(DefaultDatabase).Collection("user").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"user_id": 1},
		Options: &options.IndexOptions{
			Unique: unique,
		},
	}); err != nil {
		panic(err)
	}

	if _, err = client.Database(DefaultDatabase).Collection("account").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"account": 1},
		Options: &options.IndexOptions{
			Unique: unique,
		},
	}); err != nil {
		panic(err)
	}

	if _, err = client.Database(DefaultDatabase).Collection("app").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"app_id": 1},
		Options: &options.IndexOptions{
			Unique: unique,
		},
	}); err != nil {
		panic(err)
	}

	if _, err = client.Database(DefaultDatabase).Collection("chat").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"trace_id": 1},
	}); err != nil {
		panic(err)
	}

	if _, err = client.Database(DefaultDatabase).Collection("chat").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{bson.E{Key: "req_time", Value: -1}, bson.E{Key: "model_id", Value: 1}, bson.E{Key: "user_id", Value: 1}, bson.E{Key: "status", Value: 1}, bson.E{Key: "created_at", Value: -1}},
	}); err != nil {
		panic(err)
	}
}
