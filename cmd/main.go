package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yagoluiz/user-api/api" // Swag CLI
	"github.com/yagoluiz/user-api/configs"
	"github.com/yagoluiz/user-api/internal/api/handlers"
	"github.com/yagoluiz/user-api/internal/api/healths"
	"github.com/yagoluiz/user-api/internal/api/routers"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/internal/repositories/seed"
	"github.com/yagoluiz/user-api/internal/usercase"
	"github.com/yagoluiz/user-api/pkg/db"
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
	cfg, err := configs.GetConfigs()
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

	routers.UserRouters(r, h)

	healths := healths.NewHealthChecks(cfg)
	r.Get("/health", healths.HandlerFunc)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{Addr: cfg.Port, Handler: r}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
