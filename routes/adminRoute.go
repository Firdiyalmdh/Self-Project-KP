package routes

import (
	"golang/controllers"

	"github.com/labstack/echo/v4"
)

func AdminRouter(e *echo.Echo) {
	e.GET("/api/admin/dosen", controllers.GetAllDosen)
	e.GET("/api/admin/mahasiswa", controllers.GetAllMahasiswa)
	e.GET("/api/admin/pengumpulan", controllers.GetAllPengumpulan)
	e.GET("/api/admin/permohonan", controllers.GetAllPermohonan)
	e.PUT("/api/admin/permohonan/:id", controllers.UpdatePermohonan)
	e.DELETE("/api/admin/permohonan/:id", controllers.DeletePermohonan)
}