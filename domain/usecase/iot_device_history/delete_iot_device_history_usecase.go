package iot_device_history

import (
	"context"
	"device-service/domain/repository"
)

type DeleteIoTDeviceHistoryRequest struct {
	ID string
}

type DeleteIoTDeviceHistoryResponse struct {
	Success bool
	Message string
}

type DeleteIoTDeviceHistoryUsecase struct {
	deviceHistoryRepo repository.IoTDeviceHistoryRepository
}

func NewDeleteIoTDeviceHistoryUsecase(deviceHistoryRepo repository.IoTDeviceHistoryRepository) *DeleteIoTDeviceHistoryUsecase {
	return &DeleteIoTDeviceHistoryUsecase{
		deviceHistoryRepo: deviceHistoryRepo,
	}
}

func (u *DeleteIoTDeviceHistoryUsecase) Execute(ctx context.Context, req *DeleteIoTDeviceHistoryRequest) (*DeleteIoTDeviceHistoryResponse, error) {
	_, err := u.deviceHistoryRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	err = u.deviceHistoryRepo.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteIoTDeviceHistoryResponse{
		Success: true,
		Message: "Xóa lịch sử thiết bị thành công",
	}, nil
}
