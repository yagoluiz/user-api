package main

import (
	"context"
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
	"github.com/yagoluiz/user-api/internal/usecase"
	"github.com/yagoluiz/user-api/pkg/db"
	"github.com/yagoluiz/user-api/pkg/db/seed"
	"github.com/yagoluiz/user-api/pkg/logger"
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
		panic(err)
	}

	logger, err := logger.NewLogger(cfg)
	if err != nil {
		panic(err)
	}

	logger.Infof("API debug: %v", cfg.Debug)

	database, err := db.NewConnection(cfg.MongoConnection)
	if err != nil {
		logger.Fatal(err)
	}

	err = database.CreateIndexes()
	if err != nil {
		logger.Fatal(err)
	}

	err = seed.NewUserSeed(logger, cfg, database)
	if err != nil {
		logger.Fatal(err)
	}

	ur := repositories.NewUserRepository(logger, database)
	uc := usecase.NewUserSearchUseCase(logger, ur)
	h := handlers.NewUserSearchHandler(logger, uc)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	routers.UserRouters(r, h)

	healths := healths.NewHealthChecks(cfg)
	r.Get("/health", healths.HandlerFunc)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{Addr: cfg.Port, Handler: r}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	gracefulShutdownServer(serverCtx, logger, server, serverStopCtx, err)
}

func gracefulShutdownServer(serverCtx context.Context, logger logger.Logger, server *http.Server, serverStopCtx context.CancelFunc, err error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			logger.Fatal(err)
		}
		serverStopCtx()
	}()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	<-serverCtx.Done()
}
