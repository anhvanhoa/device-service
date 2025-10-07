package sensor_data_service

import (
	"context"
	"device-service/domain/usecase/sensor_data"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
)

func (s *SensorDataService) ListSensorData(ctx context.Context, req *proto_sensor_data.ListSensorDataRequest) (*proto_sensor_data.ListSensorDataResponse, error) {
	listRequest := s.convertRequestListSensorData(req)
	listResponse, err := s.sensorDataUsecase.List(ctx, listRequest)
	if err != nil {
		return nil, err
	}
	return s.convertResponseListSensorData(listResponse), nil
}

func (s *SensorDataService) convertRequestListSensorData(req *proto_sensor_data.ListSensorDataRequest) *sensor_data.ListSensorDataRequest {
	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	return &sensor_data.ListSensorDataRequest{
		Pagination: pagination,
	}
}

func (s *SensorDataService) convertResponseListSensorData(response *sensor_data.ListSensorDataResponse) *proto_sensor_data.ListSensorDataResponse {
	sensorDataList := make([]*proto_sensor_data.SensorData, len(response.Data))
	for i, item := range response.Data {
		sensorData := &proto_sensor_data.SensorData{
			Id:         item.ID,
			DeviceId:   item.DeviceID,
			SensorType: item.SensorType,
			Value:      item.Value,
			Unit:       item.Unit,
			RecordedAt: item.RecordedAt,
			IsAlert:    item.IsAlert,
			CreatedAt:  item.CreatedAt,
		}

		if item.QualityScore != nil {
			sensorData.QualityScore = *item.QualityScore
		}

		sensorDataList[i] = sensorData
	}

	pagination := &common_proto.PaginationResponse{
		Total:      int32(response.Total),
		Page:       int32(response.Page),
		PageSize:   int32(response.PageSize),
		TotalPages: int32(response.TotalPages),
	}

	return &proto_sensor_data.ListSensorDataResponse{
		Data:       sensorDataList,
		Pagination: pagination,
	}
}
