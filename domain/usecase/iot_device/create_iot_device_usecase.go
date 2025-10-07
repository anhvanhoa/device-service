package iot_device

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type CreateIoTDeviceRequest struct {
	DeviceName       string
	DeviceTypeID     string
	Model            string
	MacAddress       string
	IPAddress        string
	GreenhouseID     string
	GrowingZoneID    string
	InstallationDate *time.Time
	BatteryLevel     int
	Configuration    map[string]any
	DefaultConfig    map[string]any
	CreatedBy        string
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
	Configuration    map[string]any
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

	var installationDate *time.Time
	if req.InstallationDate != nil {
		installationDate = req.InstallationDate
	}

	device := &entity.IoTDevice{
		DeviceName:       req.DeviceName,
		DeviceTypeID:     req.DeviceTypeID,
		Model:            req.Model,
		MacAddress:       req.MacAddress,
		IPAddress:        req.IPAddress,
		GreenhouseID:     req.GreenhouseID,
		GrowingZoneID:    req.GrowingZoneID,
		InstallationDate: installationDate,
		BatteryLevel:     req.BatteryLevel,
		Status:           entity.DeviceStatusActive,
		Configuration:    entity.JSONB(req.Configuration),
		DefaultConfig:    entity.JSONB(req.DefaultConfig),
		CreatedBy:        req.CreatedBy,
		CreatedAt:        time.Now(),
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
		Configuration: map[string]any(device.Configuration),
		DefaultConfig: map[string]any(device.DefaultConfig),
		CreatedBy:     device.CreatedBy,
		CreatedAt:     device.CreatedAt,
	}

	if device.InstallationDate != nil {
		response.InstallationDate = device.InstallationDate
	}

	return response, nil
}
