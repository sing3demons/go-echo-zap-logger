package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sing3demons/echo-logger/logger"
	"github.com/sing3demons/echo-logger/router"
	"go.uber.org/zap"
)

func main() {
	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	logger, close := logger.FileLogger("logs")
	defer close()
	e := router.RegRoute(logger)

	addr := ":8080"

	go func() {
		err := e.Start(addr)
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("unexpected shutdown the server", zap.Error(err))
		}
		logger.Info("gracefully shutdown the server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gCtx := context.Background()
	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}

}
