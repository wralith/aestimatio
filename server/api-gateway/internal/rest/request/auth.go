package request

import (
	pb "github.com/wralith/aestimatio/server/pb/gen/auth"
)

type LoginRequest struct {
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=6"`
}

func (r *LoginRequest) ToProto() *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"min=4"`
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=6"`
}

func (r *RegisterRequest) ToProto() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Email:    r.Email,
		Password: r.Password,
		Username: r.Username,
	}
}

type ValidateRequest struct {
	JWT string `json:"jwt"`
}

func (r *ValidateRequest) ToProto() *pb.ValidateRequest {
	return &pb.ValidateRequest{
		Jwt: r.JWT,
	}
}
