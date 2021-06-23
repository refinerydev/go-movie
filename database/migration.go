package database

import (
	"errors"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/refinerydev/go-movie/model"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// ID: fmt.Sprintf("%d", time.Now().Unix()),
			ID: "1618983623",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(model.User{}).Error; err != nil {
					return errors.New("Migration failed")
				}
				return nil
			},
		},
	})
}
