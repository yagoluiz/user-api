package repositories

import (
	"context"

	"github.com/yagoluiz/user-api/internal/db"
	"github.com/yagoluiz/user-api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	database   = "User"
	collection = "Users"
)

type UserRepository struct {
	database *db.MongoClient
}

func NewUserRepository(db *db.MongoClient) *UserRepository {
	return &UserRepository{database: db}
}

func (r *UserRepository) Search(term string) ([]*entity.User, error) {
	coll := r.database.Client.Database(database).Collection(collection)

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: term}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
