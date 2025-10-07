package iot_device

import "context"

type IoTDeviceUsecase interface {
	Create(ctx context.Context, req *CreateIoTDeviceRequest) (*CreateIoTDeviceResponse, error)
	Get(ctx context.Context, req *GetIoTDeviceRequest) (*GetIoTDeviceResponse, error)
	Update(ctx context.Context, req *UpdateIoTDeviceRequest) (*UpdateIoTDeviceResponse, error)
	Delete(ctx context.Context, req *DeleteIoTDeviceRequest) (*DeleteIoTDeviceResponse, error)
	List(ctx context.Context, req *ListIoTDeviceRequest) (*ListIoTDeviceResponse, error)
}
