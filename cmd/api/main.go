package main

import (
	"fmt"

	"github.com/QwertyAkane/ticket-booking-project/config"
	"github.com/QwertyAkane/ticket-booking-project/db"
	"github.com/QwertyAkane/ticket-booking-project/handlers"
	"github.com/QwertyAkane/ticket-booking-project/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repository
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
