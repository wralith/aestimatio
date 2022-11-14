package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var testDelta = time.Millisecond * 5

func TestNewTaskTimes(t *testing.T) {
	got := NewTaskTimes()
	require.WithinDuration(t, time.Now(), got.CreatedAt, testDelta)
}

func TestTaskTimes_Start(t *testing.T) {
	v := NewTaskTimes()
	v.start()

	require.WithinDuration(t, time.Now(), v.StartedAt, testDelta)
}

func TestTaskTimes_Complete(t *testing.T) {
	v := NewTaskTimes()
	v.complete()

	require.WithinDuration(t, time.Now(), v.CompletedAt, testDelta)
}

func TestTaskTimes_Abandon(t *testing.T) {
	v := NewTaskTimes()
	v.abandon()

	require.WithinDuration(t, time.Now(), v.AbandonedAt, testDelta)
}

func TestTaskTimes_SetDeadline(t *testing.T) {
	v := NewTaskTimes()
	dl := time.Now().Add(30 * 24 * time.Hour)
	v.setDeadline(dl)

	require.Equal(t, v.DeadlineAt, dl)
}

func TestTaskTimes_Update(t *testing.T) {
	v := NewTaskTimes()
	v2 := NewTaskTimes()
	v2.StartedAt = time.Now().Add(5 * time.Minute)
	v2.CompletedAt = time.Now().Add(60 * time.Minute)
	v2.AbandonedAt = time.Now().Add(30 * time.Minute)
	v2.DeadlineAt = time.Now().Add(3 * 60 * time.Minute)

	v.update(v2)
	require.Equal(t, v, v2)
}

func Test_afterExpected(t *testing.T) {
	got := afterExpected(time.Now())
	require.True(t, got)

	got = afterExpected(time.Now().Add(50 * time.Minute))
	require.True(t, got)

	got = afterExpected(time.Now().Add(-30 * time.Second))
	require.True(t, got)
}
