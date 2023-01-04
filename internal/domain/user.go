package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty" csv:"-"`
	UserId    string             `json:"userId" bson:"userId" csv:"id"`
	Name      string             `json:"name" bson:"name" csv:"name"`
	Username  string             `json:"username" bson:"username" csv:"username"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" csv:"-"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" csv:"-"`
}

type UserRepository interface {
	Search(term string, limit, page int) ([]*User, error)
}

type UserUsecase interface {
	FindUser(term string, limit, page int) ([]*User, error)
}
