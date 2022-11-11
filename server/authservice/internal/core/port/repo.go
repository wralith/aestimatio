package port

import (
	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/authservice/internal/core/domain"
)

type Repo interface {
	Get(id uuid.UUID) (*domain.User, error)
	Create(username, password, email string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
