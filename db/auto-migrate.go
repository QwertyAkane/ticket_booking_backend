package db

import (
	"github.com/QwertyAkane/ticket-booking-project/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	db.Migrator().DropTable(&models.Event{})
	return db.AutoMigrate(&models.Event{})
}
