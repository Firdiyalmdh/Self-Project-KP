package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func MahasiswaRouter(e *echo.Echo) {
	e.GET("/api/mahasiswa", controllers.GetAllMahasiswa)
	e.GET("/api/mahasiswa/:id", controllers.GetMahasiswa)
	e.POST("/api/mahasiswa", controllers.CreateMahasiswa)
	e.PUT("/api/mahasiswa/:id", controllers.UpdateMahasiswa)
	e.DELETE("/api/mahasiswa/:id", controllers.DeleteMahasiswa)
}
