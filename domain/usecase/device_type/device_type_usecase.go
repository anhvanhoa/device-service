package device_type

import (
	"context"
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
)

type DeviceTypeUsecaseImpl struct {
	createUsecase *CreateDeviceTypeUsecase
	getUsecase    *GetDeviceTypeUsecase
	updateUsecase *UpdateDeviceTypeUsecase
	deleteUsecase *DeleteDeviceTypeUsecase
	listUsecase   *ListDeviceTypeUsecase
}

func NewDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository, helper utils.Helper) DeviceTypeUsecase {
	return &DeviceTypeUsecaseImpl{
		createUsecase: NewCreateDeviceTypeUsecase(deviceTypeRepo),
		getUsecase:    NewGetDeviceTypeUsecase(deviceTypeRepo),
		updateUsecase: NewUpdateDeviceTypeUsecase(deviceTypeRepo),
		deleteUsecase: NewDeleteDeviceTypeUsecase(deviceTypeRepo),
		listUsecase:   NewListDeviceTypeUsecase(deviceTypeRepo, helper),
	}
}

func (u *DeviceTypeUsecaseImpl) Create(ctx context.Context, req *CreateDeviceTypeRequest) (*CreateDeviceTypeResponse, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *DeviceTypeUsecaseImpl) Get(ctx context.Context, req *GetDeviceTypeRequest) (*GetDeviceTypeResponse, error) {
	return u.getUsecase.Execute(ctx, req)
}

func (u *DeviceTypeUsecaseImpl) Update(ctx context.Context, req *UpdateDeviceTypeRequest) (*UpdateDeviceTypeResponse, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *DeviceTypeUsecaseImpl) Delete(ctx context.Context, req *DeleteDeviceTypeRequest) (*DeleteDeviceTypeResponse, error) {
	return u.deleteUsecase.Execute(ctx, req)
}

func (u *DeviceTypeUsecaseImpl) List(ctx context.Context, req *ListDeviceTypeRequest) (*ListDeviceTypeResponse, error) {
	return u.listUsecase.Execute(ctx, req)
}
