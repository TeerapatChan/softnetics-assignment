package usecases

import (
	"github.com/TeerapatChan/inventory-management-api/internal/entities"
	"github.com/TeerapatChan/inventory-management-api/internal/repository"
)

type InventoryService struct {
	repo *repository.InventoryRepository
}

func NewInventoryService(repo *repository.InventoryRepository) *InventoryService {
	return &InventoryService{repo}
}

func (s *InventoryService) CreateItem(item *entities.InventoryItem) error {
	// Mock PNL calculation
	item.PNL = 0
	return s.repo.Save(item)
}
