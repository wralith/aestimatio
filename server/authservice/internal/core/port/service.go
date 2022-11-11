package port

import "github.com/wralith/aestimatio/server/authservice/internal/core/domain"

type Service interface {
	Register(username, password, email string) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
}
