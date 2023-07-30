package main

import (
	"golang/configs"
	"golang/routes"

	"github.com/labstack/echo/v4"
)

func setupMain() *echo.Echo {
	e := echo.New()
	configs.ConnectDB()
	routes.MahasiswaRouter(e)
	routes.DosenRouter(e)
	routes.PermohonanRouter(e)
	return e
}

func main() {
	e := setupMain()
	e.Logger.Fatal(e.Start(":2500"))
}
