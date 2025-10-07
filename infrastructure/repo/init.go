package repo

import (
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type Repositories interface {
	DeviceType() repository.DeviceTypeRepository
	IoTDevice() repository.IoTDeviceRepository
	IoTDeviceHistory() repository.IoTDeviceHistoryRepository
	SensorData() repository.SensorDataRepository
}

type RepositoriesImpl struct {
	deviceType       repository.DeviceTypeRepository
	iotDevice        repository.IoTDeviceRepository
	iotDeviceHistory repository.IoTDeviceHistoryRepository
	sensorData       repository.SensorDataRepository
}

func InitRepositories(db *pg.DB, helper utils.Helper) Repositories {
	return &RepositoriesImpl{
		deviceType:       NewDeviceTypeRepository(db, helper),
		iotDevice:        NewIoTDeviceRepository(db, helper),
		iotDeviceHistory: NewIoTDeviceHistoryRepository(db, helper),
		sensorData:       NewSensorDataRepository(db, helper),
	}
}

func (r *RepositoriesImpl) DeviceType() repository.DeviceTypeRepository {
	return r.deviceType
}

func (r *RepositoriesImpl) IoTDevice() repository.IoTDeviceRepository {
	return r.iotDevice
}

func (r *RepositoriesImpl) IoTDeviceHistory() repository.IoTDeviceHistoryRepository {
	return r.iotDeviceHistory
}

func (r *RepositoriesImpl) SensorData() repository.SensorDataRepository {
	return r.sensorData
}
