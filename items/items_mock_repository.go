package items

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type ItemsMockRepository struct {
	MockGetAllItems   func() ([]domain.Item, error)
	MockGetItemByName func(name string) (domain.Item, error)
}

func (r *ItemsMockRepository) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	return r.MockGetAllItems()
}

func (r *ItemsMockRepository) GetItemByName(ctx context.Context, name string) (domain.Item, error) {
	return r.MockGetItemByName(name)
}

func NewMockRepository() *ItemsMockRepository {
	getAllItems := func() ([]domain.Item, error) {
		return []domain.Item{
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
		}, nil
	}

	getItemByName := func(name string) (domain.Item, error) {
		items, _ := getAllItems()
		for _, item := range items {
			if item.Name == name {
				return item, nil
			}
		}
		return domain.Item{}, errors.New(domain.ErrCodeNotFound)
	}

	return &ItemsMockRepository{
		MockGetAllItems:   getAllItems,
		MockGetItemByName: getItemByName,
	}
}
