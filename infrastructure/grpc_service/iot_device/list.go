package iot_device_service

import (
	"context"
	"device-service/domain/usecase/iot_device"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IoTDeviceService) ListIoTDevice(ctx context.Context, req *proto_iot_device.ListIoTDeviceRequest) (*proto_iot_device.ListIoTDeviceResponse, error) {
	listRequest := s.convertRequestListIoTDevice(req)
	listResponse, err := s.iotDeviceUsecase.List(ctx, listRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseListIoTDevice(listResponse), nil
}

func (s *IoTDeviceService) convertRequestListIoTDevice(req *proto_iot_device.ListIoTDeviceRequest) *iot_device.ListIoTDeviceRequest {
	if req.Pagination == nil {
		req.Pagination = &common_proto.PaginationRequest{
			Page:     1,
			PageSize: 10,
		}
	}

	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	filters := iot_device.IoTDeviceFilters{}
	if req.Filters != nil {
		filters = iot_device.IoTDeviceFilters{
			DeviceTypeID:  req.Filters.DeviceTypeId,
			Status:        req.Filters.Status,
			GreenhouseID:  req.Filters.GreenhouseId,
			GrowingZoneID: req.Filters.GrowingZoneId,
		}
	}

	return &iot_device.ListIoTDeviceRequest{
		Pagination: pagination,
		Filters:    filters,
	}
}

func (s *IoTDeviceService) convertResponseListIoTDevice(response *iot_device.ListIoTDeviceResponse) *proto_iot_device.ListIoTDeviceResponse {
	devices := make([]*proto_iot_device.IoTDevice, len(response.Data))
	for i, item := range response.Data {
		device := &proto_iot_device.IoTDevice{
			Id:            item.ID,
			DeviceName:    item.DeviceName,
			Status:        item.Status,
			Model:         item.Model,
			MacAddress:    item.MacAddress,
			IpAddress:     item.IPAddress,
			GreenhouseId:  item.GreenhouseID,
			GrowingZoneId: item.GrowingZoneID,
			BatteryLevel:  int32(item.BatteryLevel),
			CreatedBy:     item.CreatedBy,
			CreatedAt:     timestamppb.New(item.CreatedAt),
		}

		if item.InstallationDate != nil {
			device.InstallationDate = timestamppb.New(*item.InstallationDate)
		}

		devices[i] = device
	}

	pagination := &common_proto.PaginationResponse{
		Total:      int32(response.Total),
		Page:       int32(response.Page),
		PageSize:   int32(response.PageSize),
		TotalPages: int32(response.TotalPages),
	}

	return &proto_iot_device.ListIoTDeviceResponse{
		Data:       devices,
		Pagination: pagination,
	}
}
