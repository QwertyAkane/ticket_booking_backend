package models

import (
	"context"
	"fmt"
	"hash/fnv"
	"time"
)

type Event struct {
	ID        uint      `json:"id" gorm:"primarykey"`
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

// GenerateID creates a unique uint hash based on Name, Location, and Date
func GenerateID(name, location string, date time.Time) uint {
	h := fnv.New32a()
	_, _ = h.Write([]byte(fmt.Sprintf("%s-%s-%d", name, location, date.Unix())))
	return uint(h.Sum32()) // Convert hash to uint
}
