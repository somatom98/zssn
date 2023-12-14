package trade

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type TradeService struct {
	itemsRepository     domain.ItemsRepository
	inventoryRepository domain.InventoryRepository
}

func NewTradeService(itemsRepository domain.ItemsRepository, inventoryRepository domain.InventoryRepository) *TradeService {
	return &TradeService{
		itemsRepository:     itemsRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *TradeService) Trade(ctx context.Context, offerA domain.TradeOffer, offerB domain.TradeOffer) error {
	items, err := s.itemsRepository.GetAllItems(ctx)
	if err != nil {
		return err
	}

	valueA, err := s.evaluate(offerA, items)
	if err != nil {
		return err
	}

	valueB, err := s.evaluate(offerB, items)
	if err != nil {
		return err
	}

	if valueA != valueB {
		return errors.New(domain.ErrCodeValidation)
	}

	err = s.checkItemsAvailability(ctx, offerA)
	if err != nil {
		return err
	}

	err = s.checkItemsAvailability(ctx, offerB)
	if err != nil {
		return err
	}

	return s.inventoryRepository.TradeItems(ctx, offerA.SID, offerA.Items, offerB.SID, offerB.Items)
}

func (s *TradeService) evaluate(offer domain.TradeOffer, items []domain.Item) (int64, error) {
	var value int64 = 0
	for name, quantity := range offer.Items {
		var points int64 = 0

		for _, item := range items {
			if item.Name == name {
				points = int64(item.Points)
				break
			}
		}

		if points == 0 {
			return 0, errors.New(domain.ErrCodeValidation)
		}

		value += quantity * points
	}
	return value, nil
}

func (s *TradeService) checkItemsAvailability(ctx context.Context, offer domain.TradeOffer) error {
	inventory, err := s.inventoryRepository.GetSurvivorInventory(ctx, offer.SID)
	if err != nil {
		return err
	}

	for item, quantity := range offer.Items {
		availableQuantity := inventory.Items[item]
		if availableQuantity < quantity {
			return errors.New(domain.ErrCodeValidation)
		}
	}

	return nil
}
