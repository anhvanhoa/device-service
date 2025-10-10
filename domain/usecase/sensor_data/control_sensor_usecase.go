package sensor_data

import (
	"context"
	"encoding/json"
	"fmt"
)

type ControlSensorRequest struct {
	DeviceID           string  `json:"deviceId"`
	SensorType         string  `json:"sensorType"`
	Enabled            bool    `json:"enabled"`
	ReadInterval       int     `json:"readInterval"`
	AlertThresholdHigh float64 `json:"alertThresholdHigh"`
	AlertThresholdLow  float64 `json:"alertThresholdLow"`
}

type ControlSensorResponse struct {
	Success bool
	Message string
}

type ControlSensorUsecase struct {
	// Có thể thêm repository nếu cần lưu cấu hình
}

func NewControlSensorUsecase() *ControlSensorUsecase {
	return &ControlSensorUsecase{}
}

func (u *ControlSensorUsecase) Execute(ctx context.Context, req *ControlSensorRequest) (*ControlSensorResponse, error) {
	// Validate request
	if req.DeviceID == "" {
		return &ControlSensorResponse{
			Success: false,
			Message: "DeviceID không được để trống",
		}, nil
	}

	if req.SensorType == "" {
		return &ControlSensorResponse{
			Success: false,
			Message: "SensorType không được để trống",
		}, nil
	}

	if req.ReadInterval <= 0 {
		return &ControlSensorResponse{
			Success: false,
			Message: "ReadInterval phải lớn hơn 0",
		}, nil
	}

	if req.AlertThresholdHigh <= req.AlertThresholdLow {
		return &ControlSensorResponse{
			Success: false,
			Message: "AlertThresholdHigh phải lớn hơn AlertThresholdLow",
		}, nil
	}

	// TODO: Implement logic để điều khiển cảm biến
	// - Gửi lệnh điều khiển đến thiết bị qua MQTT
	// - Lưu cấu hình vào database nếu cần
	// - Log hoạt động điều khiển

	return &ControlSensorResponse{
		Success: true,
		Message: fmt.Sprintf("Đã điều khiển cảm biến %s cho thiết bị %s", req.SensorType, req.DeviceID),
	}, nil
}

// ParseControlSensorFromJSON parses control sensor data from JSON message
func ParseControlSensorFromJSON(jsonData []byte) (*ControlSensorRequest, error) {
	var req ControlSensorRequest
	err := json.Unmarshal(jsonData, &req)
	if err != nil {
		return nil, fmt.Errorf("không thể parse JSON: %w", err)
	}
	return &req, nil
}
