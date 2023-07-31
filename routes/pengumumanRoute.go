package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PengumumanRouter(e *echo.Echo) {
	e.GET("/api/pengumuman", controllers.GetAllPengumuman)
	e.GET("/api/pengumuman/:id", controllers.GetPengumuman)
	e.POST("/api/pengumuman", controllers.CreatePengumuman)
	e.PUT("/api/pengumuman/:id", controllers.UpdatePengumuman)
	e.DELETE("/api/pengumuman/:id", controllers.DeletePengumuman)
}