package repository

import (
	"context"
	"device-service/domain/entity"
	"time"

	"github.com/anhvanhoa/service-core/common"
)

type SensorDataRepository interface {
	Create(ctx context.Context, sensorData *entity.SensorData) error
	CreateBatch(ctx context.Context, sensorDataList []*entity.SensorData) error
	GetByID(ctx context.Context, id string) (*entity.SensorData, error)
	ListByDeviceID(ctx context.Context, deviceID string, pagination *common.Pagination) ([]*entity.SensorData, int64, error)
	ListByDateRange(ctx context.Context, deviceID string, startDate, endDate time.Time, pagination *common.Pagination) ([]*entity.SensorData, int64, error)
	ListAlerts(ctx context.Context, pagination *common.Pagination) ([]*entity.SensorData, int64, error)
	List(ctx context.Context, filters SensorDataFilters, pagination *common.Pagination) ([]*entity.SensorData, int64, error)
	GetLatestByDeviceID(ctx context.Context, deviceID string) (*entity.SensorData, error)
	GetAverageByDateRange(ctx context.Context, deviceID string, sensorType entity.SensorType, startDate, endDate time.Time) (float64, error)
	Delete(ctx context.Context, id string) error
	DeleteOldRecords(ctx context.Context, beforeDate time.Time) (int64, error)
}

type SensorDataFilters struct {
	DeviceID   *string
	SensorType *entity.SensorType
	IsAlert    *bool
	StartDate  *time.Time
	EndDate    *time.Time
}
