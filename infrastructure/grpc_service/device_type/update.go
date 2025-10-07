package device_type_service

import (
	"context"
	"device-service/domain/usecase/device_type"

	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *DeviceTypeService) UpdateDeviceType(ctx context.Context, req *proto_device_type.UpdateDeviceTypeRequest) (*proto_device_type.UpdateDeviceTypeResponse, error) {
	updateRequest := s.convertRequestUpdateDeviceType(req)
	updateResponse, err := s.deviceTypeUsecase.Update(ctx, updateRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseUpdateDeviceType(updateResponse), nil
}

func (s *DeviceTypeService) convertRequestUpdateDeviceType(req *proto_device_type.UpdateDeviceTypeRequest) *device_type.UpdateDeviceTypeRequest {
	return &device_type.UpdateDeviceTypeRequest{
		ID:          req.Id,
		TypeCode:    req.TypeCode,
		Description: req.Description,
	}
}

func (s *DeviceTypeService) convertResponseUpdateDeviceType(response *device_type.UpdateDeviceTypeResponse) *proto_device_type.UpdateDeviceTypeResponse {
	deviceType := &proto_device_type.DeviceType{
		Id:          response.ID,
		TypeCode:    response.TypeCode,
		Description: response.Description,
		UpdatedAt:   timestamppb.New(response.UpdatedAt),
	}

	return &proto_device_type.UpdateDeviceTypeResponse{
		DeviceType: deviceType,
	}
}
