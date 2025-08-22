package routes

import (
	"github.com/TeerapatChan/inventory-management-api/internal/delivery/http/handler"
	"github.com/TeerapatChan/inventory-management-api/internal/repository"
	"github.com/TeerapatChan/inventory-management-api/internal/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterInventoryRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewInventoryRepository(db)
	service := usecases.NewInventoryService(repo)

	validator := validator.New()
	h := handler.NewInventoryHandler(service, validator)

	router.Get("/:name", h.GetItemSummaryByName)

	items := router.Group("/items")

	items.Post("/", h.CreateItem)
	items.Get("/:id", h.GetItemById)
	items.Delete("/:id", h.DeleteItemById)
	items.Patch("/:id", h.UpdateItemById)
}
