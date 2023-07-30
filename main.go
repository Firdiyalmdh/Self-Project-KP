package main

import (
	"golang/configs"
	"golang/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMain() *echo.Echo {
	e := echo.New()
	configs.ConnectDB()
	routes.MahasiswaRouter(e)
	routes.DosenRouter(e)
	routes.PermohonanRouter(e)
	e.Use(middleware.CORS())
	return e
}

func main() {
	e := setupMain()
	e.Logger.Fatal(e.Start(":2500"))
}
