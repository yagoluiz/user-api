package seed

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/yagoluiz/user-api/internal/entity"
	"log"
	"os"

	"github.com/yagoluiz/user-api/pkg/db"
)

func NewSeed(db *db.MongoClient, base, collection string) error {
	users, err := importUserData()
	if err != nil {
		log.Fatal(err)
	}

	err = insertUserData(db, base, collection, users)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func importUserData() ([]*entity.User, error) {
	file, err := os.OpenFile("resources/data/users.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []*entity.User

	if err := gocsv.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func insertUserData(db *db.MongoClient, base, collection string, users []*entity.User) error {
	data := make([]interface{}, len(users))

	for i := range data {
		data[i] = users[i]
	}

	coll := db.Client.Database(base).Collection(collection)

	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
