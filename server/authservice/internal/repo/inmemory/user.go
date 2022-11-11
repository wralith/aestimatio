package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/authservice/internal/core/domain"
	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type InMemoryDB map[uuid.UUID]*domain.User

type UserRepo struct {
	db InMemoryDB
}

func NewUserRepo() port.Repo {
	db := make(map[uuid.UUID]*domain.User)

	return &UserRepo{
		db: db,
	}
}

// Create implements repo.Repo
func (r *UserRepo) Create(username, password, email string) (*domain.User, error) {
	user, err := domain.NewUser(username, password, email)
	if err != nil {
		return nil, err
	}

	r.db[user.ID] = user
	return user, nil
}

// Get implements repo.Repo
func (r *UserRepo) Get(id uuid.UUID) (*domain.User, error) {
	if v, ok := r.db[id]; ok {
		return v, nil
	}

	return nil, ErrUserNotFound
}

// GetByEmail implements repo.Repo
func (r *UserRepo) GetByEmail(email string) (*domain.User, error) {
	for _, v := range r.db {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, ErrUserNotFound
}
