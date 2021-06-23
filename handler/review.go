package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/refinerydev/go-movie/helper"
)

func PostReview(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusCreated, "success", "Sucessfully Add Reviews For this Movie!", nil)
	return c.JSON(http.StatusCreated, response)
}

func GetReviews(c echo.Context) error {
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully Get Review List!", nil)
	return c.JSON(http.StatusOK, response)
}
