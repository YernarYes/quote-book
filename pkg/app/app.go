package app

import (
	"context"
	"fmt"
	"log"
	"quotes/config"
	"quotes/internal/core/domain/quote"
	"quotes/internal/core/handler"
	"quotes/internal/core/infra/storage"
	"quotes/pkg/database"
	httpserver "quotes/pkg/http"
	"quotes/pkg/logger"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func StartApp(cfg *config.Config) error {

	dsn := database.PostgresDSN(*cfg)

	conn, err := database.OpenDB(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	logger := logger.LoggerInit()

	storage := storage.NewStorage(conn)
	service := quote.NewService(logger, storage)
	h := handler.NewHandler(service, logger)
	router := mux.NewRouter()
	listen := httpserver.NewHTTPServer(h, router)

	logger.Info("Succesfully connected", "host:", cfg.HTTP.Host, "port:", cfg.HTTP.Port)
	if err := listen.StartHTTPServer(cfg.HTTP.Host, cfg.HTTP.Port); err != nil {
		log.Fatalf("failed to run http server %v", err)
	}

	return nil
}
