package device_type_service

import (
	"context"
	"device-service/domain/usecase/device_type"

	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *DeviceTypeService) GetDeviceType(ctx context.Context, req *proto_device_type.GetDeviceTypeRequest) (*proto_device_type.GetDeviceTypeResponse, error) {
	getRequest := s.convertRequestGetDeviceType(req)
	getResponse, err := s.deviceTypeUsecase.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseGetDeviceType(getResponse), nil
}

func (s *DeviceTypeService) convertRequestGetDeviceType(req *proto_device_type.GetDeviceTypeRequest) *device_type.GetDeviceTypeRequest {
	return &device_type.GetDeviceTypeRequest{
		ID: req.Id,
	}
}

func (s *DeviceTypeService) convertResponseGetDeviceType(response *device_type.GetDeviceTypeResponse) *proto_device_type.GetDeviceTypeResponse {
	deviceType := &proto_device_type.DeviceType{
		Id:          response.ID,
		TypeCode:    response.TypeCode,
		Description: response.Description,
		CreatedAt:   timestamppb.New(response.CreatedAt),
	}

	if response.UpdatedAt != nil {
		deviceType.UpdatedAt = timestamppb.New(*response.UpdatedAt)
	}

	return &proto_device_type.GetDeviceTypeResponse{
		DeviceType: deviceType,
	}
}
