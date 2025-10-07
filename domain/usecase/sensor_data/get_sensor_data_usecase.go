package sensor_data

import (
	"context"
	"device-service/domain/repository"
	"time"
)

type GetSensorDataRequest struct {
	ID string
}

type GetSensorDataResponse struct {
	ID           string
	DeviceID     string
	SensorType   string
	Value        float64
	Unit         string
	RecordedAt   time.Time
	IsAlert      bool
	QualityScore float64
	CreatedAt    time.Time
}

type GetSensorDataUsecase struct {
	sensorDataRepo repository.SensorDataRepository
}

func NewGetSensorDataUsecase(sensorDataRepo repository.SensorDataRepository) *GetSensorDataUsecase {
	return &GetSensorDataUsecase{
		sensorDataRepo: sensorDataRepo,
	}
}

func (u *GetSensorDataUsecase) Execute(ctx context.Context, req *GetSensorDataRequest) (*GetSensorDataResponse, error) {
	sensorData, err := u.sensorDataRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if sensorData == nil {
		return nil, ErrSensorDataNotFound
	}

	return &GetSensorDataResponse{
		ID:           sensorData.ID,
		DeviceID:     sensorData.DeviceID,
		SensorType:   sensorData.SensorType,
		Value:        sensorData.Value,
		Unit:         sensorData.Unit,
		RecordedAt:   sensorData.RecordedAt,
		IsAlert:      sensorData.IsAlert,
		QualityScore: sensorData.QualityScore,
		CreatedAt:    sensorData.CreatedAt,
	}, nil
}
