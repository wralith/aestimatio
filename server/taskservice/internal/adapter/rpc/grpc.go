package rpc

import (
	"log"
	"net"

	pb "github.com/wralith/aestimatio/server/pb/gen/task"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCHandler struct {
	pb.UnimplementedTaskServiceServer
	service    port.TaskService
	serverPort string
}

func NewGRPCHandler(service port.TaskService, serverPort string) *GRPCHandler {
	return &GRPCHandler{service: service, serverPort: serverPort}
}

func (h *GRPCHandler) Run() {
	lst, err := net.Listen("tcp", ":"+h.serverPort)
	if err != nil {
		log.Fatalf("failed to listen at port %s, err: %v", h.serverPort, err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, h)

	reflection.Register(s)

	err = s.Serve(lst)
	if err != nil {
		log.Fatalf("failed to server grpc over port %s, err: %v", h.serverPort, err)
	}

	defer s.Stop()
}
