package iot_device_history_service

import (
	"context"
	"device-service/domain/usecase/iot_device_history"

	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *IoTDeviceHistoryService) CreateIoTDeviceHistory(ctx context.Context, req *proto_iot_device_history.CreateIoTDeviceHistoryRequest) (*proto_iot_device_history.CreateIoTDeviceHistoryResponse, error) {
	createRequest := s.convertRequestCreateIoTDeviceHistory(req)
	createResponse, err := s.iotDeviceHistoryUsecase.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseCreateIoTDeviceHistory(createResponse), nil
}

func (s *IoTDeviceHistoryService) convertRequestCreateIoTDeviceHistory(req *proto_iot_device_history.CreateIoTDeviceHistoryRequest) *iot_device_history.CreateIoTDeviceHistoryRequest {
	request := &iot_device_history.CreateIoTDeviceHistoryRequest{
		DeviceID:    req.DeviceId,
		Action:      req.Action,
		PerformedBy: req.PerformedBy,
	}

	if req.OldValue != nil {
		request.OldValue = req.OldValue.AsMap()
	}
	if req.NewValue != nil {
		request.NewValue = req.NewValue.AsMap()
	}
	if req.Notes != "" {
		request.Notes = &req.Notes
	}

	return request
}

func (s *IoTDeviceHistoryService) convertResponseCreateIoTDeviceHistory(response *iot_device_history.CreateIoTDeviceHistoryResponse) *proto_iot_device_history.CreateIoTDeviceHistoryResponse {
	resp := &proto_iot_device_history.CreateIoTDeviceHistoryResponse{
		DeviceId:    response.DeviceID,
		Action:      response.Action,
		ActionDate:  response.ActionDate,
		PerformedBy: response.PerformedBy,
	}

	if response.OldValue != nil {
		if oldValueStruct, err := structpb.NewStruct(response.OldValue); err == nil {
			resp.OldValue = oldValueStruct
		}
	}
	if response.NewValue != nil {
		if newValueStruct, err := structpb.NewStruct(response.NewValue); err == nil {
			resp.NewValue = newValueStruct
		}
	}
	if response.Notes != nil {
		resp.Notes = *response.Notes
	}

	return resp
}
