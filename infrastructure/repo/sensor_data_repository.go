package repo

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type sensorDataRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewSensorDataRepository(db *pg.DB, helper utils.Helper) repository.SensorDataRepository {
	return &sensorDataRepository{db: db, helper: helper}
}

func (r *sensorDataRepository) Create(ctx context.Context, sensorData *entity.SensorData) error {
	_, err := r.db.Model(sensorData).Context(ctx).Insert()
	return err
}

func (r *sensorDataRepository) CreateBatch(ctx context.Context, sensorDataList []*entity.SensorData) error {
	if len(sensorDataList) == 0 {
		return nil
	}
	_, err := r.db.Model(&sensorDataList).Context(ctx).Insert()
	return err
}

func (r *sensorDataRepository) GetByID(ctx context.Context, id string) (*entity.SensorData, error) {
	sensorData := &entity.SensorData{}
	err := r.db.Model(sensorData).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return sensorData, nil
}

func (r *sensorDataRepository) ListByDeviceID(ctx context.Context, deviceID string, pagination *common.Pagination) ([]*entity.SensorData, int64, error) {
	var sensorDataList []*entity.SensorData

	query := r.db.Model(&sensorDataList).Context(ctx).Where("device_id = ?", deviceID).Order("recorded_at DESC")

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	query = r.applyPagination(query, pagination)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return sensorDataList, int64(total), nil
}

func (r *sensorDataRepository) ListByDateRange(ctx context.Context, deviceID string, startDate, endDate time.Time, pagination *common.Pagination) ([]*entity.SensorData, int64, error) {
	var sensorDataList []*entity.SensorData

	query := r.db.Model(&sensorDataList).Context(ctx).
		Where("device_id = ?", deviceID).
		Where("recorded_at >= ?", startDate).
		Where("recorded_at <= ?", endDate).
		Order("recorded_at DESC")

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	query = r.applyPagination(query, pagination)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return sensorDataList, int64(total), nil
}

func (r *sensorDataRepository) ListAlerts(ctx context.Context, pagination *common.Pagination) ([]*entity.SensorData, int64, error) {
	var sensorDataList []*entity.SensorData

	query := r.db.Model(&sensorDataList).Context(ctx).Where("is_alert = ?", true).Order("recorded_at DESC")

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	query = r.applyPagination(query, pagination)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return sensorDataList, int64(total), nil
}

func (r *sensorDataRepository) List(ctx context.Context, filters repository.SensorDataFilters, pagination *common.Pagination) ([]*entity.SensorData, int64, error) {
	var sensorDataList []*entity.SensorData

	query := r.db.Model(&sensorDataList).Context(ctx)
	query = r.applyFilters(query, filters)
	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	query = r.applyPagination(query, pagination)
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return sensorDataList, int64(total), nil
}

func (r *sensorDataRepository) GetLatestByDeviceID(ctx context.Context, deviceID string) (*entity.SensorData, error) {
	sensorData := &entity.SensorData{}
	err := r.db.Model(sensorData).Context(ctx).Where("device_id = ?", deviceID).Select()
	if err != nil {
		return nil, err
	}
	return sensorData, nil
}

func (r *sensorDataRepository) GetAverageByDateRange(ctx context.Context, deviceID string, sensorType entity.SensorType, startDate, endDate time.Time) (float64, error) {
	var avg float64
	err := r.db.Model((*entity.SensorData)(nil)).Context(ctx).
		ColumnExpr("AVG(value)").
		Where("device_id = ?", deviceID).
		Where("sensor_type = ?", sensorType).
		Where("recorded_at >= ?", startDate).
		Where("recorded_at <= ?", endDate).
		Where("value IS NOT NULL").
		Select(&avg)
	return avg, err
}

func (r *sensorDataRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.SensorData)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *sensorDataRepository) DeleteOldRecords(ctx context.Context, beforeDate time.Time) (int64, error) {
	result, err := r.db.Model((*entity.SensorData)(nil)).Context(ctx).
		Where("recorded_at < ?", beforeDate).
		Delete()
	if err != nil {
		return 0, err
	}
	return int64(result.RowsAffected()), nil
}

func (r *sensorDataRepository) applyFilters(query *pg.Query, filters repository.SensorDataFilters) *pg.Query {
	if filters.DeviceID != nil {
		query = query.Where("device_id = ?", *filters.DeviceID)
	}
	if filters.SensorType != nil {
		query = query.Where("sensor_type = ?", *filters.SensorType)
	}
	if filters.IsAlert != nil {
		query = query.Where("is_alert = ?", *filters.IsAlert)
	}
	if filters.StartDate != nil {
		query = query.Where("recorded_at >= ?", *filters.StartDate)
	}
	if filters.EndDate != nil {
		query = query.Where("recorded_at <= ?", *filters.EndDate)
	}
	return query
}

func (r *sensorDataRepository) applyPagination(query *pg.Query, pagination *common.Pagination) *pg.Query {
	query = query.Order("recorded_at DESC")
	if pagination != nil {
		limit := pagination.PageSize
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Limit(limit).Offset(offset)
	}
	return query
}
