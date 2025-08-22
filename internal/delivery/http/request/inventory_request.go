package request

import (
	"time"

	"github.com/TeerapatChan/inventory-management-api/internal/entities"
)

type CreateItemRequest struct {
	ProductName string          `json:"productName" validate:"required"`
	Status      entities.Status `json:"status" validate:"required,oneof=BUY SELL"`
	Price       float64         `json:"price" validate:"required,min=0"`
	Amount      int             `json:"amount" validate:"required,min=0"`
	At          time.Time       `json:"at" validate:"required"`
}

type UpdateItemRequest struct {
	ProductName string          `json:"productName"`
	Status      entities.Status `json:"status" validate:"oneof=BUY SELL"`
	Price       float64         `json:"price" validate:"min=0"`
	Amount      int             `json:"amount" validate:"min=0"`
	At          time.Time       `json:"at"`
}
