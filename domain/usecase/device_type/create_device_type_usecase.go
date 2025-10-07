package device_type

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type CreateDeviceTypeRequest struct {
	TypeCode    string
	Description string
}

type CreateDeviceTypeResponse struct {
	ID          string
	TypeCode    string
	Description string
	CreatedAt   time.Time
}

type CreateDeviceTypeUsecase struct {
	deviceTypeRepo repository.DeviceTypeRepository
}

func NewCreateDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository) *CreateDeviceTypeUsecase {
	return &CreateDeviceTypeUsecase{
		deviceTypeRepo: deviceTypeRepo,
	}
}

func (u *CreateDeviceTypeUsecase) Execute(ctx context.Context, req *CreateDeviceTypeRequest) (*CreateDeviceTypeResponse, error) {
	exists, err := u.deviceTypeRepo.Exists(ctx, req.TypeCode)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDeviceTypeAlreadyExists
	}

	deviceType := &entity.DeviceType{
		TypeCode:    req.TypeCode,
		Description: req.Description,
	}

	err = u.deviceTypeRepo.Create(ctx, deviceType)
	if err != nil {
		return nil, err
	}

	return &CreateDeviceTypeResponse{
		ID:          deviceType.ID,
		TypeCode:    deviceType.TypeCode,
		Description: deviceType.Description,
		CreatedAt:   deviceType.CreatedAt,
	}, nil
}
