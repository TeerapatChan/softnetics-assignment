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

func (h *InventoryHandler) GetItemById(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := h.service.GetItemById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	resp := response.GetItemByIdResponse{
		ID:          item.ID,
		ProductName: item.ProductName,
		Status:      item.Status,
		Price:       item.Price,
		Amount:      item.Amount,
		At:          item.At,
		PNL:         item.PNL,
	}

	return c.JSON(resp)
}

func (h *InventoryHandler) DeleteItemById(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.DeleteItemById(id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(response.DeleteItemByIdResponse{ID: id})
}

func (h *InventoryHandler) UpdateItemById(c *fiber.Ctx) error {
	id := c.Params("id")

	var req request.UpdateItemRequest
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
	}

	if err := h.service.UpdateItemById(id, &item); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(response.UpdateItemByIdResponse{ID: id})
}
