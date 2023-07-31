package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func PengumumanRouter(e *echo.Echo) {
	e.GET("/api/pengumuman", controllers.GetAllPng)
	e.GET("/api/pengumuman/:id", controllers.GetOnePng)
	e.POST("/api/pengumuman", controllers.CreatePng)
	e.PUT("/api/pengumuman/:id", controllers.EditAPng)
	e.DELETE("/api/pengumuman/:id", controllers.DeletePng)
}