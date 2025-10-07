package iot_device_history

import (
	"context"
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
)

type IoTDeviceHistoryUsecaseImpl struct {
	createUsecase *CreateIoTDeviceHistoryUsecase
	getUsecase    *GetIoTDeviceHistoryUsecase
	listUsecase   *ListIoTDeviceHistoryUsecase
	deleteUsecase *DeleteIoTDeviceHistoryUsecase
}

func NewIoTDeviceHistoryUsecase(deviceHistoryRepo repository.IoTDeviceHistoryRepository, helper utils.Helper) IoTDeviceHistoryUsecase {
	return &IoTDeviceHistoryUsecaseImpl{
		createUsecase: NewCreateIoTDeviceHistoryUsecase(deviceHistoryRepo),
		getUsecase:    NewGetIoTDeviceHistoryUsecase(deviceHistoryRepo),
		listUsecase:   NewListIoTDeviceHistoryUsecase(deviceHistoryRepo, helper),
		deleteUsecase: NewDeleteIoTDeviceHistoryUsecase(deviceHistoryRepo),
	}
}

func (u *IoTDeviceHistoryUsecaseImpl) Create(ctx context.Context, req *CreateIoTDeviceHistoryRequest) (*CreateIoTDeviceHistoryResponse, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *IoTDeviceHistoryUsecaseImpl) Get(ctx context.Context, req *GetIoTDeviceHistoryRequest) (*GetIoTDeviceHistoryResponse, error) {
	return u.getUsecase.Execute(ctx, req)
}

func (u *IoTDeviceHistoryUsecaseImpl) List(ctx context.Context, req *ListIoTDeviceHistoryRequest) (*ListIoTDeviceHistoryResponse, error) {
	return u.listUsecase.Execute(ctx, req)
}

func (u *IoTDeviceHistoryUsecaseImpl) Delete(ctx context.Context, req *DeleteIoTDeviceHistoryRequest) (*DeleteIoTDeviceHistoryResponse, error) {
	return u.deleteUsecase.Execute(ctx, req)
}
