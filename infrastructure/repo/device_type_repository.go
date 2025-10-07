package repo

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type deviceTypeRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewDeviceTypeRepository(db *pg.DB, helper utils.Helper) repository.DeviceTypeRepository {
	return &deviceTypeRepository{db: db, helper: helper}
}

func (r *deviceTypeRepository) Create(ctx context.Context, deviceType *entity.DeviceType) error {
	_, err := r.db.Model(deviceType).Context(ctx).Insert()
	return err
}

func (r *deviceTypeRepository) GetByID(ctx context.Context, id string) (*entity.DeviceType, error) {
	deviceType := &entity.DeviceType{}
	err := r.db.Model(deviceType).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return deviceType, nil
}

func (r *deviceTypeRepository) GetByTypeCode(ctx context.Context, typeCode string) (*entity.DeviceType, error) {
	deviceType := &entity.DeviceType{}
	err := r.db.Model(deviceType).Context(ctx).Where("type_code = ?", typeCode).Select()
	if err != nil {
		return nil, err
	}
	return deviceType, nil
}

func (r *deviceTypeRepository) List(ctx context.Context, pagination *common.Pagination) ([]*entity.DeviceType, int64, error) {
	var deviceTypes []*entity.DeviceType

	query := r.db.Model(&deviceTypes).Context(ctx)

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

	return deviceTypes, int64(total), nil
}

func (r *deviceTypeRepository) Update(ctx context.Context, deviceType *entity.DeviceType) error {
	_, err := r.db.Model(deviceType).Context(ctx).Where("id = ?", deviceType.ID).UpdateNotZero()
	return err
}

func (r *deviceTypeRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.DeviceType)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *deviceTypeRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model((*entity.DeviceType)(nil)).Context(ctx).Count()
	return int64(count), err
}

func (r *deviceTypeRepository) Exists(ctx context.Context, typeCode string) (bool, error) {
	count, err := r.db.Model((*entity.DeviceType)(nil)).Context(ctx).Where("type_code = ?", typeCode).Count()
	return count > 0, err
}
