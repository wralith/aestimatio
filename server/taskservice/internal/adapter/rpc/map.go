package rpc

import (
	"time"

	"github.com/google/uuid"
	pb "github.com/wralith/aestimatio/server/pb/gen/task"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
)

func taskToProtoResponse(task *domain.Task) *pb.Task {
	return &pb.Task{
		Id:          task.ID.String(),
		UserId:      task.UserID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      pb.TaskStatus(task.Status),
		CreatedAt:   task.CreatedAt.Unix(),
		StartedAt:   task.StartedAt.Unix(),
		CompletedAt: task.CompletedAt.Unix(),
		DeadlineAt:  task.DeadlineAt.Unix(),
		AbandonedAt: task.AbandonedAt.Unix(),
	}
}

func protoToTask(task *pb.Task) (*domain.Task, error) {
	id, err := uuid.Parse(task.Id)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(task.UserId)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:     id,
		UserID: userId,
		TaskDetails: &domain.TaskDetails{
			Title:       task.GetTitle(),
			Description: task.GetDescription(),
			Status:      domain.TaskStatus(task.GetStatus()),
		},
		TaskTimes: &domain.TaskTimes{
			CreatedAt:   time.Unix(task.GetCreatedAt(), 0),
			StartedAt:   time.Unix(task.GetStartedAt(), 0),
			CompletedAt: time.Unix(task.GetCompletedAt(), 0),
			DeadlineAt:  time.Unix(task.GetDeadlineAt(), 0),
			AbandonedAt: time.Unix(task.GetAbandonedAt(), 0),
		},
	}, nil
}
