package device_type_service

import (
	"context"
	"device-service/domain/usecase/device_type"

	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *DeviceTypeService) CreateDeviceType(ctx context.Context, req *proto_device_type.CreateDeviceTypeRequest) (*proto_device_type.CreateDeviceTypeResponse, error) {
	createRequest := s.convertRequestCreateDeviceType(req)
	createResponse, err := s.deviceTypeUsecase.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseCreateDeviceType(createResponse), nil
}

func (s *DeviceTypeService) convertRequestCreateDeviceType(req *proto_device_type.CreateDeviceTypeRequest) *device_type.CreateDeviceTypeRequest {
	return &device_type.CreateDeviceTypeRequest{
		TypeCode:    req.TypeCode,
		Description: req.Description,
	}
}

func (s *DeviceTypeService) convertResponseCreateDeviceType(response *device_type.CreateDeviceTypeResponse) *proto_device_type.CreateDeviceTypeResponse {
	dt := &proto_device_type.CreateDeviceTypeResponse{
		Id:          response.ID,
		TypeCode:    response.TypeCode,
		Description: response.Description,
		CreatedAt:   timestamppb.New(response.CreatedAt),
	}
	return dt
}
