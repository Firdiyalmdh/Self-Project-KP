package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func DosenRouter(e *echo.Echo) {
	e.GET("/dosen", controllers.GetAllDsn)
	e.GET("/dosen/:id", controllers.GetOneDsn)
	e.POST("/dosen", controllers.CreateDsn)
	e.PUT("/dosen/:id", controllers.EditADsn)
	e.DELETE("/dosen/:id", controllers.DeleteDsn)
}