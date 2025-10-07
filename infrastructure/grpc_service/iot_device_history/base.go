package iot_device_history_service

import (
	"device-service/domain/repository"
	"device-service/domain/usecase/iot_device_history"

	"github.com/anhvanhoa/service-core/utils"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
)

type IoTDeviceHistoryService struct {
	proto_iot_device_history.UnimplementedIoTDeviceHistoryServiceServer
	iotDeviceHistoryUsecase iot_device_history.IoTDeviceHistoryUsecase
}

func NewIoTDeviceHistoryService(iotDeviceHistoryRepo repository.IoTDeviceHistoryRepository, helper utils.Helper) proto_iot_device_history.IoTDeviceHistoryServiceServer {
	iotDeviceHistoryUsecase := iot_device_history.NewIoTDeviceHistoryUsecase(iotDeviceHistoryRepo, helper)
	return &IoTDeviceHistoryService{iotDeviceHistoryUsecase: iotDeviceHistoryUsecase}
}
