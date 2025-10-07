package device_type_service

import (
	"context"
	"device-service/domain/usecase/device_type"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *DeviceTypeService) ListDeviceType(ctx context.Context, req *proto_device_type.ListDeviceTypeRequest) (*proto_device_type.ListDeviceTypeResponse, error) {
	listRequest := s.convertRequestListDeviceType(req)
	listResponse, err := s.deviceTypeUsecase.List(ctx, listRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseListDeviceType(listResponse), nil
}

func (s *DeviceTypeService) convertRequestListDeviceType(req *proto_device_type.ListDeviceTypeRequest) *device_type.ListDeviceTypeRequest {
	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	return &device_type.ListDeviceTypeRequest{
		Pagination: pagination,
	}
}

func (s *DeviceTypeService) convertResponseListDeviceType(response *device_type.ListDeviceTypeResponse) *proto_device_type.ListDeviceTypeResponse {
	deviceTypes := make([]*proto_device_type.DeviceType, len(response.Data))
	for i, item := range response.Data {
		deviceType := &proto_device_type.DeviceType{
			Id:          item.ID,
			TypeCode:    item.TypeCode,
			Description: item.Description,
			CreatedAt:   timestamppb.New(item.CreatedAt),
		}

		if item.UpdatedAt != nil {
			deviceType.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}

		deviceTypes[i] = deviceType
	}

	pagination := &common_proto.PaginationResponse{
		Total:      int32(response.Total),
		Page:       int32(response.Page),
		PageSize:   int32(response.PageSize),
		TotalPages: int32(response.TotalPages),
	}

	return &proto_device_type.ListDeviceTypeResponse{
		Data:       deviceTypes,
		Pagination: pagination,
	}
}
