package device_type

import (
	"context"
	"device-service/domain/repository"
	"time"
)

type GetDeviceTypeRequest struct {
	ID string
}

type GetDeviceTypeResponse struct {
	ID          string
	TypeCode    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type GetDeviceTypeUsecase struct {
	deviceTypeRepo repository.DeviceTypeRepository
}

func NewGetDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository) *GetDeviceTypeUsecase {
	return &GetDeviceTypeUsecase{
		deviceTypeRepo: deviceTypeRepo,
	}
}

func (u *GetDeviceTypeUsecase) Execute(ctx context.Context, req *GetDeviceTypeRequest) (*GetDeviceTypeResponse, error) {
	deviceType, err := u.deviceTypeRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if deviceType == nil {
		return nil, ErrDeviceTypeNotFound
	}

	response := &GetDeviceTypeResponse{
		ID:          deviceType.ID,
		TypeCode:    deviceType.TypeCode,
		Description: deviceType.Description,
		CreatedAt:   deviceType.CreatedAt,
	}

	if deviceType.UpdatedAt != nil {
		response.UpdatedAt = deviceType.UpdatedAt
	}

	return response, nil
}
