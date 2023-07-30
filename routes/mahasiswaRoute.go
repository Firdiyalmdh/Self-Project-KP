package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func MahasiswaRouter(e *echo.Echo) {
	e.GET("/mahasiswa", controllers.GetAllMhs)
	e.GET("/mahasiswa/:id", controllers.GetOneMhs)
	e.POST("/mahasiswa", controllers.CreateMhs)
	e.PUT("/mahasiswa/:id", controllers.EditAMhs)
	e.DELETE("/mahasiswa/:id", controllers.DeleteMhs)
}
