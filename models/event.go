package models

import (
	"context"
	"time"

	"log"

	"gorm.io/gorm"
)

type Event struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}

func Seed(db *gorm.DB) {
	events := []Event{
		{ID: 1, Name: "Tech Conference 2025", Location: "San Francisco", Date: time.Now().AddDate(0, 1, 0)},
	}

	for _, event := range events {
		// Check if the event already exists to avoid duplicates
		var existing Event
		if err := db.First(&existing, event.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Insert new event
				if err := db.Create(&event).Error; err != nil {
					log.Printf("Failed to seed event: %v", err)
				} else {
					log.Printf("Seeded event: %s", event.Name)
				}
			}
		}
	}
}
