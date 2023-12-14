package survivor

import (
	"errors"

	"github.com/somatom98/zssn/domain"
)

func ParseSurvivorStatus(value string) (domain.SurvivorStatus, error) {
	switch value {
	case string(domain.SurvivorStatusHealthy):
		return domain.SurvivorStatusHealthy, nil
	case string(domain.SurvivorStatusInfected):
		return domain.SurvivorStatusInfected, nil
	case string(domain.SurvivorStatusDead):
		return domain.SurvivorStatusDead, nil
	default:
		return "", errors.New(domain.ErrCodeParsing)
	}
}
