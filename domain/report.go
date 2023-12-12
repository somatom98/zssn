package domain

import "context"

type Report struct {
	Statuses             map[SurvivorStatus]float32 `json:"statuses"` // % of survivors grouped by status
	AverageItemsQuantity map[string]float32         `json:"avg_items_quantity`
	PointsLost           int64                      `json:"points_lost"`
}

type ReportService interface {
	GetReport(ctx context.Context) (Report, error)
}
