package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/port"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Tasks map[uuid.UUID]*domain.Task
type InMemoryTaskRepo struct {
	db Tasks
}

func NewInMemoryTaskRepo() port.TaskRepo {
	db := make(map[uuid.UUID]*domain.Task)
	return &InMemoryTaskRepo{db: db}
}

// Get implements port.TaskRepo
func (r *InMemoryTaskRepo) Get(id uuid.UUID) (*domain.Task, error) {
	if v, ok := r.db[id]; ok {
		return v, nil
	}
	return nil, ErrTaskNotFound
}

// Create implements port.TaskRepo
func (r *InMemoryTaskRepo) Create(userID uuid.UUID, title string, description string) (*domain.Task, error) {
	task, err := domain.CreateTask(userID, title, description)
	if err != nil {
		return nil, err
	}

	r.db[task.ID] = task
	return task, nil
}

// Update implements port.TaskRepo
func (r *InMemoryTaskRepo) Update(task *domain.Task) (*domain.Task, error) {
	if _, err := r.Get(task.ID); err != nil {
		return nil, ErrTaskNotFound
	}

	r.db[task.ID] = task
	return task, nil
}

// Delete implements port.TaskRepo
func (r *InMemoryTaskRepo) Delete(id uuid.UUID) error {
	if _, err := r.Get(id); err != nil {
		return ErrTaskNotFound
	}
	delete(r.db, id)
	return nil
}

func (r *InMemoryTaskRepo) List(userID uuid.UUID, limit, offset uint32) []*domain.Task {
	keys := []uuid.UUID{}
	for _, v := range r.db {
		if v.UserID == userID {
			keys = append(keys, v.ID)
		}
	}

	if limit > uint32(len(keys)) {
		limit = uint32(len(keys))
	}

	var res []*domain.Task
	for i := offset * limit; i < offset*limit+limit; i++ {
		if int(i) < len(keys) {
			res = append(res, r.db[keys[i]])
		}
	}

	return res
}
