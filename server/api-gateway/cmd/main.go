package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/wralith/aestimatio/server/api-gateway/config"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/handler"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/router"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rpc"
	"github.com/wralith/aestimatio/server/api-gateway/pkg/logger"
)

// @title       Aestimatio API
// @version     1.0
// @description Aestimatio API-Gateway.

// @license.name MIT

// @host     localhost:8080
// @BasePath /
// @schemes  http

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization

func main() {
	config := config.Get()
	logger.InitLogger(config.Logger.Pretty, config.Logger.Level)

	authService, err := rpc.NewAuthClient(config.Services.Auth)
	if err != nil {
		log.Fatal().Err(err)
	}

	taskService, err := rpc.NewTaskClient(config.Services.Task)
	if err != nil {
		log.Fatal().Err(err)
	}

	authHandler := handler.NewAuthHandler(authService)
	taskHandler := handler.NewTaskHandler(taskService)

	r := router.New(authHandler, taskHandler)

	go func() {
		if err := r.Echo.Start(":" + config.Server.Port); err != nil && err != http.ErrServerClosed {
			r.Echo.Logger.Fatal("Shutting down the server")
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.Echo.Shutdown(ctx); err != nil {
		r.Echo.Logger.Fatal(err)
	}
}
