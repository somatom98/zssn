package domain

type TradeOffer struct {
	SID   string           `json:"sid"`
	Items map[string]int64 `json:"items"`
}

type TradeService interface {
	Trade(offerA TradeOffer, offerB TradeOffer) error
}
