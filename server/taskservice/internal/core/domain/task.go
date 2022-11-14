package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID
	UserID uuid.UUID
	*TaskDetails
	*TaskTimes
}

func CreateTask(userID uuid.UUID, title string, description string) (*Task, error) {
	id := uuid.New()
	details, err := NewTaskDetails(title, description)
	if err != nil {
		return nil, err
	}

	t := &Task{
		ID:          id,
		UserID:      userID,
		TaskDetails: details,
		TaskTimes:   NewTaskTimes(),
	}

	return t, nil
}

func (t *Task) Start() {
	t.TaskTimes.start()
	t.TaskDetails.switchStarted()
}

func (t *Task) Complete() {
	t.TaskTimes.complete()
	if t.TaskTimes.DeadlineAt.After(time.Now()) {
		t.TaskDetails.switchCompleted()
	} else {
		t.TaskDetails.switchCompletedAfterDeadlinePassed()
	}
}

func (t *Task) Abandon() {
	t.TaskDetails.switchAbandoned()
	t.TaskTimes.abandon()
}

func (t *Task) SetDeadline(new time.Time) {
	t.TaskTimes.setDeadline(new)
}

func (t *Task) Update(new *Task) (*Task, error) {
	t = new
	return t, nil
}
