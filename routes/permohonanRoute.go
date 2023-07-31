package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PermohonanRouter(e *echo.Echo) {
	e.GET("/api/permohonan", controllers.GetAllPermohonan)
	e.GET("/api/permohonan/:id", controllers.GetPermohonan)
	e.POST("/api/permohonan", controllers.CreatePermohonan)
	e.PUT("/api/permohonan/:id", controllers.UpdatePermohonan)
	e.DELETE("/api/permohonan/:id", controllers.DeletePermohonan)
}