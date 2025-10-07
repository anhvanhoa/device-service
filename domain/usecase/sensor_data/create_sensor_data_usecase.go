package sensor_data

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type CreateSensorDataRequest struct {
	DeviceID     string
	SensorType   string
	Value        float64
	Unit         string
	IsAlert      bool
	QualityScore *float64
}

type CreateSensorDataResponse struct {
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

type CreateSensorDataUsecase struct {
	sensorDataRepo repository.SensorDataRepository
}

func NewCreateSensorDataUsecase(sensorDataRepo repository.SensorDataRepository) *CreateSensorDataUsecase {
	return &CreateSensorDataUsecase{
		sensorDataRepo: sensorDataRepo,
	}
}

func (u *CreateSensorDataUsecase) Execute(ctx context.Context, req *CreateSensorDataRequest) (*CreateSensorDataResponse, error) {
	sensorType := entity.SensorType(req.SensorType)
	if sensorType != entity.SensorTypeTemperature && sensorType != entity.SensorTypeHumidity &&
		sensorType != entity.SensorTypePH && sensorType != entity.SensorTypeLightIntensity &&
		sensorType != entity.SensorTypeSoilMoisture && sensorType != entity.SensorTypeCO2 &&
		sensorType != entity.SensorTypeWaterLevel {
		return nil, ErrInvalidSensorType
	}

	qualityScore := 1.0
	if req.QualityScore != nil {
		if *req.QualityScore < 0 || *req.QualityScore > 1 {
			return nil, ErrInvalidQualityScore
		}
		qualityScore = *req.QualityScore
	}

	sensorData := &entity.SensorData{
		DeviceID:     req.DeviceID,
		SensorType:   req.SensorType,
		Value:        req.Value,
		Unit:         req.Unit,
		IsAlert:      req.IsAlert,
		QualityScore: qualityScore,
	}

	err := u.sensorDataRepo.Create(ctx, sensorData)
	if err != nil {
		return nil, err
	}

	return &CreateSensorDataResponse{
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
