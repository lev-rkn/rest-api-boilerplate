package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"rest-api-service/internal/app"
	"rest-api-service/internal/logger"
	"syscall"
	"time"
)

func main() {
	server := app.NewServer()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorLog("Listen server", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.ErrorLog("Server Shutdown:", err)
	}

	<-ctx.Done()
	slog.Info("timeout of 5 seconds, server exiting")
}
