package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func DosenRouter(e *echo.Echo) {
	e.GET("/api/dosen", controllers.GetAllDsn)
	e.GET("/api/dosen/:id", controllers.GetOneDsn)
	e.POST("/api/dosen", controllers.CreateDsn)
	e.PUT("/api/dosen/:id", controllers.EditADsn)
	e.DELETE("/api/dosen/:id", controllers.DeleteDsn)
}