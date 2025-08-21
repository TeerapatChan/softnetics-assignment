package main

import (
	"log"

	httpd "github.com/TeerapatChan/inventory-management-api/internal/delivery/http"
	"github.com/TeerapatChan/inventory-management-api/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	httpd.RegisterRoutes(app, db)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
