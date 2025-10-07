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

type iotDeviceHistoryRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewIoTDeviceHistoryRepository(db *pg.DB, helper utils.Helper) repository.IoTDeviceHistoryRepository {
	return &iotDeviceHistoryRepository{db: db, helper: helper}
}

func (r *iotDeviceHistoryRepository) Create(ctx context.Context, history *entity.IoTDeviceHistory) error {
	_, err := r.db.Model(history).Context(ctx).Insert()
	return err
}

func (r *iotDeviceHistoryRepository) GetByID(ctx context.Context, id string) (*entity.IoTDeviceHistory, error) {
	history := &entity.IoTDeviceHistory{}
	err := r.db.Model(history).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (r *iotDeviceHistoryRepository) ListByDeviceID(ctx context.Context, deviceID string, pagination *common.Pagination) ([]*entity.IoTDeviceHistory, int64, error) {
	var histories []*entity.IoTDeviceHistory

	query := r.db.Model(&histories).Context(ctx).Where("device_id = ?", deviceID).Order("action_date DESC")

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination != nil {
		limit := pagination.PageSize
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Limit(limit).Offset(offset)
	}

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return histories, int64(total), nil
}

func (r *iotDeviceHistoryRepository) List(ctx context.Context, filters repository.IoTDeviceHistoryFilters, pagination *common.Pagination) ([]*entity.IoTDeviceHistory, int64, error) {
	var histories []*entity.IoTDeviceHistory

	query := r.db.Model(&histories).Context(ctx)

	if filters.DeviceID != "" {
		query = query.Where("device_id = ?", filters.DeviceID)
	}
	if filters.Action != "" {
		query = query.Where("action = ?", filters.Action)
	}
	if filters.PerformedBy != "" {
		query = query.Where("performed_by = ?", filters.PerformedBy)
	}
	if filters.StartDate != nil {
		query = query.Where("action_date >= ?", *filters.StartDate)
	}
	if filters.EndDate != nil {
		query = query.Where("action_date <= ?", *filters.EndDate)
	}

	query = query.Order("action_date DESC")

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination != nil {
		limit := pagination.PageSize
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Limit(limit).Offset(offset)
	}

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return histories, int64(total), nil
}

func (r *iotDeviceHistoryRepository) GetLatestByDeviceID(ctx context.Context, deviceID string) (*entity.IoTDeviceHistory, error) {
	history := &entity.IoTDeviceHistory{}
	err := r.db.Model(history).Context(ctx).
		Where("device_id = ?", deviceID).
		Order("action_date DESC").
		Limit(1).
		Select()
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (r *iotDeviceHistoryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.IoTDeviceHistory)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *iotDeviceHistoryRepository) DeleteOldRecords(ctx context.Context, beforeDate time.Time) (int64, error) {
	result, err := r.db.Model((*entity.IoTDeviceHistory)(nil)).Context(ctx).
		Where("action_date < ?", beforeDate).
		Delete()
	if err != nil {
		return 0, err
	}
	return int64(result.RowsAffected()), nil
}
