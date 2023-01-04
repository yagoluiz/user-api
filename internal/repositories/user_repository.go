package repositories

import (
	"context"

	"github.com/yagoluiz/user-api/internal/entity"
	"github.com/yagoluiz/user-api/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *UserRepository) Search(term string, limit, page int) ([]*entity.User, error) {
	coll := r.database.Client.Database(database).Collection(collection)

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: term}}}}
	opts := options.Find().SetSkip(int64(limit) * int64(page)).SetLimit(int64(limit))

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
