package repositories

import (
	"github.com/yagoluiz/user-api/internal/entity"
	"github.com/yagoluiz/user-api/pkg/db"
)

type UserRepository struct {
	database *db.MongoClient
}

func NewUserRepository(db *db.MongoClient) *UserRepository {
	return &UserRepository{database: db}
}

func (r *UserRepository) Search() (*entity.User, error) {
	return nil, nil
}
