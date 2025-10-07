package sensor_data_service

import (
	"context"
	"device-service/domain/usecase/sensor_data"

	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
)

func (s *SensorDataService) GetSensorData(ctx context.Context, req *proto_sensor_data.GetSensorDataRequest) (*proto_sensor_data.GetSensorDataResponse, error) {
	getRequest := s.convertRequestGetSensorData(req)
	getResponse, err := s.sensorDataUsecase.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseGetSensorData(getResponse), nil
}

func (s *SensorDataService) convertRequestGetSensorData(req *proto_sensor_data.GetSensorDataRequest) *sensor_data.GetSensorDataRequest {
	return &sensor_data.GetSensorDataRequest{
		ID: req.Id,
	}
}

func (s *SensorDataService) convertResponseGetSensorData(response *sensor_data.GetSensorDataResponse) *proto_sensor_data.GetSensorDataResponse {
	sensorData := &proto_sensor_data.SensorData{
		Id:         response.ID,
		DeviceId:   response.DeviceID,
		SensorType: response.SensorType,
		Value:      response.Value,
		Unit:       response.Unit,
		RecordedAt: response.RecordedAt,
		IsAlert:    response.IsAlert,
		CreatedAt:  response.CreatedAt,
	}

	if response.QualityScore != nil {
		sensorData.QualityScore = *response.QualityScore
	}

	return &proto_sensor_data.GetSensorDataResponse{
		SensorData: sensorData,
	}
}
