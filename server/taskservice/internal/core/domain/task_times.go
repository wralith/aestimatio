package domain

import "time"

type TaskTimes struct {
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	DeadlineAt  time.Time
	AbandonedAt time.Time
}

// NewTaskTimes returns *TaskTimes with CreatedAt populated with Now() and else zero value for time.Time
func NewTaskTimes() *TaskTimes {
	return &TaskTimes{
		CreatedAt:   time.Now(),
		StartedAt:   time.Time{},
		CompletedAt: time.Time{},
		DeadlineAt:  time.Time{},
		AbandonedAt: time.Time{},
	}
}

// start updates StartedAt to time.Now
func (t *TaskTimes) start() {
	t.StartedAt = time.Now()
}

// complete updates CompletedAt to time.Now
func (t *TaskTimes) complete() {
	t.CompletedAt = time.Now()
}

// abandon updates AbandonedAt to time.Now
func (t *TaskTimes) abandon() {
	t.AbandonedAt = time.Now()
}

func (t *TaskTimes) setDeadline(new time.Time) {
	if new.After(time.Now()) {
		t.DeadlineAt = new
	}
}

// update updates provided fields
func (t *TaskTimes) update(new *TaskTimes) {
	if !new.CreatedAt.IsZero() || afterExpected(new.CreatedAt) {
		t.CreatedAt = new.CreatedAt
	}
	if !new.StartedAt.IsZero() || afterExpected(new.StartedAt) {
		t.StartedAt = new.StartedAt
	}
	if !new.CompletedAt.IsZero() || afterExpected(new.CompletedAt) {
		t.CompletedAt = new.CompletedAt
	}
	if !new.DeadlineAt.IsZero() || afterExpected(new.DeadlineAt) {
		t.DeadlineAt = new.DeadlineAt
	}
	if !new.AbandonedAt.IsZero() || afterExpected(new.AbandonedAt) {
		t.AbandonedAt = new.AbandonedAt
	}
}

// TODO: How to name this function??
// afterExpected checks if the time at least a min before the current time
func afterExpected(t time.Time) bool {
	return t.After(time.Now().Add(-1 * time.Minute))
}
