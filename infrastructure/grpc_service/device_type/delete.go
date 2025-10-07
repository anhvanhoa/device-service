package device_type_service

import (
	"context"
	"device-service/domain/usecase/device_type"

	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
)

func (s *DeviceTypeService) DeleteDeviceType(ctx context.Context, req *proto_device_type.DeleteDeviceTypeRequest) (*proto_device_type.DeleteDeviceTypeResponse, error) {
	deleteRequest := s.convertRequestDeleteDeviceType(req)
	deleteResponse, err := s.deviceTypeUsecase.Delete(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseDeleteDeviceType(deleteResponse), nil
}

func (s *DeviceTypeService) convertRequestDeleteDeviceType(req *proto_device_type.DeleteDeviceTypeRequest) *device_type.DeleteDeviceTypeRequest {
	return &device_type.DeleteDeviceTypeRequest{
		ID: req.Id,
	}
}

func (s *DeviceTypeService) convertResponseDeleteDeviceType(response *device_type.DeleteDeviceTypeResponse) *proto_device_type.DeleteDeviceTypeResponse {
	return &proto_device_type.DeleteDeviceTypeResponse{
		Success: response.Success,
	}
}
