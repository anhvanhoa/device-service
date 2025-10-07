package sensor_data

import (
	"context"
	"device-service/domain/repository"
)

type DeleteSensorDataRequest struct {
	ID string
}

type DeleteSensorDataResponse struct {
	Success bool
	Message string
}

type DeleteSensorDataUsecase struct {
	sensorDataRepo repository.SensorDataRepository
}

func NewDeleteSensorDataUsecase(sensorDataRepo repository.SensorDataRepository) *DeleteSensorDataUsecase {
	return &DeleteSensorDataUsecase{
		sensorDataRepo: sensorDataRepo,
	}
}

func (u *DeleteSensorDataUsecase) Execute(ctx context.Context, req *DeleteSensorDataRequest) (*DeleteSensorDataResponse, error) {

	_, err := u.sensorDataRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	err = u.sensorDataRepo.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteSensorDataResponse{
		Success: true,
		Message: "Xóa dữ liệu cảm biến thành công",
	}, nil
}
