package usecases

import (
	"errors"

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

func (s *InventoryService) GetItemById(id string) (*entities.InventoryItem, error) {
	item, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("item not found")
	}

	return item, nil
}

func (s *InventoryService) DeleteItemById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *InventoryService) UpdateItemById(id string, item *entities.InventoryItem) error {
	return s.repo.UpdateById(id, item)
}
