package response

import pb "github.com/wralith/aestimatio/server/pb/gen/task"

type TaskStatus int32

const (
	TaskStatus_TASK_STATUS_UNSPECIFIED              TaskStatus = 0
	TaskStatus_TASK_STATUS_PLANNED                  TaskStatus = 1
	TaskStatus_TASK_STATUS_STARTED                  TaskStatus = 2
	TaskStatus_TASK_STATUS_COMPLETED                TaskStatus = 3
	TaskStatus_TASK_STATUS_ABANDONED                TaskStatus = 4
	TaskStatus_TASK_STATUS_DEADLINE_PASSED          TaskStatus = 5
	TaskStatus_TASK_STATUS_COMPLETED_AFTER_DEADLINE TaskStatus = 6
)

type TaskResponse struct {
	ID          string     `json:"id,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty"`
	CreatedAt   int64      `json:"created_at,omitempty"`
	StartedAt   int64      `json:"started_at,omitempty"`
	CompletedAt int64      `json:"completed_at,omitempty"`
	DeadlineAt  int64      `json:"deadline_at,omitempty"`
	AbandonedAt int64      `json:"abandoned_at,omitempty"`
}

func (m *TaskResponse) ToGRPC() *pb.Task {
	return &pb.Task{
		Id:          m.ID,
		UserId:      m.UserID,
		Title:       m.Title,
		Description: m.Description,
		Status:      pb.TaskStatus(m.Status),
		CreatedAt:   m.CreatedAt,
		StartedAt:   m.StartedAt,
		CompletedAt: m.CompletedAt,
		DeadlineAt:  m.DeadlineAt,
		AbandonedAt: m.AbandonedAt,
	}
}

func TaskResponseFromProto(t *pb.Task) *TaskResponse {
	return &TaskResponse{
		ID:          t.Id,
		UserID:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Status:      TaskStatus(t.Status),
		CreatedAt:   t.CreatedAt,
		StartedAt:   t.StartedAt,
		CompletedAt: t.CompletedAt,
		DeadlineAt:  t.DeadlineAt,
		AbandonedAt: t.AbandonedAt,
	}
}
