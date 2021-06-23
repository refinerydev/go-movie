package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  int
	MovieID int
	Review  string
	Rate    int
}
