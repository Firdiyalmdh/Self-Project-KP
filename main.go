package main

import (
	"golang/configs"
	"golang/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	configs.ConnectDB()
	routes.MahasiswaRouter(e)
	routes.DosenRouter(e)
	routes.PermohonanRouter(e)
	routes.PengumumanRouter(e)
	routes.PengumpulanRouter(e)
	e.Logger.Fatal(e.Start(":2500"))
}
