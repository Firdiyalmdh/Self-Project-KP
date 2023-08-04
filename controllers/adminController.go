package controllers

import (
	"github.com/labstack/echo/v4"
)

func Ad_GetAllPermohonan(c echo.Context) error {
	return GetAllPermohonan(c)
}

func Ad_UpdatePermohonan(c echo.Context) error {
	return UpdatePermohonan(c)
}

func Ad_GetAllPengumpulan(c echo.Context) error {
	return GetAllPengumpulan(c)
}

func Ad_GetAllMahasiswa(c echo.Context) error {
	return GetAllMahasiswa(c)
}

func Ad_GetAllDosen(c echo.Context) error {
	return GetAllDosen(c)
}

func Ad_DeletePermohonan(c echo.Context) error {
	return DeletePermohonan(c)
}