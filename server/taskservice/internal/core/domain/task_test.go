package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	userID := uuid.New()
	title, desc := "Test", "Test Description"

	v, err := CreateTask(userID, title, desc)
	require.NoError(t, err)
	require.Equal(t, title, v.Title)
	require.Equal(t, desc, v.Description)

	_, err = CreateTask(userID, "", "")
	require.Error(t, err)
}

func TestTask_Start(t *testing.T) {
	v, _ := CreateTask(uuid.New(), "Test", "Test Description")
	v.Start()

	require.Equal(t, STATUS_STARTED, v.Status)
	require.WithinDuration(t, v.StartedAt, time.Now(), testDelta)
}

func TestTask_Complete(t *testing.T) {
	v, _ := CreateTask(uuid.New(), "Test", "Test Description")
	v.TaskTimes.setDeadline(time.Now().Add(5 * time.Hour))
	v.Complete()

	require.Equal(t, STATUS_COMPLETED, v.Status)
	require.WithinDuration(t, v.CompletedAt, time.Now(), testDelta)

	v2, _ := CreateTask(uuid.New(), "Test", "Test Description")
	v2.TaskTimes.setDeadline(time.Now().Add(-5 * time.Hour))
	v2.Complete()

	require.Equal(t, STATUS_COMPLETED_AFTER_DEADLINE, v2.Status)
	require.WithinDuration(t, v.CompletedAt, time.Now(), testDelta)
}

func TestTask_Abandon(t *testing.T) {
	v, _ := CreateTask(uuid.New(), "Test", "Test Description")
	v.Abandon()

	require.Equal(t, STATUS_ABANDONED, v.Status)
	require.WithinDuration(t, time.Now(), v.AbandonedAt, testDelta)
}

func TestTask_SetDeadline(t *testing.T) {
	v, _ := CreateTask(uuid.New(), "Test", "Test Description")
	dl := time.Now().Add(24 * 60 * time.Minute)
	v.SetDeadline(dl)

	require.WithinDuration(t, dl, v.DeadlineAt, testDelta)
}
