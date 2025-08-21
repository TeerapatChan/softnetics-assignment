package handler

import (
	"github.com/TeerapatChan/inventory-management-api/internal/usecases"
	"github.com/go-playground/validator/v10"
)

type InventoryHandler struct {
	service   *usecases.InventoryService
	validator *validator.Validate
}

func NewInventoryHandler(service *usecases.InventoryService, validator *validator.Validate) *InventoryHandler {
	return &InventoryHandler{service, validator}
}
