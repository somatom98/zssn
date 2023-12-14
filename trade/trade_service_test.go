package trade

import (
	"context"
	"errors"
	"testing"

	"github.com/somatom98/zssn/domain"
	"github.com/somatom98/zssn/items"
)

func TestTradeService_Trade(t *testing.T) {
	type fields struct {
		itemsRepository     domain.ItemsRepository
		inventoryRepository domain.InventoryRepository
	}
	type args struct {
		ctx    context.Context
		offerA domain.TradeOffer
		offerB domain.TradeOffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "noitems_error",
			fields: fields{
				itemsRepository: &items.ItemsMockRepository{
					MockGetAllItems: func() ([]domain.Item, error) {
						return []domain.Item{}, errors.New("error")
					},
					MockGetItemByName: func(name string) (domain.Item, error) {
						return domain.Item{}, nil
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				offerA: domain.TradeOffer{},
				offerB: domain.TradeOffer{},
			},
			err: errors.New("error"),
		},
		{
			name: "offeritem_notfound_error",
			fields: fields{
				itemsRepository: items.NewMockRepository(),
			},
			args: args{
				ctx: context.Background(),
				offerA: domain.TradeOffer{
					SID: "survivor",
					Items: map[string]int64{
						"madeup": 4,
					},
				},
				offerB: domain.TradeOffer{},
			},
			err: errors.New(domain.ErrCodeValidation),
		},
		{
			name: "different_prices_error",
			fields: fields{
				itemsRepository: items.NewMockRepository(),
			},
			args: args{
				ctx: context.Background(),
				offerA: domain.TradeOffer{
					SID: "survivor",
					Items: map[string]int64{
						"water": 4,
					},
				},
				offerB: domain.TradeOffer{
					SID: "survivor",
					Items: map[string]int64{
						"water": 1,
					},
				},
			},
			err: errors.New(domain.ErrCodeValidation),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTradeService(tt.fields.itemsRepository, tt.fields.inventoryRepository)
			if err := s.Trade(tt.args.ctx, tt.args.offerA, tt.args.offerB); (tt.err == nil && err != nil) || (tt.err != nil && err == nil) || (tt.err != nil && tt.err.Error() != err.Error()) {
				t.Errorf("TradeService.Trade() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
