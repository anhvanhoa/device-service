package iot_device_history_service

import (
	"context"
	"device-service/domain/usecase/iot_device_history"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *IoTDeviceHistoryService) ListIoTDeviceHistory(ctx context.Context, req *proto_iot_device_history.ListIoTDeviceHistoryRequest) (*proto_iot_device_history.ListIoTDeviceHistoryResponse, error) {
	listRequest := s.convertRequestListIoTDeviceHistory(req)
	listResponse, err := s.iotDeviceHistoryUsecase.List(ctx, listRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseListIoTDeviceHistory(listResponse), nil
}

func (s *IoTDeviceHistoryService) convertRequestListIoTDeviceHistory(req *proto_iot_device_history.ListIoTDeviceHistoryRequest) *iot_device_history.ListIoTDeviceHistoryRequest {
	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	return &iot_device_history.ListIoTDeviceHistoryRequest{
		Pagination: pagination,
	}
}

func (s *IoTDeviceHistoryService) convertResponseListIoTDeviceHistory(response *iot_device_history.ListIoTDeviceHistoryResponse) *proto_iot_device_history.ListIoTDeviceHistoryResponse {
	histories := make([]*proto_iot_device_history.IoTDeviceHistory, len(response.Data))
	for i, item := range response.Data {
		history := &proto_iot_device_history.IoTDeviceHistory{
			Id:          item.ID,
			DeviceId:    item.DeviceID,
			Action:      item.Action,
			ActionDate:  item.ActionDate,
			PerformedBy: item.PerformedBy,
		}

		if item.OldValue != nil {
			if oldValueStruct, err := structpb.NewStruct(item.OldValue); err == nil {
				history.OldValue = oldValueStruct
			}
		}
		if item.NewValue != nil {
			if newValueStruct, err := structpb.NewStruct(item.NewValue); err == nil {
				history.NewValue = newValueStruct
			}
		}
		if item.Notes != nil {
			history.Notes = *item.Notes
		}

		histories[i] = history
	}

	pagination := &common_proto.PaginationResponse{
		Total:      int32(response.Total),
		Page:       int32(response.Page),
		PageSize:   int32(response.PageSize),
		TotalPages: int32(response.TotalPages),
	}

	return &proto_iot_device_history.ListIoTDeviceHistoryResponse{
		Data:       histories,
		Pagination: pagination,
	}
}
