package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
)

func (s *IoTDeviceService) DeleteIoTDevice(ctx context.Context, req *proto_iot_device.DeleteIoTDeviceRequest) (*proto_iot_device.DeleteIoTDeviceResponse, error) {
	deleteRequest := s.convertRequestDeleteIoTDevice(req)
	deleteResponse, err := s.iotDeviceUsecase.Delete(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseDeleteIoTDevice(deleteResponse), nil
}

func (s *IoTDeviceService) convertRequestDeleteIoTDevice(req *proto_iot_device.DeleteIoTDeviceRequest) *iot_device.DeleteIoTDeviceRequest {
	return &iot_device.DeleteIoTDeviceRequest{
		ID: req.Id,
	}
}

func (s *IoTDeviceService) convertResponseDeleteIoTDevice(response *iot_device.DeleteIoTDeviceResponse) *proto_iot_device.DeleteIoTDeviceResponse {
	return &proto_iot_device.DeleteIoTDeviceResponse{
		Success: response.Success,
	}
}
