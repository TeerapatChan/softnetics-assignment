package entities

import "time"

type Status string

const (
	BUY  Status = "BUY"
	SELL Status = "SELL"
)

type InventoryItem struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	ProductName string    `json:"productName"`
	Status      Status    `json:"status"`
	Price       float64   `json:"price"`
	Amount      int64     `json:"amount"`
	At          time.Time `json:"at"`
	PNL         float64   `json:"PNL,omitempty"`
}
