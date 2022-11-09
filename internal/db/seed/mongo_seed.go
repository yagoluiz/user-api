package seed

import (
	"context"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/yagoluiz/user-api/internal/db"
	"github.com/yagoluiz/user-api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "User"
	collection = "Users"
)

func NewUserSeed(db *db.MongoClient) error {
	done, err := importUserDone(db)
	if err != nil {
		return err
	}

	if done {
		return nil
	}

	users, err := importUserData()
	if err != nil {
		return err
	}

	err = insertUserData(db, users)
	if err != nil {
		return err
	}

	return nil
}

func importUserDone(db *db.MongoClient) (bool, error) {
	coll := db.Client.Database(database).Collection(collection)

	var user entity.User
	err := coll.FindOne(context.TODO(), bson.D{}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func importUserData() ([]*entity.User, error) {
	file, err := os.OpenFile("resources/data/users.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []*entity.User

	if err := gocsv.Unmarshal(file, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func insertUserData(db *db.MongoClient, users []*entity.User) error {
	date := time.Now().UTC()

	data := make([]interface{}, len(users))
	for i := range data {
		users[i].CreatedAt = date
		users[i].UpdatedAt = date
		data[i] = users[i]
	}

	coll := db.Client.Database(database).Collection(collection)

	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}

	return err
}
