package inventory

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type InventoryMockRepository struct {
	MockGetAllInventories    func() ([]domain.Inventory, error)
	MockGetSurvivorInventory func(sid string) (domain.Inventory, error)
	MockAddItem              func(sid string, item string, quantity int64) error
	MockRemoveItem           func(sid string, item string, quantity int64) error
	MockTradeItems           func(sidA string, itemsA map[string]int64, sidB string, itemsB map[string]int64) error
}

func (r *InventoryMockRepository) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	return r.MockGetAllInventories()
}

func (r *InventoryMockRepository) GetSurvivorInventory(ctx context.Context, sid string) (domain.Inventory, error) {
	return r.MockGetSurvivorInventory(sid)
}

func (r *InventoryMockRepository) AddItem(ctx context.Context, sid string, item string, quantity int64) error {
	return r.MockAddItem(sid, item, quantity)
}

func (r *InventoryMockRepository) RemoveItem(ctx context.Context, sid string, item string, quantity int64) error {
	return r.MockRemoveItem(sid, item, quantity)
}

func (r *InventoryMockRepository) TradeItems(ctx context.Context, sidA string, itemsA map[string]int64, sidB string, itemsB map[string]int64) error {
	return r.MockTradeItems(sidA, itemsA, sidB, itemsB)
}

type InventoryFakeRepository struct {
	inventories map[string]domain.Inventory
}

func NewMockRepository() *InventoryFakeRepository {
	return &InventoryFakeRepository{
		inventories: map[string]domain.Inventory{
			"657b4ea4d54e4b7c3870f8c3": {
				Items: map[string]int64{
					"water":      4,
					"food":       25,
					"medication": 13,
					"ammunition": 299,
				},
			},
			"657b4ea4d54e4b7c3870f8c7": {
				Items: map[string]int64{
					"ammunition": 57,
				},
			},
		},
	}
}

func (r *InventoryFakeRepository) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	inventories := []domain.Inventory{}

	for _, inventory := range r.inventories {
		inventories = append(inventories, inventory)
	}
	return inventories, nil
}

func (r *InventoryFakeRepository) GetSurvivorInventory(ctx context.Context, sid string) (domain.Inventory, error) {
	inventory, ok := r.inventories[sid]
	if !ok {
		return domain.Inventory{}, nil
	}

	return inventory, nil
}

func (r *InventoryFakeRepository) AddItem(ctx context.Context, sid string, item string, quantity int64) error {
	inventory, ok := r.inventories[sid]
	if !ok {
		inventory = domain.Inventory{
			Items: map[string]int64{},
		}
		r.inventories[sid] = inventory
	}

	currentQuantity := inventory.Items[item]
	currentQuantity += quantity
	r.inventories[sid].Items[item] = currentQuantity

	return nil
}

func (r *InventoryFakeRepository) RemoveItem(ctx context.Context, sid string, item string, quantity int64) error {
	inventory, ok := r.inventories[sid]
	if !ok {
		return errors.New(domain.ErrCodeNotFound)
	}

	currentQuantity, ok := inventory.Items[item]
	if !ok {
		return errors.New(domain.ErrCodeNotFound)
	}

	if quantity > currentQuantity {
		return errors.New(domain.ErrCodeValidation)
	}

	r.inventories[sid].Items[item] -= quantity

	return nil
}

func (r *InventoryFakeRepository) TradeItems(ctx context.Context, sidA string, itemsA map[string]int64, sidB string, itemsB map[string]int64) error {
	// in the real repository we would create a transaction with a rollback logic

	err := r.swapItems(ctx, sidA, sidB, itemsA)
	if err != nil {
		return err
	}

	r.swapItems(ctx, sidB, sidA, itemsB)
	if err != nil {
		return err
	}

	return nil
}

func (r *InventoryFakeRepository) swapItems(ctx context.Context, giver string, receiver string, items map[string]int64) error {
	for item, quantity := range items {
		err := r.RemoveItem(ctx, giver, item, quantity)
		if err != nil {
			return err
		}

		err = r.AddItem(ctx, receiver, item, quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
