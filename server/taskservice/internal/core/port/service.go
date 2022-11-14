package port

import (
	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
)

type TaskService interface {
	GetTask(id uuid.UUID) (*domain.Task, error)
	CreateTask(userID uuid.UUID, title, description string) (*domain.Task, error)
	Update(*domain.Task) (*domain.Task, error)
	UpdateStatus(id uuid.UUID, status domain.TaskStatus) (*domain.Task, error)
	UpdateDetails(id uuid.UUID, title, description string) (*domain.Task, error)
	DeleteTask(id uuid.UUID) error
	ListTasks(userID uuid.UUID, limit, offset uint32) []*domain.Task
}
