package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/handler"
	"github.com/wralith/aestimatio/server/api-gateway/pkg/vld"
)

type router struct {
	Echo        *echo.Echo
	authHandler *handler.AuthHandler
	taskHandler *handler.TaskHandler
}

func New(authH *handler.AuthHandler, taskH *handler.TaskHandler) *router {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	e.Validator = &vld.Validator{Validator: validator.New()}

	r := &router{Echo: e, authHandler: authH, taskHandler: taskH}
	r.initRoutes()

	return r
}
