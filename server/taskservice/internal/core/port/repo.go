package port

import (
	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
)

type TaskRepo interface {
	Get(id uuid.UUID) (*domain.Task, error)
	Create(userID uuid.UUID, title, description string) (*domain.Task, error)
	Update(*domain.Task) (*domain.Task, error)
	Delete(id uuid.UUID) error
	List(userID uuid.UUID, limit, offset uint32) []*domain.Task
}
