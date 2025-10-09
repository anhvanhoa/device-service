package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
)

func (s *IoTDeviceService) ControlIoTDevice(ctx context.Context, req *proto_iot_device.ControlIoTDeviceRequest) (*proto_iot_device.ControlIoTDeviceResponse, error) {
	controlRequest := s.convertRequestControlIoTDevice(req)
	controlResponse, err := s.iotDeviceUsecase.Control(ctx, controlRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseControlIoTDevice(controlResponse), nil
}

func (s *IoTDeviceService) convertRequestControlIoTDevice(req *proto_iot_device.ControlIoTDeviceRequest) *iot_device.ControlIoTDeviceRequest {
	return &iot_device.ControlIoTDeviceRequest{
		DeviceID: req.Id,
		Action:   req.Action,
	}
}

func (s *IoTDeviceService) convertResponseControlIoTDevice(response *iot_device.ControlIoTDeviceResponse) *proto_iot_device.ControlIoTDeviceResponse {
	return &proto_iot_device.ControlIoTDeviceResponse{
		Id:      response.DeviceID,
		Status:  string(response.Status),
		Action:  response.Action,
		Message: response.Message,
	}
}
