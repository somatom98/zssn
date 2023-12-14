package domain

import "context"

type Inventory struct {
	Items map[string]int64 `json:"items"` // [item]quantity
}

type InventoryRepository interface {
	GetAllInventories(ctx context.Context) ([]Inventory, error)
	GetSurvivorInventory(ctx context.Context, sid string) (Inventory, error)
	AddItem(ctx context.Context, sid string, item string, quantity int64) error
	RemoveItem(ctx context.Context, sid string, item string, quantity int64) error
	TradeItems(ctx context.Context, sidA string, itemsA map[string]int64, sidB string, itemsB map[string]int64) error
}
