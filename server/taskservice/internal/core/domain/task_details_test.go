package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTaskDetails(t *testing.T) {
	title, desc := "Test", "Test Description"
	v, err := NewTaskDetails(title, desc)

	require.NoError(t, err)
	require.Equal(t, v.Title, title)
	require.Equal(t, v.Description, desc)

	v2, err := NewTaskDetails("", "")
	require.Nil(t, v2)
	require.Equal(t, ErrInvalid, err)
}

func TestUpdateTitle(t *testing.T) {
	title, desc := "Test", "Test Description"
	v, _ := NewTaskDetails(title, desc)
	newTitle := "New Title"
	err := v.UpdateTitle(newTitle)

	require.NoError(t, err)
	require.Equal(t, newTitle, v.Title)
}

func TestUpdateDescription(t *testing.T) {
	title, desc := "Test", "Test Description"
	v, _ := NewTaskDetails(title, desc)
	newDesc := "New Desc"
	err := v.UpdateDescription(newDesc)

	require.NoError(t, err)
	require.Equal(t, newDesc, v.Description)
}
func TestTaskDetails_SwitchPlanned(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchPlanned()

	require.Equal(t, STATUS_PLANNED, v.Status)
}
func TestTaskDetails_SwitchStarted(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchStarted()

	require.Equal(t, STATUS_STARTED, v.Status)
}
func TestTaskDetails_SwitchCompleted(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchCompleted()

	require.Equal(t, STATUS_COMPLETED, v.Status)
}
func TestTaskDetails_SwitchAbandoned(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchAbandoned()

	require.Equal(t, STATUS_ABANDONED, v.Status)
}
func TestTaskDetails_SwitchDeadlinePassed(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchDeadlinePassed()

	require.Equal(t, STATUS_DEADLINE_PASSED, v.Status)
}
func TestTaskDetails_SwitchCompletedAfterDeadlinePassed(t *testing.T) {
	v, _ := NewTaskDetails("Test", "Test")
	v.switchCompletedAfterDeadlinePassed()

	require.Equal(t, STATUS_COMPLETED_AFTER_DEADLINE, v.Status)
}

func TestTaskDetails_validate(t *testing.T) {
	v := &TaskDetails{
		Title:       "Test",
		Description: "Test",
		Status:      1,
	}
	err := v.validate()
	require.NoError(t, err)

	v = &TaskDetails{
		Title:       "",
		Description: "",
		Status:      0,
	}
	err = v.validate()

	require.Error(t, err)
}
