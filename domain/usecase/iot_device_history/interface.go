package iot_device_history

import "context"

type IoTDeviceHistoryUsecase interface {
	Create(ctx context.Context, req *CreateIoTDeviceHistoryRequest) (*CreateIoTDeviceHistoryResponse, error)
	Get(ctx context.Context, req *GetIoTDeviceHistoryRequest) (*GetIoTDeviceHistoryResponse, error)
	Delete(ctx context.Context, req *DeleteIoTDeviceHistoryRequest) (*DeleteIoTDeviceHistoryResponse, error)
	List(ctx context.Context, req *ListIoTDeviceHistoryRequest) (*ListIoTDeviceHistoryResponse, error)
}
