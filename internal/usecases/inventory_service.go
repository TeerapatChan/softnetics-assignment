package usecases

import (
	"errors"
	"time"

	"github.com/TeerapatChan/inventory-management-api/internal/delivery/http/response"
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

	revenue := 0.0
	currentAverage := 0.0
	currentTotalItems := 0
	currentSoldItems := 0

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

func (s *InventoryService) GetItemSummaryByName(productName string) ([]response.InventoryItemResponse, int, int, int, float64, error) {
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	allItems, err := s.repo.FindItemsByProduct(productName)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	totalAmount := 0
	for _, item := range allItems {
		totalAmount += item.Amount
	}

	data := make([]response.InventoryItemResponse, 0, len(allItems))
	for _, item := range allItems {
		var pnl *float64
		if item.Status == entities.SELL {
			calculatedPNL := s.CalculatePNL(&item)
			pnl = &calculatedPNL
		}
		data = append(data, response.InventoryItemResponse{
			ID:          item.ID,
			ProductName: item.ProductName,
			Status:      item.Status,
			Price:       item.Price,
			Amount:      item.Amount,
			At:          item.At,
			PNL:         pnl,
		})
	}

	boughtThisMonth, err := s.repo.FindBoughtItemsSince(productName, monthStart, now)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	soldThisMonth, err := s.repo.FindSoldItemsSince(productName, monthStart, now)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	cost := 0.0
	productsBoughtLatestMonth := 0
	for _, b := range boughtThisMonth {
		productsBoughtLatestMonth += b.Amount
		cost += b.Price * float64(b.Amount)
	}

	revenue := 0.0
	productsSoldLatestMonth := 0
	for _, s := range soldThisMonth {
		productsSoldLatestMonth += s.Amount
		revenue += s.Price * float64(s.Amount)
	}

	latestMonthProfit := revenue - cost
	return data, totalAmount, productsBoughtLatestMonth, productsSoldLatestMonth, latestMonthProfit, nil
}
