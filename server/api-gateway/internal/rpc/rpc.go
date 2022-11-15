package rpc

import (
	pbAuth "github.com/wralith/aestimatio/server/pb/gen/auth"
	pbTask "github.com/wralith/aestimatio/server/pb/gen/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient pbAuth.AuthServiceClient

func NewAuthClient(port string) (AuthClient, error) {
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	client := pbAuth.NewAuthServiceClient(conn)
	return client, nil
}

type TaskClient pbTask.TaskServiceClient

func NewTaskClient(port string) (TaskClient, error) {
	conn, err := grpc.Dial(":"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	client := pbTask.NewTaskServiceClient(conn)
	return client, nil
}
