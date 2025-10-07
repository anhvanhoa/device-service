package device_type_service

import (
	"device-service/domain/repository"
	"device-service/domain/usecase/device_type"

	"github.com/anhvanhoa/service-core/utils"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
)

type DeviceTypeService struct {
	proto_device_type.UnimplementedDeviceTypeServiceServer
	deviceTypeUsecase device_type.DeviceTypeUsecase
}

func NewDeviceTypeService(deviceTypeRepo repository.DeviceTypeRepository, helper utils.Helper) proto_device_type.DeviceTypeServiceServer {
	deviceTypeUsecase := device_type.NewDeviceTypeUsecase(deviceTypeRepo, helper)
	return &DeviceTypeService{deviceTypeUsecase: deviceTypeUsecase}
}
