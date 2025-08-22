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

func (s *InventoryService) CalculatePNL(item *entities.InventoryItem) float64 {
	allItems, err := s.repo.FindItemsByProductUntil(item.ProductName, item.At)
	if err != nil {
		return 0
	}

	var revenue float64
	var currentAverage float64
	var currentTotalItems int64
	var currentSoldItems int64

	for _, p := range allItems {
		if p.Status == entities.BUY {
			currentAverage = (currentAverage*float64(currentTotalItems) + p.Price*float64(p.Amount)) / float64(currentTotalItems+p.Amount)
			currentTotalItems += p.Amount
		}
		if p.Status == entities.SELL {
			currentTotalItems -= p.Amount
			revenue += p.Price * float64(p.Amount)
			currentSoldItems += p.Amount
		}
	}

	if item.Status == entities.BUY {
		return 0
	}

	pnl := revenue - currentAverage*float64(currentSoldItems)
	return pnl
}
