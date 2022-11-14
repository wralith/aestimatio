package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/port"
)

var (
	ErrNotFound      = errors.New("task not found")
	ErrStatusInvalid = errors.New("invalid status")
	ErrUnknown       = errors.New("unknown error")
)

type TaskService struct {
	repo port.TaskRepo
}

func New(repo port.TaskRepo) port.TaskService {
	return &TaskService{repo: repo}
}

// CreateTask implements port.TaskService
func (t *TaskService) CreateTask(userID uuid.UUID, title string, description string) (*domain.Task, error) {
	task, err := t.repo.Create(userID, title, description)
	if err != nil {
		return nil, err
	}
	return task, err
}

// DeleteTask implements port.TaskService
func (t *TaskService) DeleteTask(id uuid.UUID) error {
	return t.repo.Delete(id)
}

// GetTask implements port.TaskService
func (t *TaskService) GetTask(id uuid.UUID) (*domain.Task, error) {
	task, err := t.repo.Get(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return task, nil
}

// ListTasks implements port.TaskService
func (t *TaskService) ListTasks(userID uuid.UUID, limit, offset uint32) []*domain.Task {
	list := t.repo.List(userID, limit, offset)
	if len(list) == 0 {
		return nil
	}
	return t.repo.List(userID, limit, offset)
}

// Update implements port.TaskService
func (t *TaskService) Update(task *domain.Task) (*domain.Task, error) {
	return t.repo.Update(task)
}

// UpdateDetails implements port.TaskService
func (t *TaskService) UpdateDetails(id uuid.UUID, title string, description string) (*domain.Task, error) {
	task, err := t.GetTask(id)
	if err != nil {
		return nil, err
	}

	task.UpdateTitle(title)
	task.UpdateDescription(description)
	task, err = t.repo.Update(task)
	if err != nil {
		return nil, ErrUnknown
	}

	return task, nil
}

// UpdateStatus implements port.TaskService
func (t *TaskService) UpdateStatus(id uuid.UUID, status domain.TaskStatus) (*domain.Task, error) {
	task, err := t.GetTask(id)
	if err != nil {
		return nil, err
	} else if status == domain.STATUS_STARTED {
		task.Start()
	} else if status == domain.STATUS_COMPLETED {
		task.Complete()
	} else if status == domain.STATUS_ABANDONED {
		task.Abandon()
	} else {
		return nil, ErrStatusInvalid
	}

	task, err = t.repo.Update(task)
	if err != nil {
		return nil, ErrUnknown
	}

	return task, nil
}
