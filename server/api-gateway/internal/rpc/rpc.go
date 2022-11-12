package rpc

import (
	pb "github.com/wralith/aestimatio/server/pb/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient pb.AuthServiceClient

func NewAuthClient() (AuthClient, error) {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewAuthServiceClient(conn)
	return client, nil
}
