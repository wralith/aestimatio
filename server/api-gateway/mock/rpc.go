package mock

import (
	"context"

	pb "github.com/wralith/aestimatio/server/pb/gen/auth"
	"google.golang.org/grpc"
)

type MockAuthClient struct {
}

func NewMockAuthClient() pb.AuthServiceClient {
	return &MockAuthClient{}
}

// Login implements auth.AuthServiceClient
func (*MockAuthClient) Login(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		User: &pb.User{
			Id:    "",
			Email: in.Email,
		},
		Jwt: "",
	}, nil
}

// Register implements auth.AuthServiceClient
func (*MockAuthClient) Register(ctx context.Context, in *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		User: &pb.User{
			Id:       "",
			Email:    in.Email,
			Username: in.Username,
		},
		Jwt: "",
	}, nil
}

// Validate expects "isValid" value from the context as a bool and returns new response with that, else return valid response
func (*MockAuthClient) Validate(ctx context.Context, in *pb.ValidateRequest, opts ...grpc.CallOption) (*pb.ValidateResponse, error) {
	isValid, ok := ctx.Value("isValid").(bool)

	if ok {
		return &pb.ValidateResponse{
			Valid: isValid,
		}, nil
	}

	return &pb.ValidateResponse{
		Valid: true,
	}, nil
}
