package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IoTDeviceService) UpdateIoTDevice(ctx context.Context, req *proto_iot_device.UpdateIoTDeviceRequest) (*proto_iot_device.UpdateIoTDeviceResponse, error) {
	updateRequest := s.convertRequestUpdateIoTDevice(req)
	updateResponse, err := s.iotDeviceUsecase.Update(ctx, updateRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseUpdateIoTDevice(updateResponse), nil
}

func (s *IoTDeviceService) convertRequestUpdateIoTDevice(req *proto_iot_device.UpdateIoTDeviceRequest) *iot_device.UpdateIoTDeviceRequest {
	request := &iot_device.UpdateIoTDeviceRequest{
		ID:            req.Id,
		DeviceName:    req.DeviceName,
		DeviceTypeID:  req.DeviceTypeId,
		Model:         req.Model,
		MacAddress:    req.MacAddress,
		IPAddress:     req.IpAddress,
		GreenhouseID:  req.GreenhouseId,
		GrowingZoneID: req.GrowingZoneId,
		BatteryLevel:  []int{int(req.BatteryLevel)}[0],
		Status:        req.Status,
	}

	if req.InstallationDate != nil {
		installationDate := req.InstallationDate.AsTime()
		request.InstallationDate = &installationDate
	}
	if req.LastMaintenanceDate != nil {
		lastMaintenanceDate := req.LastMaintenanceDate.AsTime()
		request.LastMaintenanceDate = &lastMaintenanceDate
	}

	if req.Configuration != nil {
		request.Configuration = req.Configuration.AsMap()
	}
	if req.DefaultConfig != nil {
		request.DefaultConfig = req.DefaultConfig.AsMap()
	}

	return request
}

func (s *IoTDeviceService) convertResponseUpdateIoTDevice(response *iot_device.UpdateIoTDeviceResponse) *proto_iot_device.UpdateIoTDeviceResponse {
	device := &proto_iot_device.IoTDevice{
		Id:            response.ID,
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
		UpdatedAt:     timestamppb.New(*response.UpdatedAt),
	}

	if response.InstallationDate != nil {
		device.InstallationDate = timestamppb.New(*response.InstallationDate)
	}
	if response.LastMaintenanceDate != nil {
		device.LastMaintenanceDate = timestamppb.New(*response.LastMaintenanceDate)
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

	return &proto_iot_device.UpdateIoTDeviceResponse{
		Device: device,
	}
}
