package app

import (
	"context"
	"rest-api-service/internal/config"
	"rest-api-service/internal/controller"
	"rest-api-service/internal/logger"
	"rest-api-service/internal/service"
	"rest-api-service/internal/storage"

	"net/http"

	_ "github.com/lib/pq"
)

func NewServer() *http.Server {
	config.MustLoad()
	logger.MustLoad()

	mainCtx := context.Background()

	repository := storage.NewStorage(mainCtx)
	service := service.NewService(repository)
	router := controller.NewRouter(mainCtx, service)

	server := &http.Server{
		Addr:    config.Cfg.HTTPServerAddress,
		Handler: router.Handler(),
	}

	return server
}
