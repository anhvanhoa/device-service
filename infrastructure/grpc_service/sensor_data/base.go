package sensor_data_service

import (
	"device-service/domain/repository"
	"device-service/domain/usecase/sensor_data"

	"github.com/anhvanhoa/service-core/utils"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
)

type SensorDataService struct {
	proto_sensor_data.UnimplementedSensorDataServiceServer
	sensorDataUsecase sensor_data.SensorDataUsecase
}

func NewSensorDataService(sensorDataRepo repository.SensorDataRepository, helper utils.Helper) proto_sensor_data.SensorDataServiceServer {
	sensorDataUsecase := sensor_data.NewSensorDataUsecase(sensorDataRepo, helper)
	return &SensorDataService{sensorDataUsecase: sensorDataUsecase}
}
