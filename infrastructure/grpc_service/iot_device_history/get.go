package iot_device_history_service

import (
	"context"
	"device-service/domain/usecase/iot_device_history"

	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *IoTDeviceHistoryService) GetIoTDeviceHistory(ctx context.Context, req *proto_iot_device_history.GetIoTDeviceHistoryRequest) (*proto_iot_device_history.GetIoTDeviceHistoryResponse, error) {
	getRequest := s.convertRequestGetIoTDeviceHistory(req)
	getResponse, err := s.iotDeviceHistoryUsecase.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseGetIoTDeviceHistory(getResponse), nil
}

func (s *IoTDeviceHistoryService) convertRequestGetIoTDeviceHistory(req *proto_iot_device_history.GetIoTDeviceHistoryRequest) *iot_device_history.GetIoTDeviceHistoryRequest {
	return &iot_device_history.GetIoTDeviceHistoryRequest{
		ID: req.Id,
	}
}

func (s *IoTDeviceHistoryService) convertResponseGetIoTDeviceHistory(response *iot_device_history.GetIoTDeviceHistoryResponse) *proto_iot_device_history.GetIoTDeviceHistoryResponse {
	history := &proto_iot_device_history.IoTDeviceHistory{
		Id:          response.ID,
		DeviceId:    response.DeviceID,
		Action:      response.Action,
		ActionDate:  response.ActionDate,
		PerformedBy: response.PerformedBy,
	}

	if response.OldValue != nil {
		if oldValueStruct, err := structpb.NewStruct(response.OldValue); err == nil {
			history.OldValue = oldValueStruct
		}
	}
	if response.NewValue != nil {
		if newValueStruct, err := structpb.NewStruct(response.NewValue); err == nil {
			history.NewValue = newValueStruct
		}
	}
	if response.Notes != nil {
		history.Notes = *response.Notes
	}

	return &proto_iot_device_history.GetIoTDeviceHistoryResponse{
		History: history,
	}
}
