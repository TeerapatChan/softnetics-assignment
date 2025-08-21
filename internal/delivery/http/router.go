package http

import (
	"github.com/TeerapatChan/inventory-management-api/internal/delivery/http/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	inventoryGroup := app.Group("/inventory")
	routes.RegisterInventoryRoutes(inventoryGroup, db)
}
