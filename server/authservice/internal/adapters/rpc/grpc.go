package rpc

import (
	"log"
	"net"

	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
	pb "github.com/wralith/aestimatio/server/pb/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	service    port.Service
	jwt        port.Token
	serverPort string
}

func NewGRPCHandler(service port.Service, jwt port.Token, serverPort string) *GRPCHandler {
	return &GRPCHandler{service: service, jwt: jwt, serverPort: serverPort}
}

func (h *GRPCHandler) Run() {
	lst, err := net.Listen("tcp", ":"+h.serverPort)
	if err != nil {
		log.Fatalf("failed to listen at port %s, err: %v", h.serverPort, err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, h)

	reflection.Register(s)

	err = s.Serve(lst)
	if err != nil {
		log.Fatalf("failed to server grpc over port %s, err: %v", h.serverPort, err)
	}

	defer s.Stop()
}
