package sensor_data_service

import (
	"context"
	"device-service/domain/usecase/sensor_data"

	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
)

func (s *SensorDataService) DeleteSensorData(ctx context.Context, req *proto_sensor_data.DeleteSensorDataRequest) (*proto_sensor_data.DeleteSensorDataResponse, error) {
	deleteRequest := s.convertRequestDeleteSensorData(req)
	deleteResponse, err := s.sensorDataUsecase.Delete(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseDeleteSensorData(deleteResponse), nil
}

func (s *SensorDataService) convertRequestDeleteSensorData(req *proto_sensor_data.DeleteSensorDataRequest) *sensor_data.DeleteSensorDataRequest {
	return &sensor_data.DeleteSensorDataRequest{
		ID: req.Id,
	}
}

func (s *SensorDataService) convertResponseDeleteSensorData(response *sensor_data.DeleteSensorDataResponse) *proto_sensor_data.DeleteSensorDataResponse {
	return &proto_sensor_data.DeleteSensorDataResponse{
		Success: response.Success,
	}
}
