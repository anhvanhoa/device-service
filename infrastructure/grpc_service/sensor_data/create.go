package sensor_data_service

import (
	"context"
	"device-service/domain/usecase/sensor_data"

	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
)

func (s *SensorDataService) CreateSensorData(ctx context.Context, req *proto_sensor_data.CreateSensorDataRequest) (*proto_sensor_data.CreateSensorDataResponse, error) {
	createRequest := s.convertRequestCreateSensorData(req)
	createResponse, err := s.sensorDataUsecase.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseCreateSensorData(createResponse), nil
}

func (s *SensorDataService) convertRequestCreateSensorData(req *proto_sensor_data.CreateSensorDataRequest) *sensor_data.CreateSensorDataRequest {
	request := &sensor_data.CreateSensorDataRequest{
		DeviceID:   req.DeviceId,
		SensorType: req.SensorType,
		Value:      req.Value,
		Unit:       req.Unit,
		IsAlert:    req.IsAlert,
	}

	if req.QualityScore != 0 {
		request.QualityScore = &req.QualityScore
	}

	return request
}

func (s *SensorDataService) convertResponseCreateSensorData(response *sensor_data.CreateSensorDataResponse) *proto_sensor_data.CreateSensorDataResponse {
	resp := &proto_sensor_data.CreateSensorDataResponse{
		DeviceId:   response.DeviceID,
		SensorType: response.SensorType,
		Value:      response.Value,
		Unit:       response.Unit,
		RecordedAt: response.RecordedAt,
		IsAlert:    response.IsAlert,
		CreatedAt:  response.CreatedAt,
	}

	if response.QualityScore != nil {
		resp.QualityScore = *response.QualityScore
	}

	return resp
}
