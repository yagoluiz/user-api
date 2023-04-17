package mocks

import (
	"time"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/yagoluiz/user-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsersMock() []*domain.User {
	usersMock := []*domain.User{
		{
			ID:        primitive.NewObjectID(),
			UserID:    mock.Anything,
			Name:      mock.Anything,
			Username:  mock.Anything,
			Priority:  0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID(),
			UserID:    mock.Anything,
			Name:      mock.Anything,
			Username:  mock.Anything,
			Priority:  0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return usersMock
}
