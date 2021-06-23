package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/refinerydev/go-movie/helper"
)

func PostUserRegistration(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusCreated, "success", "Sucessfully Register!", nil)
	return c.JSON(http.StatusCreated, response)
}

func PostUserLogin(c echo.Context) error {
	response := helper.M{"code": http.StatusOK, "expire": time.Now(), "token": "tokentokentoken"}
	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusOK, "success", "Sucessfully Get Data!", nil)
	return c.JSON(http.StatusOK, response)
}

func PutUser(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusOK, "success", "Sucessfully Updated Data!", nil)
	return c.JSON(http.StatusOK, response)
}
