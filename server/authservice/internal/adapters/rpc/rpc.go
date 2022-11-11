package rpc

import (
	context "context"

	pb "github.com/wralith/aestimatio/server/pb/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login implements auth.AuthServiceServer
func (h *GRPCHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := h.service.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unknown, "unexpected error") // Need more errors!
	}

	jwtString, err := h.jwt.GenerateJWT(result.ID.String(), result.Email)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unexpected error") // Need more errors!
	}

	return &pb.LoginResponse{
		User: &pb.User{
			Id:       result.ID.String(),
			Email:    result.Email,
			Username: result.Username,
		},
		Jwt: jwtString,
	}, nil
}

// Register implements auth.AuthServiceServer
func (h *GRPCHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	result, err := h.service.Register(req.GetUsername(), req.GetPassword(), req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Unknown, "unexpected error")
	}

	jwtString, err := h.jwt.GenerateJWT(result.ID.String(), result.Email)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unexpected error") // Need more errors!
	}

	return &pb.RegisterResponse{
		User: &pb.User{
			Id:       result.ID.String(),
			Email:    result.Email,
			Username: result.Username,
		},
		Jwt: jwtString,
	}, nil
}

func (h *GRPCHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	res := h.jwt.VerifyJWT(req.GetJwt())
	return &pb.ValidateResponse{Valid: res}, nil
}
