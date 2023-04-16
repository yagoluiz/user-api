package seed

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/yagoluiz/user-api/configs"
	"github.com/yagoluiz/user-api/internal/domain"
	"github.com/yagoluiz/user-api/pkg/db"
	"github.com/yagoluiz/user-api/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "User"
	collection = "Users"
)

func NewUserSeed(logger logger.Logger, cfg *configs.Config, db *db.MongoClient) error {
	done, err := importUserDone(db)
	if err != nil {
		return err
	}

	if done {
		logger.Info("Import has already been done.")
		return nil
	}

	logger.Info("Import will be done...")

	resourcePath := setResourcePath(cfg)

	users, err := importUserData(fmt.Sprintf("%s/users.csv", resourcePath))
	if err != nil {
		return err
	}

	priorities1, err := importPriorityData(fmt.Sprintf("%s/user_priority1.csv", resourcePath))
	if err != nil {
		return err
	}

	priorities2, err := importPriorityData(fmt.Sprintf("%s/user_priority2.csv", resourcePath))
	if err != nil {
		return err
	}

	err = insertUserData(db, users, priorities1, priorities2)
	if err != nil {
		return err
	}

	logger.Info("Import done!")

	return nil
}

func setResourcePath(cfg *configs.Config) string {
	var path string

	if cfg.Debug {
		path = "../resources/data"
	} else {
		path = "resources/data"
	}

	return path
}

func importUserDone(db *db.MongoClient) (bool, error) {
	coll := db.Client.Database(database).Collection(collection)

	var user domain.User
	err := coll.FindOne(context.TODO(), bson.D{}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func importUserData(dir string) ([]*domain.UserCSV, error) {
	file, err := os.OpenFile(dir, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []*domain.UserCSV

	if err := gocsv.Unmarshal(file, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func importPriorityData(dir string) ([]*domain.UserPriorityCSV, error) {
	file, err := os.OpenFile(dir, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var priorities []*domain.UserPriorityCSV

	if err := gocsv.Unmarshal(file, &priorities); err != nil {
		return nil, err
	}

	return priorities, nil
}

func insertUserData(db *db.MongoClient, users []*domain.UserCSV, priorities1 []*domain.UserPriorityCSV, priorities2 []*domain.UserPriorityCSV) error {
	data := make([]interface{}, len(users))

	for i := range data {
		date := time.Now().UTC()
		user := users[i]

		var priority int
		for _, v := range priorities1 {
			if v.Priority == user.UserID {
				priority = 1
				break
			}
		}
		if priority == 0 {
			for _, v := range priorities2 {
				if v.Priority == user.UserID {
					priority = 2
					break
				}
			}
		}

		userCreate := domain.User{
			UserID:    user.UserID,
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: date,
			UpdatedAt: date,
		}
		if priority != 0 {
			userCreate.Priority = priority
		}

		data[i] = userCreate
	}

	coll := db.Client.Database(database).Collection(collection)

	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}

	return err
}
