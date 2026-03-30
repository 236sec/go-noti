package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goboilerplate.com/config"
)

type MongoDB struct {
	DB *mongo.Database
}

func (m *MongoDB) Create(ctx context.Context, collectionName string, doc interface{}) error {
	_, err := m.DB.Collection(collectionName).InsertOne(ctx, doc)
	return err
}

func (m *MongoDB) Find(ctx context.Context, collectionName string, filter Filter, dest interface{}) error {
	cursor, err := m.DB.Collection(collectionName).Find(ctx, filter)
	if err != nil {
		return err
	}
	return cursor.All(ctx, dest)
}

func (m *MongoDB) First(ctx context.Context, collectionName string, filter Filter, dest interface{}) error {
	err := m.DB.Collection(collectionName).FindOne(ctx, filter).Decode(dest)
	if err == mongo.ErrNoDocuments {
		return ErrRecordNotFound
	}
	return err
}

func initMongoDB() IDatabase {
	config := config.GetConfig().EnvConfig.MongoDB
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.MongodbUser, config.MongodbPassword, config.MongodbHost, config.MongodbPort)

	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Database connection established successfully")

	return &MongoDB{DB: client.Database(config.MongodbDbname)}
}
