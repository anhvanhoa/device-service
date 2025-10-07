package sensor_data

import (
	"context"
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
)

type SensorDataUsecaseImpl struct {
	createUsecase *CreateSensorDataUsecase
	getUsecase    *GetSensorDataUsecase
	listUsecase   *ListSensorDataUsecase
	deleteUsecase *DeleteSensorDataUsecase
}

func NewSensorDataUsecase(sensorDataRepo repository.SensorDataRepository, helper utils.Helper) SensorDataUsecase {
	return &SensorDataUsecaseImpl{
		createUsecase: NewCreateSensorDataUsecase(sensorDataRepo),
		getUsecase:    NewGetSensorDataUsecase(sensorDataRepo),
		listUsecase:   NewListSensorDataUsecase(sensorDataRepo, helper),
		deleteUsecase: NewDeleteSensorDataUsecase(sensorDataRepo),
	}
}

func (u *SensorDataUsecaseImpl) Create(ctx context.Context, req *CreateSensorDataRequest) (*CreateSensorDataResponse, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *SensorDataUsecaseImpl) Get(ctx context.Context, req *GetSensorDataRequest) (*GetSensorDataResponse, error) {
	return u.getUsecase.Execute(ctx, req)
}

func (u *SensorDataUsecaseImpl) List(ctx context.Context, req *ListSensorDataRequest) (*ListSensorDataResponse, error) {
	return u.listUsecase.Execute(ctx, req)
}

func (u *SensorDataUsecaseImpl) Delete(ctx context.Context, req *DeleteSensorDataRequest) (*DeleteSensorDataResponse, error) {
	return u.deleteUsecase.Execute(ctx, req)
}
