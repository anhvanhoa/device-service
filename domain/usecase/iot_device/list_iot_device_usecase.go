package iot_device

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListIoTDeviceRequest struct {
	Filters    IoTDeviceFilters
	Pagination *common.Pagination
}

type ListIoTDeviceResponse common.PaginationResult[IoTDeviceItem]

type IoTDeviceItem struct {
	ID                 string
	DeviceName         string
	DeviceTypeID       string
	Model              string
	MacAddress         string
	IPAddress          string
	GreenhouseID       string
	GrowingZoneID      string
	InstallationDate   *time.Time
	CreatedBy          string
	BatteryLevel       int
	Status             string
	DefaultConfig      map[string]any
	ReadInterval       int
	AlertEnabled       bool
	AlertThresholdHigh float64
	AlertThresholdLow  float64
	CreatedAt          time.Time
}

type IoTDeviceFilters struct {
	DeviceTypeID  *string
	Status        *string
	GreenhouseID  *string
	GrowingZoneID *string
	Search        *string
}

type ListIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
	helper        utils.Helper
}

func NewListIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository, helper utils.Helper) *ListIoTDeviceUsecase {
	return &ListIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
		helper:        helper,
	}
}

func (u *ListIoTDeviceUsecase) Execute(ctx context.Context, req *ListIoTDeviceRequest) (*ListIoTDeviceResponse, error) {
	pagination := req.Pagination
	if pagination == nil {
		pagination = &common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}

	filters := repository.IoTDeviceFilters{
		DeviceTypeID:  req.Filters.DeviceTypeID,
		GreenhouseID:  req.Filters.GreenhouseID,
		GrowingZoneID: req.Filters.GrowingZoneID,
		Search:        req.Filters.Search,
	}

	if req.Filters.Status != nil {
		status := entity.DeviceStatus(*req.Filters.Status)
		filters.Status = &status
	}

	devices, total, err := u.iotDeviceRepo.List(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}

	items := make([]IoTDeviceItem, len(devices))
	for i, device := range devices {
		item := IoTDeviceItem{
			ID:                 device.ID,
			DeviceName:         device.DeviceName,
			DeviceTypeID:       device.DeviceTypeID,
			Model:              device.Model,
			MacAddress:         device.MacAddress,
			IPAddress:          device.IPAddress,
			GreenhouseID:       device.GreenhouseID,
			GrowingZoneID:      device.GrowingZoneID,
			BatteryLevel:       device.BatteryLevel,
			Status:             string(device.Status),
			DefaultConfig:      map[string]any(device.DefaultConfig),
			ReadInterval:       device.ReadInterval,
			AlertEnabled:       device.AlertEnabled,
			AlertThresholdHigh: device.AlertThresholdHigh,
			AlertThresholdLow:  device.AlertThresholdLow,
			CreatedAt:          device.CreatedAt,
		}

		if device.InstallationDate != nil {
			item.InstallationDate = device.InstallationDate
		}

		items[i] = item
	}

	return &ListIoTDeviceResponse{
		Data:       items,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(int64(total), int64(pagination.PageSize)),
	}, nil
}
