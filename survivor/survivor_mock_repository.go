package survivor

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
)

type SurvivorMockRepository struct {
	survivors []domain.Survivor
}

func NewMockRepository() *SurvivorMockRepository {
	return &SurvivorMockRepository{
		survivors: []domain.Survivor{
			{
				ID:     "survivor",
				Name:   "Tommaso",
				Age:    25,
				Gender: domain.SurvivorGenderMale,
				Location: domain.Location{
					Latitude:  15,
					Longitude: 15,
				},
				Inventory: domain.Inventory{
					Items: map[string]int64{
						"water": 1,
					},
				},
			},
		},
	}
}

func (r *SurvivorMockRepository) GetAllSurvivors(ctx context.Context) ([]domain.Survivor, error) {
	return r.survivors, nil
}

func (r *SurvivorMockRepository) GetSurvivor(ctx context.Context, sid string) (domain.Survivor, error) {
	for _, s := range r.survivors {
		if s.ID == sid {
			return s, nil
		}
	}

	return domain.Survivor{}, errors.New(domain.ErrCodeNotFound)
}

func (r *SurvivorMockRepository) AddSurvivor(ctx context.Context, survivor domain.Survivor) (string, error) {
	for _, s := range r.survivors {
		if s.ID == survivor.ID {
			return "", errors.New(domain.ErrCodeDuplicate)
		}
	}

	r.survivors = append(r.survivors, survivor)

	return survivor.ID, nil
}

func (r *SurvivorMockRepository) UpdateSurvivorLocation(ctx context.Context, sid string, location domain.Location) error {
	for i, s := range r.survivors {
		if s.ID == sid {
			r.survivors[i].Location = location
			return nil
		}
	}

	return errors.New(domain.ErrCodeNotFound)
}

func (r *SurvivorMockRepository) UpdateSurvivorStatus(ctx context.Context, sid string, status domain.SurvivorStatus) error {
	for i, s := range r.survivors {
		if s.ID == sid {
			r.survivors[i].Status = status
			return nil
		}
	}

	return errors.New(domain.ErrCodeNotFound)
}

func (r *SurvivorMockRepository) UpdateSurvivorStatusReports(ctx context.Context, sid string, statusReports map[domain.SurvivorStatus]int) error {
	for i, s := range r.survivors {
		if s.ID == sid {
			r.survivors[i].StatusReports = statusReports
			return nil
		}
	}

	return errors.New(domain.ErrCodeNotFound)
}
