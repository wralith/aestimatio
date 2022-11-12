package response

import pb "github.com/wralith/aestimatio/server/pb/gen/auth"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type AuthResponse struct {
	User User   `json:"user,omitempty"`
	JWT  string `json:"jwt,omitempty"`
}

func LoginResponseFromProto(res *pb.LoginResponse) *AuthResponse {
	return &AuthResponse{
		User: User{
			ID:       res.User.GetId(),
			Email:    res.User.GetEmail(),
			Username: res.User.GetUsername(),
		},
		JWT: res.GetJwt(),
	}
}

func RegisterResponseFromProto(res *pb.RegisterResponse) *AuthResponse {
	return &AuthResponse{
		User: User{
			ID:       res.User.GetId(),
			Email:    res.User.GetEmail(),
			Username: res.User.GetUsername(),
		},
		JWT: res.GetJwt(),
	}
}
