package request

import (
	pb "github.com/wralith/aestimatio/server/pb/gen/task"
)

type CreateTask struct {
	Title       string `json:"title,omitempty" validate:"required,min=3"`
	Description string `json:"description,omitempty" validate:"required,min=3"`
	DeadlineAt  int64  `json:"deadline_at,omitempty" validate:"required"`
}

func (t *CreateTask) ToProto() *pb.CreateTaskRequest {
	return &pb.CreateTaskRequest{
		Title:       t.Title,
		Description: t.Description,
		DeadlineAt:  t.DeadlineAt,
	}
}
