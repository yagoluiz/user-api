package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/yagoluiz/user-api/internal/handlers"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/internal/routers"
	"github.com/yagoluiz/user-api/internal/usercase"
	"github.com/yagoluiz/user-api/pkg/db"
	"github.com/yagoluiz/user-api/pkg/db/seed"
)

type config struct {
	Port            string `env:"PORT" env-default:":8080"`
	MongoConnection string `env:"MONGO_CONNECTION" env-default:"mongodb://localhost:27017"`
	MongoDatabase   string `env:"MONGO_DATABASE" env-default:"User"`
	MongoCollection string `env:"MONGO_COLLECTION" env-default:"Users"`
}

func main() {
	var cfg config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewConnection(cfg.MongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	err = seed.NewSeed(database, cfg.MongoDatabase, cfg.MongoCollection)
	if err != nil {
		log.Fatal(err)
	}

	ur := repositories.NewUserRepository(database)
	uc := usercase.NewUserSearchUserCase(ur)
	h := handlers.NewUserSearchHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))

	routers.UserRouters(r, h)

	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
