package iot_device

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type UpdateIoTDeviceRequest struct {
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
	Configuration       map[string]any
	DefaultConfig       map[string]any
}

type UpdateIoTDeviceResponse struct {
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
	Configuration       map[string]any
	DefaultConfig       map[string]any
	UpdatedAt           *time.Time
	CreatedBy           string
	CreatedAt           time.Time
}

type UpdateIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
}

func NewUpdateIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository) *UpdateIoTDeviceUsecase {
	return &UpdateIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
	}
}

func (u *UpdateIoTDeviceUsecase) Execute(ctx context.Context, req *UpdateIoTDeviceRequest) (*UpdateIoTDeviceResponse, error) {
	existingDevice, err := u.iotDeviceRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingDevice == nil {
		return nil, ErrIoTDeviceNotFound
	}

	if req.MacAddress != "" && existingDevice.MacAddress != "" && req.MacAddress != existingDevice.MacAddress {
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

	var lastMaintenanceDate *time.Time
	if req.LastMaintenanceDate != nil {
		lastMaintenanceDate = req.LastMaintenanceDate
	}

	status := entity.DeviceStatus(req.Status)
	now := time.Now()
	device := &entity.IoTDevice{
		ID:                  req.ID,
		DeviceName:          req.DeviceName,
		DeviceTypeID:        req.DeviceTypeID,
		Model:               req.Model,
		MacAddress:          req.MacAddress,
		IPAddress:           req.IPAddress,
		GreenhouseID:        req.GreenhouseID,
		GrowingZoneID:       req.GrowingZoneID,
		InstallationDate:    installationDate,
		LastMaintenanceDate: lastMaintenanceDate,
		BatteryLevel:        req.BatteryLevel,
		Status:              status,
		Configuration:       entity.JSONB(req.Configuration),
		DefaultConfig:       entity.JSONB(req.DefaultConfig),
		CreatedBy:           existingDevice.CreatedBy,
		CreatedAt:           existingDevice.CreatedAt,
		UpdatedAt:           &now,
	}

	err = u.iotDeviceRepo.Update(ctx, device)
	if err != nil {
		return nil, err
	}

	response := &UpdateIoTDeviceResponse{
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
