package routes

import (
	"golang/middlewares"

	"github.com/labstack/echo/v4"
)

func HandlerRouter(e *echo.Echo) {
	e.POST("/api/login", middlewares.LoginHandler)
	e.GET("/api/session/:id", middlewares.GetSession)
	e.DELETE("/api/logout/:id", middlewares.LogoutHandler)
}