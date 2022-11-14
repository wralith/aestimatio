package rpc

import (
	"context"
	"time"

	pb "github.com/wralith/aestimatio/server/pb/gen/task"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown    = status.Error(codes.Unknown, "unknown error")
	ErrBadRequest = status.Error(codes.InvalidArgument, "bad request")
)

// CreateTask implements task.TaskServiceServer
func (h *GRPCHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	id, err := getUserFromMetadata(ctx)
	if err != nil {
		return nil, ErrBadRequest
	}

	task, err := h.service.CreateTask(id, req.Title, req.Description)
	if err != nil {
		return nil, ErrUnknown
	}

	return &pb.CreateTaskResponse{Task: taskToProtoResponse(task)}, nil
}

// GetTask implements task.TaskServiceServer
func (h *GRPCHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := h.validateUserAndGetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskResponse{Task: taskToProtoResponse(task)}, nil
}

// DeleteTask implements task.TaskServiceServer
func (h *GRPCHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	task, err := h.validateUserAndGetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = h.service.DeleteTask(task.ID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTaskResponse{}, nil
}

// UpdateTaskDetails implements task.TaskServiceServer
func (h *GRPCHandler) UpdateTaskDetails(ctx context.Context, req *pb.UpdateTaskDetailsRequest) (*pb.UpdateTaskDetailsResponse, error) {
	task, err := h.validateUserAndGetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res, err := h.service.UpdateDetails(task.ID, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskDetailsResponse{Task: taskToProtoResponse(res)}, nil
}

// UpdateTaskStatus implements task.TaskServiceServer
func (h *GRPCHandler) UpdateTaskStatus(ctx context.Context, req *pb.UpdateTaskStatusRequest) (*pb.UpdateTaskStatusResponse, error) {
	task, err := h.validateUserAndGetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if req.Status == pb.TaskStatus_TASK_STATUS_STARTED {
		task.Start()
	} else if req.Status == pb.TaskStatus_TASK_STATUS_COMPLETED || req.Status == pb.TaskStatus_TASK_STATUS_COMPLETED_AFTER_DEADLINE {
		task.Complete()
	} else if req.Status == pb.TaskStatus_TASK_STATUS_ABANDONED {
		task.Abandon()
	} else if req.Status == pb.TaskStatus_TASK_STATUS_DEADLINE_PASSED {
		task.Status = domain.STATUS_DEADLINE_PASSED
	} else {
		return nil, ErrBadRequest
	}

	res, err := h.service.Update(task)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskStatusResponse{Task: taskToProtoResponse(res)}, nil
}

func (h *GRPCHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task, err := h.validateUserAndGetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Status = domain.TaskStatus(req.Status)
	task.StartedAt = time.Unix(req.StartedAt, 0)
	task.CompletedAt = time.Unix(req.CompletedAt, 0)
	task.DeadlineAt = time.Unix(req.DeadlineAt, 0)
	task.AbandonedAt = time.Unix(req.AbandonedAt, 0)

	task, err = h.service.Update(task)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{Task: taskToProtoResponse(task)}, err
}

func (h *GRPCHandler) ListTasks(req *pb.ListTasksRequest, stream pb.TaskService_ListTasksServer) error {
	id, err := getUserFromMetadata(stream.Context())
	if err != nil {
		return err
	}

	list := h.service.ListTasks(id, req.Limit, req.Offset)

	for _, task := range list {
		if err := stream.Send(&pb.ListTasksResponse{Task: taskToProtoResponse(task)}); err != nil {
			return err
		}
	}

	return nil
}
