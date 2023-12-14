package survivor

import (
	"context"
	"errors"

	"github.com/somatom98/zssn/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SurvivorMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *SurvivorMongoRepository {
	return &SurvivorMongoRepository{
		db:         db,
		collection: db.Collection("survivors"),
	}
}

func (r *SurvivorMongoRepository) GetAllSurvivors(ctx context.Context) ([]domain.Survivor, error) {
	survivors := []domain.Survivor{}

	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return survivors, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var elem MongoSurvivor
		err := cur.Decode(&elem)
		if err != nil {
			return survivors, err
		}

		survivors = append(survivors, elem.toDomain())
	}

	if err := cur.Err(); err != nil {
		return survivors, err
	}

	return survivors, nil
}

func (r *SurvivorMongoRepository) GetSurvivor(ctx context.Context, sid string) (domain.Survivor, error) {
	mongoSurvivor := MongoSurvivor{}

	objectID, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return domain.Survivor{}, errors.New(domain.ErrCodeParsing)
	}

	err = r.collection.FindOne(ctx, primitive.M{"_id": objectID}).Decode(&mongoSurvivor)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return domain.Survivor{}, errors.New(domain.ErrCodeNotFound)
		}
		return domain.Survivor{}, err
	}

	return mongoSurvivor.toDomain(), nil
}

func (r *SurvivorMongoRepository) AddSurvivor(ctx context.Context, survivor domain.Survivor) (string, error) {
	statusReports := []MongoSurvivorStatusReport{}
	for _, report := range survivor.StatusReports {
		sid, err := primitive.ObjectIDFromHex(report.SID)
		if err != nil {
			return "", err
		}

		mongoReport := MongoSurvivorStatusReport{
			SID:    sid,
			Status: report.Status,
		}

		statusReports = append(statusReports, mongoReport)
	}

	mongoSurvivor := MongoSurvivor{
		Name:          survivor.Name,
		Age:           survivor.Age,
		Gender:        survivor.Gender,
		Status:        survivor.Status,
		StatusReports: statusReports,
		Location: MongoLocation{
			Latitude:  survivor.Location.Latitude,
			Longitude: survivor.Location.Longitude,
		},
	}

	result, err := r.collection.InsertOne(ctx, mongoSurvivor)
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New(domain.ErrCodeParsing)
	}

	return insertedID.Hex(), nil
}

func (r *SurvivorMongoRepository) UpdateSurvivorLocation(ctx context.Context, sid string, location domain.Location) error {
	return errors.New(domain.ErrCodeNotFound)
}

func (r *SurvivorMongoRepository) UpdateSurvivorStatus(ctx context.Context, sid string, status domain.SurvivorStatus) error {
	return errors.New(domain.ErrCodeNotFound)
}

func (r *SurvivorMongoRepository) UpdateSurvivorStatusReports(ctx context.Context, sid string, statusReports []domain.SurvivorStatusReport) error {
	return errors.New(domain.ErrCodeNotFound)
}
