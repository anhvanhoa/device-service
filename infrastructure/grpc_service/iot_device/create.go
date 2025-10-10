package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IoTDeviceService) CreateIoTDevice(ctx context.Context, req *proto_iot_device.CreateIoTDeviceRequest) (*proto_iot_device.CreateIoTDeviceResponse, error) {
	createRequest := s.convertRequestCreateIoTDevice(req)
	createResponse, err := s.iotDeviceUsecase.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseCreateIoTDevice(createResponse), nil
}

func (s *IoTDeviceService) convertRequestCreateIoTDevice(req *proto_iot_device.CreateIoTDeviceRequest) *iot_device.CreateIoTDeviceRequest {
	request := &iot_device.CreateIoTDeviceRequest{
		DeviceName:         req.DeviceName,
		DeviceTypeID:       req.DeviceTypeId,
		Model:              req.Model,
		MacAddress:         req.MacAddress,
		IPAddress:          req.IpAddress,
		GreenhouseID:       req.GreenhouseId,
		GrowingZoneID:      req.GrowingZoneId,
		BatteryLevel:       int(req.BatteryLevel),
		CreatedBy:          req.CreatedBy,
		ReadInterval:       int(req.ReadInterval),
		AlertEnabled:       req.AlertEnabled,
		AlertThresholdHigh: req.AlertThresholdHigh,
		AlertThresholdLow:  req.AlertThresholdLow,
	}

	if req.InstallationDate != nil {
		installationDate := req.InstallationDate.AsTime()
		request.InstallationDate = &installationDate
	}

	if req.DefaultConfig != nil {
		request.DefaultConfig = req.DefaultConfig.AsMap()
	}

	return request
}

func (s *IoTDeviceService) convertResponseCreateIoTDevice(response *iot_device.CreateIoTDeviceResponse) *proto_iot_device.CreateIoTDeviceResponse {
	resp := &proto_iot_device.CreateIoTDeviceResponse{
		DeviceName:    response.DeviceName,
		DeviceTypeId:  response.DeviceTypeID,
		Model:         response.Model,
		MacAddress:    response.MacAddress,
		IpAddress:     response.IPAddress,
		GreenhouseId:  response.GreenhouseID,
		GrowingZoneId: response.GrowingZoneID,
		BatteryLevel:  int32(response.BatteryLevel),
		Status:        response.Status,
		CreatedBy:     response.CreatedBy,
		CreatedAt:     timestamppb.New(response.CreatedAt),
	}

	if response.DefaultConfig != nil {
		if defaultConfigStruct, err := structpb.NewStruct(response.DefaultConfig); err == nil {
			resp.DefaultConfig = defaultConfigStruct
		}
	}

	if response.InstallationDate != nil {
		resp.InstallationDate = timestamppb.New(*response.InstallationDate)
	}

	return resp
}
