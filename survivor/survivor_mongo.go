package survivor

import (
	"github.com/somatom98/zssn/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoSurvivorStatusReport struct {
	SID    primitive.ObjectID    `bson:"sid"`
	Status domain.SurvivorStatus `bson:"status"`
}

func (r *MongoSurvivorStatusReport) toDomain() domain.SurvivorStatusReport {
	return domain.SurvivorStatusReport{
		SID:    r.SID.Hex(),
		Status: r.Status,
	}
}

type MongoLocation struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

func (l *MongoLocation) toDomain() domain.Location {
	return domain.Location{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

type MongoSurvivor struct {
	ID            primitive.ObjectID          `bson:"_id,omitempty"`
	Name          string                      `bson:"name"`
	Age           int                         `bson:"age"`
	Gender        domain.SurvivorGender       `bson:"gender"`
	Status        domain.SurvivorStatus       `bson:"status"`
	StatusReports []MongoSurvivorStatusReport `bson:"status_reports",omitempty`
	Location      MongoLocation               `bson:"location"`
}

func (s *MongoSurvivor) toDomain() domain.Survivor {
	statusReports := []domain.SurvivorStatusReport{}
	for _, report := range s.StatusReports {
		statusReports = append(statusReports, report.toDomain())
	}

	return domain.Survivor{
		ID:            s.ID.Hex(),
		Name:          s.Name,
		Age:           s.Age,
		Gender:        s.Gender,
		Status:        s.Status,
		StatusReports: statusReports,
		Location:      s.Location.toDomain(),
	}
}
