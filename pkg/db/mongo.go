package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		log.Fatal(err)
	}

	return &MongoClient{Client: client}, nil
}

func (m *MongoClient) Ping() error {
	err := m.Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
