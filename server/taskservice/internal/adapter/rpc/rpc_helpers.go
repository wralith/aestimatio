package rpc

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getSubFromJWT(token string) (uuid.UUID, error) {
	v, _ := jwt.Parse(token, nil)
	if v == nil {
		return uuid.Nil, ErrUnknown
	}

	if claims, ok := v.Claims.(jwt.MapClaims); ok {
		sub := claims["sub"]
		return uuid.Parse(sub.(string))
	}

	return uuid.Nil, ErrUnknown
}

func getUserFromMetadata(ctx context.Context) (id uuid.UUID, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, ErrBadRequest
	}

	token := md.Get("jwt")
	if len(token) > 0 {
		id, err = getSubFromJWT(token[0])
		if err != nil {
			return uuid.Nil, ErrBadRequest
		}
	}

	return id, nil
}

func (h *GRPCHandler) validateUserAndGetTask(ctx context.Context, reqId string) (*domain.Task, error) {
	userId, err := getUserFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(reqId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "unable to parse id")
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "task not found")
	}

	if task.UserID.String() != userId.String() {
		return nil, status.Error(codes.InvalidArgument, "task not found")
	}

	return task, nil
}
