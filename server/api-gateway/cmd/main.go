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

func main() {
	config := config.Get()
	logger.InitLogger(config.Logger.Pretty, config.Logger.Level)

	authService, err := rpc.NewAuthClient()
	if err != nil {
		log.Fatal().Err(err)
	}
	authHandler := handler.NewAuthHandler(authService)

	r := router.New(authHandler)

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
