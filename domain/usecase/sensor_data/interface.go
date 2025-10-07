package sensor_data

import "context"

type SensorDataUsecase interface {
	Create(ctx context.Context, req *CreateSensorDataRequest) (*CreateSensorDataResponse, error)
	Get(ctx context.Context, req *GetSensorDataRequest) (*GetSensorDataResponse, error)
	Delete(ctx context.Context, req *DeleteSensorDataRequest) (*DeleteSensorDataResponse, error)
	List(ctx context.Context, req *ListSensorDataRequest) (*ListSensorDataResponse, error)
}
