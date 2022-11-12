package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r router) initRoutes() {
	r.Echo.POST("/auth/login", r.handler.Login)
	r.Echo.POST("/auth/register", r.handler.Register)

	g := r.Echo.Group("", r.handler.Authenticate)
	g.GET("/restricted", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello")
	})
}
