package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PermohonanRouter(e *echo.Echo) {
	e.GET("/api/permohonan", controllers.GetAllPmh)
	e.GET("/api/permohonan/:id", controllers.GetOnePmh)
	e.POST("/api/permohonan", controllers.CreatePmh)
	e.PUT("/api/permohonan/:id", controllers.EditAPmh)
	e.DELETE("/api/permohonan/:id", controllers.DeletePmh)
}