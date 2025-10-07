package device_type

import (
	"context"
	"device-service/domain/repository"
)

type DeleteDeviceTypeRequest struct {
	ID string
}

type DeleteDeviceTypeResponse struct {
	Success bool
	Message string
}

type DeleteDeviceTypeUsecase struct {
	deviceTypeRepo repository.DeviceTypeRepository
}

func NewDeleteDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository) *DeleteDeviceTypeUsecase {
	return &DeleteDeviceTypeUsecase{
		deviceTypeRepo: deviceTypeRepo,
	}
}

func (u *DeleteDeviceTypeUsecase) Execute(ctx context.Context, req *DeleteDeviceTypeRequest) (*DeleteDeviceTypeResponse, error) {
	_, err := u.deviceTypeRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	err = u.deviceTypeRepo.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteDeviceTypeResponse{
		Success: true,
		Message: "Xóa loại thiết bị thành công",
	}, nil
}
