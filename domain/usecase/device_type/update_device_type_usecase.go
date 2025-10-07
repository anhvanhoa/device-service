package device_type

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"
)

type UpdateDeviceTypeRequest struct {
	ID          string
	TypeCode    string
	Description string
}
type UpdateDeviceTypeResponse struct {
	ID          string
	TypeCode    string
	Description string
	UpdatedAt   time.Time
}

type UpdateDeviceTypeUsecase struct {
	deviceTypeRepo repository.DeviceTypeRepository
}

func NewUpdateDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository) *UpdateDeviceTypeUsecase {
	return &UpdateDeviceTypeUsecase{
		deviceTypeRepo: deviceTypeRepo,
	}
}

func (u *UpdateDeviceTypeUsecase) Execute(ctx context.Context, req *UpdateDeviceTypeRequest) (*UpdateDeviceTypeResponse, error) {
	existingDeviceType, err := u.deviceTypeRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingDeviceType == nil {
		return nil, ErrDeviceTypeNotFound
	}

	if existingDeviceType.TypeCode != req.TypeCode {
		exists, err := u.deviceTypeRepo.Exists(ctx, req.TypeCode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrDeviceTypeAlreadyExists
		}
	}

	deviceType := &entity.DeviceType{
		ID:          req.ID,
		TypeCode:    req.TypeCode,
		Description: req.Description,
	}

	err = u.deviceTypeRepo.Update(ctx, deviceType)
	if err != nil {
		return nil, err
	}

	response := &UpdateDeviceTypeResponse{
		ID:          deviceType.ID,
		TypeCode:    deviceType.TypeCode,
		Description: deviceType.Description,
	}

	if deviceType.UpdatedAt != nil {
		response.UpdatedAt = *deviceType.UpdatedAt
	}

	return response, nil
}
