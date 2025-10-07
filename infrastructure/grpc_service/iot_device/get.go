package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IoTDeviceService) GetIoTDevice(ctx context.Context, req *proto_iot_device.GetIoTDeviceRequest) (*proto_iot_device.GetIoTDeviceResponse, error) {
	getRequest := s.convertRequestGetIoTDevice(req)
	getResponse, err := s.iotDeviceUsecase.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseGetIoTDevice(getResponse), nil
}

func (s *IoTDeviceService) convertRequestGetIoTDevice(req *proto_iot_device.GetIoTDeviceRequest) *iot_device.GetIoTDeviceRequest {
	return &iot_device.GetIoTDeviceRequest{
		ID: req.Id,
	}
}

func (s *IoTDeviceService) convertResponseGetIoTDevice(response *iot_device.GetIoTDeviceResponse) *proto_iot_device.GetIoTDeviceResponse {
	device := &proto_iot_device.IoTDevice{
		Id:            response.ID,
		DeviceName:    response.DeviceName,
		Status:        response.Status,
		DeviceTypeId:  response.DeviceTypeID,
		Model:         response.Model,
		MacAddress:    response.MacAddress,
		IpAddress:     response.IPAddress,
		GreenhouseId:  response.GreenhouseID,
		GrowingZoneId: response.GrowingZoneID,
		BatteryLevel:  int32(response.BatteryLevel),
		CreatedBy:     response.CreatedBy,
		CreatedAt:     timestamppb.New(response.CreatedAt),
	}

	if response.InstallationDate != nil {
		device.InstallationDate = timestamppb.New(*response.InstallationDate)
	}
	if response.LastMaintenanceDate != nil {
		device.LastMaintenanceDate = timestamppb.New(*response.LastMaintenanceDate)
	}
	if response.UpdatedAt != nil {
		device.UpdatedAt = timestamppb.New(*response.UpdatedAt)
	}

	if response.Configuration != nil {
		if configStruct, err := structpb.NewStruct(response.Configuration); err == nil {
			device.Configuration = configStruct
		}
	}
	if response.DefaultConfig != nil {
		if defaultConfigStruct, err := structpb.NewStruct(response.DefaultConfig); err == nil {
			device.DefaultConfig = defaultConfigStruct
		}
	}

	return &proto_iot_device.GetIoTDeviceResponse{
		Device: device,
	}
}
