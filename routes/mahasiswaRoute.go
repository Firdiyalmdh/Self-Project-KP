package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func MahasiswaRouter(e *echo.Echo) {
	e.GET("/api/mahasiswa", controllers.GetAllMhs)
	e.GET("/api/mahasiswa/:id", controllers.GetOneMhs)
	e.POST("/api/mahasiswa", controllers.CreateMhs)
	e.PUT("/api/mahasiswa/:id", controllers.EditAMhs)
	e.DELETE("/api/mahasiswa/:id", controllers.DeleteMhs)
}
