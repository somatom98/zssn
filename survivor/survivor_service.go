package survivor

import (
	"context"

	"github.com/somatom98/zssn/domain"
)

type SurvivorService struct {
	survivorRepository  domain.SurvivorRepository
	inventoryRepository domain.InventoryRepository
}

func NewSurvivorService(survivorRepository domain.SurvivorRepository, inventoryRepository domain.InventoryRepository) *SurvivorService {
	return &SurvivorService{
		survivorRepository:  survivorRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *SurvivorService) GetAllSurvivors(ctx context.Context) ([]domain.Survivor, error) {
	return s.survivorRepository.GetAllSurvivors(ctx)
}

func (s *SurvivorService) GetSurvivor(ctx context.Context, sid string) (domain.Survivor, error) {
	return s.survivorRepository.GetSurvivor(ctx, sid)
}

func (s *SurvivorService) AddSurvivor(ctx context.Context, survivor domain.Survivor) (string, error) {
	return s.survivorRepository.AddSurvivor(ctx, survivor)
}

func (s *SurvivorService) UpdateSurvivorLocation(ctx context.Context, sid string, location domain.Location) error {
	return s.survivorRepository.UpdateSurvivorLocation(ctx, sid, location)
}

func (s *SurvivorService) ReportSurvivorStatus(ctx context.Context, sid string, statusReport domain.SurvivorStatusReport) error {
	survivor, err := s.survivorRepository.GetSurvivor(ctx, sid)
	if err != nil {
		return err
	}

	reports := survivor.StatusReports
	if reports == nil {
		reports = []domain.SurvivorStatusReport{}
	}

	count := 0
	found := false
	for i, report := range reports {
		if report.SID == statusReport.SID {
			reports[i].Status = statusReport.Status
			found = true
		}

		if reports[i].Status == statusReport.Status {
			count++
		}
	}

	if !found {
		reports = append(reports, statusReport)
		count++
	}

	if count >= 3 { // TODO use config
		err = s.survivorRepository.UpdateSurvivorStatus(ctx, sid, statusReport.Status)
		if err != nil {
			return err
		}

		reports = []domain.SurvivorStatusReport{}
	}

	err = s.survivorRepository.UpdateSurvivorStatusReports(ctx, sid, reports)
	if err != nil {
		return err
	}

	return nil
}

func (s *SurvivorService) AddItem(ctx context.Context, sid string, item string, quantity int64) error {
	return s.inventoryRepository.AddItem(ctx, sid, item, quantity)
}

func (s *SurvivorService) RemoveItem(ctx context.Context, sid string, item string, quantity int64) error {
	return s.inventoryRepository.RemoveItem(ctx, sid, item, quantity)
}
