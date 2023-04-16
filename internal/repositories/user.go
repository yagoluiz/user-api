package repositories

import (
	"context"

	"github.com/yagoluiz/user-api/internal/domain"
	"github.com/yagoluiz/user-api/pkg/db"
	"github.com/yagoluiz/user-api/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "User"
	collection = "Users"
)

type UserRepositoryInterface interface {
	Search(term string, limit, page int) ([]*domain.User, error)
}

type UserRepository struct {
	logger   logger.Logger
	database *db.MongoClient
}

func NewUserRepository(l logger.Logger, db *db.MongoClient) *UserRepository {
	return &UserRepository{logger: l, database: db}
}

func (r *UserRepository) Search(term string, limit, page int) ([]*domain.User, error) {
	r.logger.Infof("Search user respository - Term: %s - Limit: %d - Page: %d", term, limit, page)

	coll := r.database.Client.Database(database).Collection(collection)

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: term}}}}
	opts := options.Find().SetSort(bson.D{{Key: "priority", Value: 1}}).SetSkip(int64(limit) * int64(page)).SetLimit(int64(limit))

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
