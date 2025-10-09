package iot_device

import (
	"context"
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
)

type IoTDeviceUsecaseImpl struct {
	createUsecase  *CreateIoTDeviceUsecase
	getUsecase     *GetIoTDeviceUsecase
	updateUsecase  *UpdateIoTDeviceUsecase
	deleteUsecase  *DeleteIoTDeviceUsecase
	listUsecase    *ListIoTDeviceUsecase
	controlUsecase *ControlIoTDeviceUsecase
}

func NewIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository, helper utils.Helper, mq mq.MQI) IoTDeviceUsecase {
	return &IoTDeviceUsecaseImpl{
		createUsecase:  NewCreateIoTDeviceUsecase(iotDeviceRepo),
		getUsecase:     NewGetIoTDeviceUsecase(iotDeviceRepo),
		updateUsecase:  NewUpdateIoTDeviceUsecase(iotDeviceRepo),
		deleteUsecase:  NewDeleteIoTDeviceUsecase(iotDeviceRepo),
		listUsecase:    NewListIoTDeviceUsecase(iotDeviceRepo, helper),
		controlUsecase: NewControlIoTDeviceUsecase(iotDeviceRepo, mq),
	}
}

func (u *IoTDeviceUsecaseImpl) Create(ctx context.Context, req *CreateIoTDeviceRequest) (*CreateIoTDeviceResponse, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *IoTDeviceUsecaseImpl) Get(ctx context.Context, req *GetIoTDeviceRequest) (*GetIoTDeviceResponse, error) {
	return u.getUsecase.Execute(ctx, req)
}

func (u *IoTDeviceUsecaseImpl) Update(ctx context.Context, req *UpdateIoTDeviceRequest) (*UpdateIoTDeviceResponse, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *IoTDeviceUsecaseImpl) Delete(ctx context.Context, req *DeleteIoTDeviceRequest) (*DeleteIoTDeviceResponse, error) {
	return u.deleteUsecase.Execute(ctx, req)
}

func (u *IoTDeviceUsecaseImpl) List(ctx context.Context, req *ListIoTDeviceRequest) (*ListIoTDeviceResponse, error) {
	return u.listUsecase.Execute(ctx, req)
}

func (u *IoTDeviceUsecaseImpl) Control(ctx context.Context, req *ControlIoTDeviceRequest) (*ControlIoTDeviceResponse, error) {
	return u.controlUsecase.Execute(ctx, req)
}
