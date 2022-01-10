package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todolist/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/users", uc.Register())
}
