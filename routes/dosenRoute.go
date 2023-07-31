package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func DosenRouter(e *echo.Echo) {
	e.GET("/api/dosen", controllers.GetAllDosen)
	e.GET("/api/dosen/:id", controllers.GetDosen)
	e.POST("/api/dosen", controllers.CreateDosen)
	e.PUT("/api/dosen/:id", controllers.UpdateDosen)
	e.DELETE("/api/dosen/:id", controllers.DeleteDosen)
}