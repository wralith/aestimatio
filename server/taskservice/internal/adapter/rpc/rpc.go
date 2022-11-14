package rpc

import (
	"context"

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

	res, err := h.service.UpdateStatus(task.ID, domain.TaskStatus(req.GetStatus()))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskStatusResponse{Task: taskToProtoResponse(res)}, nil
}

// TODO: UpdateTask implements task.TaskServiceServer
func (h *GRPCHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	// task, err := h.validateUserAndGetTask(ctx, req.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

// TODO: ListTasks implements task.TaskServiceServer
func (*GRPCHandler) ListTasks(*pb.ListTasksRequest, pb.TaskService_ListTasksServer) error {
	// panic("unimplemented")
	return nil
}
