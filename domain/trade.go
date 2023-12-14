package domain

import "context"

type TradeOffer struct {
	SID   string           `json:"sid"`
	Items map[string]int64 `json:"items"`
}

type TradeService interface {
	Trade(ctx context.Context, offerA TradeOffer, offerB TradeOffer) error
}
