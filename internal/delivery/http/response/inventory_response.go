package response

import (
	"time"

	"github.com/TeerapatChan/inventory-management-api/internal/entities"
)

type CreateItemResponse struct {
	ID uint `json:"id"`
}

type GetItemByIdResponse struct {
	ID          uint            `json:"id"`
	ProductName string          `json:"product_name"`
	Status      entities.Status `json:"status"`
	Price       float64         `json:"price"`
	Amount      int             `json:"amount"`
	At          time.Time       `json:"at"`
	PNL         *float64        `json:"PNL,omitempty"`
}

type DeleteItemByIdResponse struct {
	ID string `json:"id"`
}

type UpdateItemByIdResponse struct {
	ID string `json:"id"`
}

type InventoryItemResponse struct {
	ID          uint            `json:"id"`
	ProductName string          `json:"productName"`
	Status      entities.Status `json:"status"`
	Price       float64         `json:"price"`
	Amount      int             `json:"amount"`
	At          time.Time       `json:"at"`
	PNL         *float64        `json:"PNL,omitempty"`
}

type GetInventoryByProductResponse struct {
	Data                        []InventoryItemResponse `json:"data"`
	TotalAmount                 int                     `json:"totalAmount"`
	ProductsSoldInLatestMonth   int                     `json:"productsSoldInLatestMonth"`
	ProductsBoughtInLatestMonth int                     `json:"productsBoughtInLatestMonth"`
	LatestMonthProfit           float64                 `json:"latestMonthProfit"`
}
