package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yagoluiz/user-api/api" // Swag CLI
	"github.com/yagoluiz/user-api/internal/api/handlers"
	"github.com/yagoluiz/user-api/internal/api/healths"
	"github.com/yagoluiz/user-api/internal/api/routers"
	"github.com/yagoluiz/user-api/internal/config"
	"github.com/yagoluiz/user-api/internal/db"
	"github.com/yagoluiz/user-api/internal/db/seed"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/internal/usercase"
)

// @title          User API
// @version        1.0
// @description    User management.
// @termsOfService http://swagger.io/terms/

// @contact.name Yago Luiz
// @contact.url  http://www.github.com/yagoluiz

// @host     localhost:8080
// @BasePath /api
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
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	routers.UserRouters(r, h)

	healths := healths.NewHealthChecks(cfg)
	r.Get("/health", healths.HandlerFunc)

	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
