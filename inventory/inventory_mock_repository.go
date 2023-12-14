package inventory

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type InventoryMockRepository struct {
	inventories map[string]domain.Inventory
}

func NewMockRepository() *InventoryMockRepository {
	return &InventoryMockRepository{
		inventories: map[string]domain.Inventory{
			"657b4ea4d54e4b7c3870f8c3": {
				Items: map[string]int64{
					"water":      4,
					"food":       25,
					"medication": 13,
					"ammonition": 299,
				},
			},
		},
	}
}

func (r *InventoryMockRepository) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	inventories := []domain.Inventory{}

	for _, inventory := range r.inventories {
		inventories = append(inventories, inventory)
	}
	return inventories, nil
}

func (r *InventoryMockRepository) GetSurvivorInventory(ctx context.Context, sid string) (domain.Inventory, error) {
	inventory, ok := r.inventories[sid]
	if !ok {
		return domain.Inventory{}, nil
	}

	return inventory, nil
}

func (r *InventoryMockRepository) AddItem(ctx context.Context, sid string, item string, quantity int64) error {
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

func (r *InventoryMockRepository) RemoveItem(ctx context.Context, sid string, item string, quantity int64) error {
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
