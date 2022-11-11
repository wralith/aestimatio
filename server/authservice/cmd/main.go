package main

import (
	"fmt"

	"github.com/wralith/aestimatio/server/authservice/config"
	"github.com/wralith/aestimatio/server/authservice/internal/adapters/rpc"
	"github.com/wralith/aestimatio/server/authservice/internal/adapters/tkn"
	"github.com/wralith/aestimatio/server/authservice/internal/core/service"
	"github.com/wralith/aestimatio/server/authservice/internal/repo/inmemory"
)

func main() {
	config := config.Get()

	repo := inmemory.NewUserRepo()
	service := service.New(repo)
	token := tkn.New([]byte(config.JWTSecret))
	server := rpc.NewGRPCHandler(service, token, config.Port)

	fmt.Printf("grpc server started at port %s", config.Port)
	go server.Run()

	select {}
}
