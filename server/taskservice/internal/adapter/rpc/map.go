package rpc

import (
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
