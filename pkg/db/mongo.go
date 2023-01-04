package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "User"
	collection = "Users"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewConnection(conn string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(conn),
	)
	if err != nil {
		return nil, err
	}

	return &MongoClient{Client: client}, nil
}

func (m *MongoClient) CreateIndexes() error {
	coll := m.Client.Database(database).Collection(collection)

	model := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: "text"}, {Key: "username", Value: "text"}},
		Options: options.Index().SetDefaultLanguage("pt"),
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		return err
	}

	return nil
}
