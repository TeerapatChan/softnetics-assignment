package repository

import (
	"strconv"
	"time"

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

func (r *InventoryRepository) DeleteById(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := r.db.Delete(&entities.InventoryItem{}, intID).Error; err != nil {
		return err
	}
	return nil
}

func (r *InventoryRepository) UpdateById(id string, item *entities.InventoryItem) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := r.db.Model(&entities.InventoryItem{}).Where("id = ?", intID).Updates(item).Error; err != nil {
		return err
	}

	return nil
}

func (r *InventoryRepository) FindItemsByProduct(productName string) ([]entities.InventoryItem, error) {
	var items []entities.InventoryItem
	if err := r.db.Where("product_name = ?", productName).Order("at ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InventoryRepository) FindItemsByProductUntil(productName string, until time.Time) ([]entities.InventoryItem, error) {
	var items []entities.InventoryItem
	if err := r.db.Where("product_name = ? AND at <= ?", productName, until).Order("at ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InventoryRepository) FindBoughtItemsSince(productName string, from, to time.Time) ([]entities.InventoryItem, error) {
	var items []entities.InventoryItem
	if err := r.db.Where("product_name = ? AND status = ? AND at >= ? AND at <= ?", productName, entities.BUY, from, to).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InventoryRepository) FindSoldItemsSince(productName string, from, to time.Time) ([]entities.InventoryItem, error) {
	var items []entities.InventoryItem
	if err := r.db.Where("product_name = ? AND status = ? AND at >= ? AND at <= ?", productName, entities.SELL, from, to).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
