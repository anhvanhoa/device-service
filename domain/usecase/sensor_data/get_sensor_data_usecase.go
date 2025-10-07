package sensor_data

import (
	"context"
	"device-service/domain/repository"
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
	RecordedAt   string
	IsAlert      bool
	QualityScore *float64
	CreatedAt    string
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
		SensorType:   *sensorData.SensorType,
		Value:        *sensorData.Value,
		Unit:         *sensorData.Unit,
		RecordedAt:   sensorData.RecordedAt.Format("2006-01-02T15:04:05Z07:00"),
		IsAlert:      sensorData.IsAlert,
		QualityScore: sensorData.QualityScore,
		CreatedAt:    sensorData.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
