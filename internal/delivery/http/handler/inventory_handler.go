package handler

import (
	"github.com/TeerapatChan/inventory-management-api/internal/delivery/http/request"
	"github.com/TeerapatChan/inventory-management-api/internal/delivery/http/response"
	"github.com/TeerapatChan/inventory-management-api/internal/entities"
	"github.com/TeerapatChan/inventory-management-api/internal/usecases"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type InventoryHandler struct {
	service   *usecases.InventoryService
	validator *validator.Validate
}

func NewInventoryHandler(service *usecases.InventoryService, validator *validator.Validate) *InventoryHandler {
	return &InventoryHandler{service, validator}
}

func (h *InventoryHandler) CreateItem(c *fiber.Ctx) error {
	var req request.CreateItemRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	item := entities.InventoryItem{
		ProductName: req.ProductName,
		Status:      req.Status,
		Price:       req.Price,
		Amount:      req.Amount,
		At:          req.At,
		PNL:         0,
	}

	if err := h.service.CreateItem(&item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(response.CreateItemResponse{ID: item.ID})
}
