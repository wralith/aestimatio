package rpc

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
)

func Test_taskToProtoResponse(t *testing.T) {
	userId := uuid.New()
	title, desc := "Test", "Test Desc"
	task, _ := domain.CreateTask(userId, title, desc)
	proto := taskToProtoResponse(task)

	parsedId, err := uuid.Parse(proto.Id)
	require.NoError(t, err)

	require.Equal(t, task.ID, parsedId)
	require.WithinDuration(t, task.CreatedAt, time.Unix(proto.CreatedAt, 0), time.Second)
	//... Am i testing the language itself?
}
