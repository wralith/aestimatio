package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/request"
)

func (h *AuthHandler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		if len(bearer) == 0 {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}
		token := strings.Split(bearer, " ")
		if token[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		res := request.ValidateRequest{JWT: token[1]}

		isValid := h.callAuthClient(context.Background(), res)
		if !isValid {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		return next(c)
	}
}

func (h *AuthHandler) callAuthClient(ctx context.Context, req request.ValidateRequest) bool {
	isValid, err := h.svc.Validate(ctx, req.ToProto())
	if err != nil || !isValid.Valid {
		return false
	}
	return true
}
