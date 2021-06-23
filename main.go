package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/refinerydev/go-movie/database"
	"github.com/refinerydev/go-movie/handler"
	"github.com/refinerydev/go-movie/helper"
	"github.com/refinerydev/go-movie/model"
)

func main() {
	e := echo.New()

	db := database.GetDbInstance()
	err := db.AutoMigrate(
		model.User{},
		model.Movie{},
		model.Genre{},
		model.Review{},
		model.MovieGenre{},
	)
	if err != nil {
		fmt.Print("Migration failed")
	}

	// Custom Validator
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	// Main root
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	e.POST("/register", handler.PostUserRegistration)

	e.POST("/signin", handler.PostUserLogin)

	e.GET("/movie_reviews/user", handler.GetUser)

	e.PUT("/movie_reviews/user", handler.PutUser)

	e.POST("/movie_reviews/genre", handler.PostGenre)

	e.GET("/movie_reviews/genre", handler.GetGenres)

	e.POST("/movie_reviews/movies", handler.PostMovie)

	e.GET("/movie_reviews/movies", handler.GetMovies)

	e.POST("/movie_reviews/movies/genre", handler.PostGenreMovie)

	e.POST("/movie_reviews/review", handler.PostReview)

	e.GET("/movie_reviews/review", handler.GetReviews)
	// Run server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
