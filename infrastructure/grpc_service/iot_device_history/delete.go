package iot_device_history_service

import (
	"context"
	"device-service/domain/usecase/iot_device_history"

	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
)

func (s *IoTDeviceHistoryService) DeleteIoTDeviceHistory(ctx context.Context, req *proto_iot_device_history.DeleteIoTDeviceHistoryRequest) (*proto_iot_device_history.DeleteIoTDeviceHistoryResponse, error) {
	deleteRequest := s.convertRequestDeleteIoTDeviceHistory(req)
	deleteResponse, err := s.iotDeviceHistoryUsecase.Delete(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseDeleteIoTDeviceHistory(deleteResponse), nil
}

func (s *IoTDeviceHistoryService) convertRequestDeleteIoTDeviceHistory(req *proto_iot_device_history.DeleteIoTDeviceHistoryRequest) *iot_device_history.DeleteIoTDeviceHistoryRequest {
	return &iot_device_history.DeleteIoTDeviceHistoryRequest{
		ID: req.Id,
	}
}

func (s *IoTDeviceHistoryService) convertResponseDeleteIoTDeviceHistory(response *iot_device_history.DeleteIoTDeviceHistoryResponse) *proto_iot_device_history.DeleteIoTDeviceHistoryResponse {
	return &proto_iot_device_history.DeleteIoTDeviceHistoryResponse{
		Success: response.Success,
	}
}
