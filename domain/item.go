package domain

import "context"

type Item struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type ItemsRepository interface {
	GetItems(ctx context.Context) ([]Item, error)
	GetItemByName(ctx context.Context, name string) (Item, error)
}
