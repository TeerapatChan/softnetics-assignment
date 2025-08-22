package repository

import (
	"strconv"

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

func (r *InventoryRepository) FindById(id string) (*entities.InventoryItem, error) {
	var item entities.InventoryItem

	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.First(&item, intID).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
