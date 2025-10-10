package sensor_data

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"encoding/json"
	"fmt"
	"time"
)

type ProcessSensorDataRequest struct {
	DeviceID     string  `json:"deviceId"`
	SensorType   string  `json:"sensorType"`
	Value        float64 `json:"value"`
	Unit         string  `json:"unit"`
	IsAlert      bool    `json:"isAlert"`
	QualityScore float64 `json:"qualityScore"`
	Timestamp    int64   `json:"timestamp"`
}

type ProcessSensorDataResponse struct {
	ID        string
	Processed bool
	AlertSent bool
}

type ProcessSensorDataUsecase struct {
	sensorDataRepo repository.SensorDataRepository
}

func NewProcessSensorDataUsecase(sensorDataRepo repository.SensorDataRepository) *ProcessSensorDataUsecase {
	return &ProcessSensorDataUsecase{
		sensorDataRepo: sensorDataRepo,
	}
}

func (u *ProcessSensorDataUsecase) Execute(ctx context.Context, req *ProcessSensorDataRequest) (*ProcessSensorDataResponse, error) {
	// Tạo entity SensorData
	sensorData := &entity.SensorData{
		DeviceID:     req.DeviceID,
		SensorType:   req.SensorType,
		Value:        req.Value,
		Unit:         req.Unit,
		IsAlert:      req.IsAlert,
		QualityScore: req.QualityScore,
		RecordedAt:   time.Unix(req.Timestamp/1000, 0), // Chuyển đổi từ milliseconds
		CreatedAt:    time.Now(),
	}

	// Lưu vào database
	err := u.sensorDataRepo.Create(ctx, sensorData)
	if err != nil {
		return nil, fmt.Errorf("không thể lưu dữ liệu cảm biến: %w", err)
	}

	response := &ProcessSensorDataResponse{
		ID:        sensorData.ID,
		Processed: true,
		AlertSent: false,
	}

	// Kiểm tra và xử lý alert nếu cần
	if req.IsAlert {
		// TODO: Gửi thông báo alert
		response.AlertSent = true
	}

	return response, nil
}

// ParseSensorDataFromJSON parses sensor data from JSON message
func ParseSensorDataFromJSON(jsonData []byte) (*ProcessSensorDataRequest, error) {
	var req ProcessSensorDataRequest
	err := json.Unmarshal(jsonData, &req)
	if err != nil {
		return nil, fmt.Errorf("không thể parse JSON: %w", err)
	}
	return &req, nil
}
