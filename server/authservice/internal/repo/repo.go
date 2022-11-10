package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/authservice/internal/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Repo interface {
	Get(id uuid.UUID) (*domain.User, error)
	Create(username, password, email string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
