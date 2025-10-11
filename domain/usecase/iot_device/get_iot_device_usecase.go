package iot_device

import (
	"context"
	"device-service/domain/repository"
	"time"
)

type GetIoTDeviceRequest struct {
	ID string
}

type GetIoTDeviceResponse struct {
	ID                  string
	DeviceName          string
	DeviceTypeID        string
	Model               string
	MacAddress          string
	IPAddress           string
	GreenhouseID        string
	GrowingZoneID       string
	InstallationDate    *time.Time
	LastMaintenanceDate *time.Time
	BatteryLevel        int
	Status              string
	DefaultConfig       map[string]any
	ReadInterval        int
	AlertEnabled        bool
	AlertThresholdHigh  float64
	AlertThresholdLow   float64
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

type GetIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
}

func NewGetIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository) *GetIoTDeviceUsecase {
	return &GetIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
	}
}

func (u *GetIoTDeviceUsecase) Execute(ctx context.Context, req *GetIoTDeviceRequest) (*GetIoTDeviceResponse, error) {
	device, err := u.iotDeviceRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrIoTDeviceNotFound
	}

	response := &GetIoTDeviceResponse{
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

	if device.LastMaintenanceDate != nil {
		response.LastMaintenanceDate = device.LastMaintenanceDate
	}

	if device.UpdatedAt != nil {
		response.UpdatedAt = device.UpdatedAt
	}

	return response, nil
}
