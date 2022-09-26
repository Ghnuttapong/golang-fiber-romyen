package migration

import (
	"gorm.io/gorm"
	"gnutta.com/models"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Prefix{}, &models.Parent{}, &models.ClassRoom{}, &models.Category{}, &models.Finance{}, &models.Transaction{}, &models.Student{})
	if err != nil {
		return err
	}
	return nil
}
