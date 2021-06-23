package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/refinerydev/go-movie/helper"
)

func GetMovies(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusOK, "success", "Sucessfully Added Data!", nil)
	return c.JSON(http.StatusOK, response)
}

func PostMovie(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusCreated, "success", "Sucessfully Created Data!", nil)
	return c.JSON(http.StatusCreated, response)
}

func PostGenre(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusCreated, "success", "Sucessfully Added Data!", nil)
	return c.JSON(http.StatusCreated, response)
}

func GetGenres(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusOK, "success", "Sucessfully Get Genre List!", nil)
	return c.JSON(http.StatusOK, response)
}

func PostGenreMovie(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusCreated, "success", "Sucessfully Added Data!", nil)
	return c.JSON(http.StatusCreated, response)
}
