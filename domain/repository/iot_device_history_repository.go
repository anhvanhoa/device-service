package repository

import (
	"context"
	"device-service/domain/entity"
	"time"

	"github.com/anhvanhoa/service-core/common"
)

type IoTDeviceHistoryRepository interface {
	Create(ctx context.Context, history *entity.IoTDeviceHistory) error
	GetByID(ctx context.Context, id string) (*entity.IoTDeviceHistory, error)
	ListByDeviceID(ctx context.Context, deviceID string, pagination *common.Pagination) ([]*entity.IoTDeviceHistory, int64, error)
	List(ctx context.Context, filters IoTDeviceHistoryFilters, pagination *common.Pagination) ([]*entity.IoTDeviceHistory, int64, error)
	GetLatestByDeviceID(ctx context.Context, deviceID string) (*entity.IoTDeviceHistory, error)
	Delete(ctx context.Context, id string) error
	DeleteOldRecords(ctx context.Context, beforeDate time.Time) (int64, error)
}

type IoTDeviceHistoryFilters struct {
	DeviceID    string
	Action      entity.DeviceAction
	PerformedBy string
	StartDate   *time.Time
	EndDate     *time.Time
}
