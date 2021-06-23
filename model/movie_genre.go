package model

import (
	"gorm.io/gorm"
)

type MovieGenre struct {
	gorm.Model
	MovieID int
	GenreID int
}
