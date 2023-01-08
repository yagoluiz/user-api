package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId"`
	Name      string             `json:"name" bson:"name"`
	Username  string             `json:"username" bson:"username"`
	Priority  int                `json:"priority" bson:"priority,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type UserCSV struct {
	UserID   string `csv:"id"`
	Name     string `csv:"name"`
	Username string `csv:"username"`
}

type UserPriorityCSV struct {
	Priority string `csv:"priority"`
}

type UserRepository interface {
	Search(term string, limit, page int) ([]*User, error)
}

type UserUsecase interface {
	FindUser(term string, limit, page int) ([]*User, error)
}
