package iot_device_service

import (
	"device-service/domain/repository"
	"device-service/domain/usecase/iot_device"

	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
)

type IoTDeviceService struct {
	proto_iot_device.UnimplementedIoTDeviceServiceServer
	iotDeviceUsecase iot_device.IoTDeviceUsecase
}

func NewIoTDeviceService(iotDeviceRepo repository.IoTDeviceRepository, helper utils.Helper, mq mq.MQI) proto_iot_device.IoTDeviceServiceServer {
	iotDeviceUsecase := iot_device.NewIoTDeviceUsecase(iotDeviceRepo, helper, mq)
	return &IoTDeviceService{iotDeviceUsecase: iotDeviceUsecase}
}
