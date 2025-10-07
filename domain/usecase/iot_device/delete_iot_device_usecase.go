package iot_device

import (
	"context"
	"device-service/domain/repository"
)

type DeleteIoTDeviceRequest struct {
	ID string
}

type DeleteIoTDeviceResponse struct {
	Success bool
	Message string
}

type DeleteIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
}

func NewDeleteIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository) *DeleteIoTDeviceUsecase {
	return &DeleteIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
	}
}

func (u *DeleteIoTDeviceUsecase) Execute(ctx context.Context, req *DeleteIoTDeviceRequest) (*DeleteIoTDeviceResponse, error) {
	_, err := u.iotDeviceRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	err = u.iotDeviceRepo.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteIoTDeviceResponse{
		Success: true,
		Message: "Xóa thiết bị IoT thành công",
	}, nil
}
