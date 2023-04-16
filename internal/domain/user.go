package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Name      string             `bson:"name"`
	Username  string             `bson:"username"`
	Priority  int                `bson:"priority,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserCSV struct {
	UserID   string `csv:"id"`
	Name     string `csv:"name"`
	Username string `csv:"username"`
}

type UserPriorityCSV struct {
	Priority string `csv:"priority"`
}
