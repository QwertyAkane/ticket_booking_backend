package db

import (
	"github.com/QwertyAkane/ticket-booking-project/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
