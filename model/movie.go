package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Rating int
	Title  string
	Year   int
	Genre  []Genre
}
