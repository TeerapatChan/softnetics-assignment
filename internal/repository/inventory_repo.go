package repository

import (
	"github.com/TeerapatChan/inventory-management-api/internal/entities"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db}
}

func (r *InventoryRepository) Save(item *entities.InventoryItem) error {
	if err := r.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}
