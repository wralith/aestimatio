package domain

import (
	"errors"

	"github.com/wralith/aestimatio/server/taskservice/pkg/vld"
)

var (
	ErrInvalid = errors.New("invalid arguments")
)

type TaskStatus int32

// Copied from protoc generated code instead adding dependency
const (
	STATUS_UNSPECIFIED              TaskStatus = 0
	STATUS_PLANNED                  TaskStatus = 1
	STATUS_STARTED                  TaskStatus = 2
	STATUS_COMPLETED                TaskStatus = 3
	STATUS_ABANDONED                TaskStatus = 4
	STATUS_DEADLINE_PASSED          TaskStatus = 5
	STATUS_COMPLETED_AFTER_DEADLINE TaskStatus = 6
)

type TaskDetails struct {
	Title       string     `validate:"required,min=3"`
	Description string     `validate:"required,min=3"`
	Status      TaskStatus `validate:"required"`
}

func NewTaskDetails(title string, description string) (*TaskDetails, error) {

	t := &TaskDetails{
		Title:       title,
		Description: description,
		Status:      STATUS_PLANNED,
	}

	if err := t.validate(); err != nil {
		return nil, ErrInvalid
	}

	return t, nil
}

func (t *TaskDetails) UpdateTitle(title string) error {
	t.Title = title
	return t.validate()
}

func (t *TaskDetails) UpdateDescription(desc string) error {
	t.Description = desc
	return t.validate()
}

func (t *TaskDetails) switchPlanned() {
	t.Status = STATUS_PLANNED
}
func (t *TaskDetails) switchStarted() {
	t.Status = STATUS_STARTED
}
func (t *TaskDetails) switchCompleted() {
	t.Status = STATUS_COMPLETED
}
func (t *TaskDetails) switchAbandoned() {
	t.Status = STATUS_ABANDONED
}
func (t *TaskDetails) switchDeadlinePassed() {
	t.Status = STATUS_DEADLINE_PASSED
}
func (t *TaskDetails) switchCompletedAfterDeadlinePassed() {
	t.Status = STATUS_COMPLETED_AFTER_DEADLINE
}

func (t *TaskDetails) validate() error {
	v := vld.GetValidator()
	if err := v.Struct(t); err != nil {
		return err
	}
	return nil
}
