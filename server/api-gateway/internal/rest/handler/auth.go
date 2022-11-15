package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/request"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/response"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rpc"
)

type AuthHandler struct {
	svc rpc.AuthClient
}

func NewAuthHandler(svc rpc.AuthClient) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// @Summary Login
// @ID      Auth-Login
// @Tags    auth
// @Accept  json
// @Produce json
// @Param   credentials body     request.LoginRequest true "User Credentials"
// @Success 200         {object} response.AuthResponse
// @Failure 400
// @Failure 500
// @Router  /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrBadRequest.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrInvalid.Error())
	}

	res, err := h.svc.Login(context.TODO(), req.ToProto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unknown error")
	}

	return c.JSON(http.StatusOK, response.LoginResponseFromProto(res))
}

// @Summary Register
// @ID      Auth-Register
// @Tags    auth
// @Accept  json
// @Produce json
// @Param   credentials body     request.RegisterRequest true "New User Credentials"
// @Success 200         {object} response.AuthResponse
// @Failure 400
// @Failure 500
// @Router  /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	req := request.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.svc.Register(context.TODO(), req.ToProto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unknown error")
	}

	return c.JSON(http.StatusCreated, response.RegisterResponseFromProto(res))
}

func (h *AuthHandler) Validate(c echo.Context) error {
	token, err := getAuthHeader(c)
	if err != nil {
		return err
	}

	res := request.ValidateRequest{JWT: token}

	isValid, err := h.svc.Validate(context.TODO(), res.ToProto())
	if err != nil || !isValid.Valid {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	return nil
}

func getAuthHeader(c echo.Context) (string, error) {
	bearer := c.Request().Header.Get("Authorization")
	if len(bearer) == 0 {
		return "", c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	token := strings.Split(bearer, " ")
	if token[0] != "Bearer" {
		return "", c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	return token[1], nil
}
