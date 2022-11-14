package main

import (
	"fmt"

	"github.com/wralith/aestimatio/server/taskservice/config"
	"github.com/wralith/aestimatio/server/taskservice/internal/adapter/rpc"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/service"
	"github.com/wralith/aestimatio/server/taskservice/internal/repo/inmemory"
)

func main() {
	config := config.Get()

	repo := inmemory.NewInMemoryTaskRepo()
	service := service.New(repo)
	server := rpc.NewGRPCHandler(service, config.Port)

	fmt.Printf("grpc server started at port %s", config.Port)
	go server.Run()
	select {}
}
