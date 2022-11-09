package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yagoluiz/user-api/internal/config"
	"github.com/yagoluiz/user-api/internal/db"
	"github.com/yagoluiz/user-api/internal/db/seed"
	"github.com/yagoluiz/user-api/internal/handlers"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/internal/routers"
	"github.com/yagoluiz/user-api/internal/usercase"
)

func main() {
	cfg, err := config.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewConnection(cfg.MongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	err = database.CreateIndexes()
	if err != nil {
		log.Fatal(err)
	}

	err = seed.NewUserSeed(database)
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
