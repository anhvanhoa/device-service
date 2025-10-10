package iot_device

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type CreateIoTDeviceRequest struct {
	ID                 string
	DeviceName         string
	DeviceTypeID       string
	Model              string
	MacAddress         string
	IPAddress          string
	GreenhouseID       string
	GrowingZoneID      string
	InstallationDate   *time.Time
	BatteryLevel       int
	DefaultConfig      map[string]any
	CreatedBy          string
	ReadInterval       int
	AlertEnabled       bool
	AlertThresholdHigh float64
	AlertThresholdLow  float64
}

type CreateIoTDeviceResponse struct {
	ID               string
	DeviceName       string
	DeviceTypeID     string
	Model            string
	MacAddress       string
	IPAddress        string
	GreenhouseID     string
	GrowingZoneID    string
	InstallationDate *time.Time
	BatteryLevel     int
	Status           string
	DefaultConfig    map[string]any
	CreatedBy        string
	CreatedAt        time.Time
}

type CreateIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
}

func NewCreateIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository) *CreateIoTDeviceUsecase {
	return &CreateIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
	}
}

func (u *CreateIoTDeviceUsecase) Execute(ctx context.Context, req *CreateIoTDeviceRequest) (*CreateIoTDeviceResponse, error) {
	if req.MacAddress != "" {
		exists, err := u.iotDeviceRepo.ExistsByMacAddress(ctx, req.MacAddress)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrMacAddressAlreadyExists
		}
	}

	if req.InstallationDate == nil {
		now := time.Now()
		req.InstallationDate = &now
	}

	if req.DefaultConfig == nil {
		req.DefaultConfig = map[string]any{
			"read_interval":        req.ReadInterval,
			"alert_enabled":        req.AlertEnabled,
			"alert_threshold_high": req.AlertThresholdHigh,
			"alert_threshold_low":  req.AlertThresholdLow,
		}
	}

	device := &entity.IoTDevice{
		ID:                 req.ID,
		DeviceName:         req.DeviceName,
		DeviceTypeID:       req.DeviceTypeID,
		Model:              req.Model,
		MacAddress:         req.MacAddress,
		IPAddress:          req.IPAddress,
		GreenhouseID:       req.GreenhouseID,
		GrowingZoneID:      req.GrowingZoneID,
		InstallationDate:   req.InstallationDate,
		BatteryLevel:       req.BatteryLevel,
		Status:             entity.DeviceStatusActive,
		DefaultConfig:      entity.JSONB(req.DefaultConfig),
		ReadInterval:       req.ReadInterval,
		AlertEnabled:       req.AlertEnabled,
		AlertThresholdHigh: req.AlertThresholdHigh,
		AlertThresholdLow:  req.AlertThresholdLow,
		CreatedBy:          req.CreatedBy,
		CreatedAt:          time.Now(),
	}

	err := u.iotDeviceRepo.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	response := &CreateIoTDeviceResponse{
		ID:            device.ID,
		DeviceName:    device.DeviceName,
		DeviceTypeID:  device.DeviceTypeID,
		Model:         device.Model,
		MacAddress:    device.MacAddress,
		IPAddress:     device.IPAddress,
		GreenhouseID:  device.GreenhouseID,
		GrowingZoneID: device.GrowingZoneID,
		BatteryLevel:  device.BatteryLevel,
		Status:        string(device.Status),
		DefaultConfig: map[string]any(device.DefaultConfig),
		CreatedBy:     device.CreatedBy,
		CreatedAt:     device.CreatedAt,
	}

	if device.InstallationDate != nil {
		response.InstallationDate = device.InstallationDate
	}

	return response, nil
}
