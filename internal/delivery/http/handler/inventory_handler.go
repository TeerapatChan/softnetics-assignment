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

// @Summary		Create Item
// @Description	Create a new inventory item
// @Tags			inventory
// @Router			/inventory/items [post]
// @Success		201			{object}	response.CreateItemResponse
// @Failure		400			{object}	map[string]interface{}
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
	}

	if err := h.service.CreateItem(&item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(response.CreateItemResponse{ID: item.ID})
}

// @Summary		Get Item By ID
// @Description	Get an inventory item by its ID
// @Tags			inventory
// @Router			/inventory/items/{id} [get]
// @Success		200			{object}	response.GetItemByIdResponse
// @Failure		404			{object}	map[string]interface{}
func (h *InventoryHandler) GetItemById(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := h.service.GetItemById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	var pnl *float64
	if item.Status != entities.BUY {
		calculatedPNL := h.service.CalculatePNL(item)
		pnl = &calculatedPNL
	}

	resp := response.GetItemByIdResponse{
		ID:          item.ID,
		ProductName: item.ProductName,
		Status:      item.Status,
		Price:       item.Price,
		Amount:      item.Amount,
		At:          item.At,
		PNL:         pnl,
	}

	return c.JSON(resp)
}

// @Summary		Delete Item By ID
// @Description	Delete an inventory item by its ID
// @Tags			inventory
// @Router			/inventory/items/{id} [delete]
// @Success		200			{object}	response.DeleteItemByIdResponse
// @Failure		400			{object}	map[string]interface{}
// @Failure		404			{object}	map[string]interface{}
func (h *InventoryHandler) DeleteItemById(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.DeleteItemById(id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(response.DeleteItemByIdResponse{ID: id})
}

// @Summary		Update Item By ID
// @Description	Update an inventory item by its ID
// @Tags			inventory
// @Router			/inventory/items/{id} [put]
// @Success		200			{object}	response.UpdateItemByIdResponse
// @Failure		400			{object}	map[string]interface{}
// @Failure		404			{object}	map[string]interface{}
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

// @Summary		Get Item Summary By Name
// @Description	Get a summary of inventory items by product name
// @Tags			inventory
// @Router			/inventory/{productName} [get]
// @Success		200			{object}	response.GetInventoryByProductResponse
// @Failure		404			{object}	map[string]interface{}
func (h *InventoryHandler) GetItemSummaryByProductName(c *fiber.Ctx) error {
	productName := c.Params("productName")

	items, totalAmount, productsBoughtInLatestMonth, productsSoldInLatestMonth, latestMonthProfit, err := h.service.GetItemSummaryByProductName(productName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response.GetInventoryByProductResponse{
		Data:                        items,
		TotalAmount:                 totalAmount,
		ProductsBoughtInLatestMonth: productsBoughtInLatestMonth,
		ProductsSoldInLatestMonth:   productsSoldInLatestMonth,
		LatestMonthProfit:           latestMonthProfit,
	})

}
