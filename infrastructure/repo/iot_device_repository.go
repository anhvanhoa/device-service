package repo

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"strings"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type iotDeviceRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewIoTDeviceRepository(db *pg.DB, helper utils.Helper) repository.IoTDeviceRepository {
	return &iotDeviceRepository{db: db, helper: helper}
}

func (r *iotDeviceRepository) Create(ctx context.Context, device *entity.IoTDevice) error {
	_, err := r.db.Model(device).Context(ctx).Insert()
	return err
}

func (r *iotDeviceRepository) GetByID(ctx context.Context, id string) (*entity.IoTDevice, error) {
	device := &entity.IoTDevice{}
	err := r.db.Model(device).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *iotDeviceRepository) GetByMacAddress(ctx context.Context, macAddress string) (*entity.IoTDevice, error) {
	device := &entity.IoTDevice{}
	err := r.db.Model(device).Context(ctx).Where("mac_address = ?", macAddress).Select()
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *iotDeviceRepository) List(ctx context.Context, filters repository.IoTDeviceFilters, pagination *common.Pagination) ([]*entity.IoTDevice, int64, error) {
	var devices []*entity.IoTDevice

	query := r.db.Model(&devices).Context(ctx)

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

	return devices, int64(total), nil
}

func (r *iotDeviceRepository) Update(ctx context.Context, device *entity.IoTDevice) error {
	_, err := r.db.Model(device).Context(ctx).Where("id = ?", device.ID).UpdateNotZero()
	return err
}

func (r *iotDeviceRepository) UpdateStatus(ctx context.Context, id string, status entity.DeviceStatus) error {
	_, err := r.db.Model((*entity.IoTDevice)(nil)).Context(ctx).
		Set("status = ?", status).
		Set("updated_at = NOW()").
		Where("id = ?", id).UpdateNotZero()
	return err
}

func (r *iotDeviceRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.IoTDevice)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *iotDeviceRepository) GetLowBatteryDevices(ctx context.Context, threshold int, pagination *common.Pagination) ([]*entity.IoTDevice, int64, error) {
	var devices []*entity.IoTDevice

	query := r.db.Model(&devices).Context(ctx).Where("battery_level < ?", threshold)

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	r.applyPagination(query, pagination)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return devices, int64(total), nil
}

func (r *iotDeviceRepository) ExistsByMacAddress(ctx context.Context, macAddress string) (bool, error) {
	count, err := r.db.Model((*entity.IoTDevice)(nil)).Context(ctx).Where("mac_address = ?", macAddress).Count()
	return count > 0, err
}

func (r *iotDeviceRepository) applyFilters(query *pg.Query, filters repository.IoTDeviceFilters) *pg.Query {
	if filters.DeviceTypeID != "" {
		query = query.Where("device_type_id = ?", filters.DeviceTypeID)
	}
	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}
	if filters.GreenhouseID != "" {
		query = query.Where("greenhouse_id = ?", filters.GreenhouseID)
	}
	if filters.GrowingZoneID != "" {
		query = query.Where("growing_zone_id = ?", filters.GrowingZoneID)
	}
	if filters.Search != "" {
		searchTerm := "%" + strings.ToLower(filters.Search) + "%"
		query = query.Where("LOWER(device_name) LIKE ? OR LOWER(model) LIKE ?", searchTerm, searchTerm)
	}
	return query
}

func (r *iotDeviceRepository) applyPagination(query *pg.Query, pagination *common.Pagination) *pg.Query {
	if pagination != nil {
		limit := pagination.PageSize
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Limit(limit).Offset(offset)
	}
	return query
}
