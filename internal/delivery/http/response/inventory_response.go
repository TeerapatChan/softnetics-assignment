package response

import "github.com/TeerapatChan/inventory-management-api/internal/entities"

type CreateItemResponse struct {
	ID uint `json:"id"`
}

type GetItemByIdResponse = entities.InventoryItem
