package service

import (
	"errors"

	"github.com/wralith/aestimatio/server/authservice/internal/core/domain"
	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnexpected         = errors.New("unexpected error occurred")
)

type service struct {
	repo port.Repo
}

func New(repo port.Repo) port.Service {
	return &service{repo: repo}
}

// Login implements port.Service
func (s *service) Login(email string, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// Register implements port.Service
func (s *service) Register(username string, password string, email string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, ErrUnexpected
	}

	user, err := domain.NewUser(username, string(hashedPassword), email)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.repo.Create(user.Username, user.Password, user.Email)
	if err != nil {
		return nil, ErrUnexpected
	}

	return createdUser, nil
}
