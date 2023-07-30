package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PermohonanRouter(e *echo.Echo) {
	e.GET("/permohonan", controllers.GetAllPmh)
	e.GET("/permohonan/:id", controllers.GetOnePmh)
	e.POST("/permohonan", controllers.CreatePmh)
	e.PUT("/permohonan/:id", controllers.EditAPmh)
	e.DELETE("/permohonan/:id", controllers.DeletePmh)
}