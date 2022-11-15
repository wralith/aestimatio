package handler

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/request"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/response"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rpc"
	"github.com/wralith/aestimatio/server/pb/gen/task"
	"google.golang.org/grpc/metadata"
)

type TaskHandler struct {
	svc rpc.TaskClient
}

func NewTaskHandler(svc rpc.TaskClient) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) Get(c echo.Context) error {
	id := c.Param("id")
	token, err := getAuthHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	ctx := attachTokenToMetadata(token)
	task, err := h.svc.GetTask(ctx, &task.GetTaskRequest{Id: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal")
	}

	res := response.TaskResponseFromProto(task.GetTask())

	return c.JSON(http.StatusOK, res)
}

func (h *TaskHandler) Create(c echo.Context) error {
	token, err := getAuthHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	var req request.CreateTask

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrBadRequest.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrInvalid.Error())
	}

	ctx := attachTokenToMetadata(token)
	task, err := h.svc.CreateTask(ctx, req.ToProto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal")
	}

	res := response.TaskResponseFromProto(task.GetTask())

	return c.JSON(http.StatusCreated, res)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	token, err := getAuthHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	ctx := attachTokenToMetadata(token)

	_, err = h.svc.DeleteTask(ctx, &task.DeleteTaskRequest{Id: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func (h *TaskHandler) Switch(c echo.Context) error {
	id := c.Param("id")
	toQuery := c.QueryParam("switch")
	to, err := strconv.ParseInt(toQuery, 10, 32)
	if err != nil || to > 6 || to < 0 {
		return c.JSON(http.StatusBadRequest, "invalid switch")
	}

	token, err := getAuthHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	ctx := attachTokenToMetadata(token)

	task, err := h.svc.UpdateTaskStatus(ctx, &task.UpdateTaskStatusRequest{Id: id, Status: task.TaskStatus(to)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal")
	}

	return c.JSON(http.StatusOK, response.TaskResponseFromProto(task.GetTask()))
}

func (h *TaskHandler) List(c echo.Context) error {
	paramLimit := c.QueryParam("limit")
	paramOffset := c.QueryParam("offset")
	limit, err := paramToUint32(paramLimit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid limit query value")
	}
	offset, err := paramToUint32(paramOffset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid offset query value")
	}

	token, err := getAuthHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	ctx := attachTokenToMetadata(token)
	stream, err := h.svc.ListTasks(ctx, &task.ListTasksRequest{Limit: limit, Offset: offset})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal")
	}

	res := []response.TaskResponse{}
	for {
		m, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "internal")
		}
		res = append(res, *response.TaskResponseFromProto(m.GetTask()))
	}

	return c.JSON(http.StatusOK, res)
}

func paramToUint32(s string) (uint32, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

func attachTokenToMetadata(token string) context.Context {
	md := metadata.Pairs("jwt", token)
	return metadata.NewOutgoingContext(context.Background(), md)
}
