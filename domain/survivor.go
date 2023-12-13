package domain

import "context"

type Survivor struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Age           int                    `json:"age"`
	Gender        SurvivorGender         `json:"gender"`
	Status        SurvivorStatus         `json:"status"`
	StatusReports map[SurvivorStatus]int `json:"status_reports"`
	Location      Location               `json:"location"`
	Inventory     Inventory              `json:"inventory"`
}

type SurvivorGender string

const (
	SurvivorGenderMale   SurvivorGender = "male"
	SurvivorGenderFemale SurvivorGender = "female"
	SurvivorGendeerOther SurvivorGender = "other"
)

type SurvivorStatus string

const (
	SurvivorStatusHealthy  SurvivorStatus = "healthy"
	SurvivorStatusInfected SurvivorStatus = "infected"
	SurvivorStatusDead     SurvivorStatus = "dead"
)

type SurvivorRepository interface {
	GetAllSurvivors(ctx context.Context) ([]Survivor, error)
	GetSurvivor(ctx context.Context, sid string) (Survivor, error)
	AddSurvivor(ctx context.Context, survivor Survivor) (string, error)
	UpdateSurvivorLocation(ctx context.Context, sid string, location Location) error
	UpdateSurvivorStatus(ctx context.Context, sid string, status SurvivorStatus) error
	UpdateSurvivorStatusReports(ctx context.Context, sid string, statusReports map[SurvivorStatus]int) error
}

type SurvivorService interface {
	GetAllSurvivors(ctx context.Context) ([]Survivor, error)
	GetSurvivor(ctx context.Context, sid string) (Survivor, error)
	AddSurvivor(ctx context.Context, survivor Survivor) (string, error)
	UpdateSurvivorLocation(ctx context.Context, sid string, location Location) error
	ReportSurvivorStatus(ctx context.Context, sid string, status SurvivorStatus) error
	AddItem(ctx context.Context, sid string, item string, quantity int64) error
	RemoveItem(ctx context.Context, sid string, item string, quantity int64) error
}
