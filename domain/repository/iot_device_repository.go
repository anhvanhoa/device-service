package repository

import (
	"context"
	"device-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type IoTDeviceRepository interface {
	Create(ctx context.Context, device *entity.IoTDevice) error
	GetByID(ctx context.Context, id string) (*entity.IoTDevice, error)
	GetByMacAddress(ctx context.Context, macAddress string) (*entity.IoTDevice, error)
	List(ctx context.Context, filters IoTDeviceFilters, pagination *common.Pagination) ([]*entity.IoTDevice, int64, error)
	Update(ctx context.Context, device *entity.IoTDevice) error
	UpdateStatus(ctx context.Context, id string, status entity.DeviceStatus) error
	Delete(ctx context.Context, id string) error
	GetLowBatteryDevices(ctx context.Context, threshold int, pagination *common.Pagination) ([]*entity.IoTDevice, int64, error)
	ExistsByMacAddress(ctx context.Context, macAddress string) (bool, error)
}

type IoTDeviceFilters struct {
	DeviceTypeID  string
	Status        entity.DeviceStatus
	GreenhouseID  string
	GrowingZoneID string
	Search        string
}
