package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/request"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/response"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rpc"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrInvalid    = errors.New("invalid request")
)

type AuthHandler struct {
	svc rpc.AuthClient
}

func NewAuthHandler(svc rpc.AuthClient) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrBadRequest.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrInvalid.Error())
	}

	res, err := h.svc.Login(context.TODO(), req.ToProto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unknown error")
	}

	return c.JSON(http.StatusOK, response.LoginResponseFromProto(res))
}

func (h *AuthHandler) Register(c echo.Context) error {
	req := request.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(req)
	fmt.Println(req.ToProto())

	res, err := h.svc.Register(context.TODO(), req.ToProto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unknown error")
	}

	return c.JSON(http.StatusCreated, response.RegisterResponseFromProto(res))
}

func (h *AuthHandler) Validate(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	if len(bearer) == 0 {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	token := strings.Split(bearer, " ")
	if token[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	res := request.ValidateRequest{JWT: token[1]}

	isValid, err := h.svc.Validate(context.TODO(), res.ToProto())
	if err != nil || !isValid.Valid {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	return nil
}
