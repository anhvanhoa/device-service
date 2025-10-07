package device_type

import "context"

type DeviceTypeUsecase interface {
	Create(ctx context.Context, req *CreateDeviceTypeRequest) (*CreateDeviceTypeResponse, error)
	Get(ctx context.Context, req *GetDeviceTypeRequest) (*GetDeviceTypeResponse, error)
	Update(ctx context.Context, req *UpdateDeviceTypeRequest) (*UpdateDeviceTypeResponse, error)
	Delete(ctx context.Context, req *DeleteDeviceTypeRequest) (*DeleteDeviceTypeResponse, error)
	List(ctx context.Context, req *ListDeviceTypeRequest) (*ListDeviceTypeResponse, error)
}
