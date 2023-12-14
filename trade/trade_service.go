package trade

import "github.com/somatom98/zssn/domain"

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
	return nil
}
