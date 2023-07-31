package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PengumpulanRouter(e *echo.Echo) {
	e.GET("/api/pengumpulan", controllers.GetAllPengumpulan)
	e.GET("/api/pengumpulan/:id", controllers.GetPengumpulan)
	e.POST("/api/pengumpulan", controllers.CreatePengumpulan)
	e.PUT("/api/pengumpulan/:id", controllers.UpdatePengumpulan)
	e.DELETE("/api/pengumpulan/:id", controllers.DeletePengumpulan)
}