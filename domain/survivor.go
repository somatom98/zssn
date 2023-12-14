package domain

import "context"

type Survivor struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Age           int                    `json:"age"`
	Gender        SurvivorGender         `json:"gender"`
	Status        SurvivorStatus         `json:"status"`
	StatusReports []SurvivorStatusReport `json:"status_reports"`
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

type SurvivorStatusReport struct {
	SID    string         `json:"sid"`
	Status SurvivorStatus `json:"status"`
}

type SurvivorRepository interface {
	GetAllSurvivors(ctx context.Context) ([]Survivor, error)
	GetSurvivor(ctx context.Context, sid string) (Survivor, error)
	AddSurvivor(ctx context.Context, survivor Survivor) (string, error)
	UpdateSurvivorLocation(ctx context.Context, sid string, location Location) error
	UpdateSurvivorStatus(ctx context.Context, sid string, status SurvivorStatus) error
	UpdateSurvivorStatusReports(ctx context.Context, sid string, statusReports []SurvivorStatusReport) error
}

type SurvivorService interface {
	GetAllSurvivors(ctx context.Context) ([]Survivor, error)
	GetSurvivor(ctx context.Context, sid string) (Survivor, error)
	AddSurvivor(ctx context.Context, survivor Survivor) (string, error)
	UpdateSurvivorLocation(ctx context.Context, sid string, location Location) error
	ReportSurvivorStatus(ctx context.Context, sid string, statusReport SurvivorStatusReport) error
	AddItem(ctx context.Context, sid string, item string, quantity int64) error
	RemoveItem(ctx context.Context, sid string, item string, quantity int64) error
}
