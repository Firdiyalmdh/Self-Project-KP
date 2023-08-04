package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func AdminRouter(e *echo.Echo) {
	e.GET("/api/admin/dosen", controllers.Ad_GetAllDosen)
	e.GET("/api/admin/mahasiswa", controllers.Ad_GetAllMahasiswa)
	e.GET("/api/admin/pengumpulan", controllers.Ad_GetAllPengumpulan)
	e.GET("/api/admin/permohonan", controllers.Ad_GetAllPermohonan)
	e.PUT("/api/admin/permohonan/:id", controllers.Ad_UpdatePermohonan)
	e.DELETE("/api/admin/permohonan/:id", controllers.Ad_DeletePermohonan)
}