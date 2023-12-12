package items

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type ItemsMockRepository struct {
	items []domain.Item
}

func NewMockRepository() *ItemsMockRepository {
	return &ItemsMockRepository{
		items: []domain.Item{
			{
				Name:   "water",
				Points: 4,
			},
			{
				Name:   "food",
				Points: 3,
			},
			{
				Name:   "medication",
				Points: 2,
			},
			{
				Name:   "ammunition",
				Points: 1,
			},
		},
	}
}

func (r *ItemsMockRepository) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	return r.items, nil
}

func (r *ItemsMockRepository) GetItemByName(ctx context.Context, name string) (domain.Item, error) {
	for _, item := range r.items {
		if item.Name == name {
			return item, nil
		}
	}
	return domain.Item{}, errors.New(domain.ErrCodeNotFound)
}
