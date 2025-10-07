package repository

import (
	"context"
	"device-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type DeviceTypeRepository interface {
	Create(ctx context.Context, deviceType *entity.DeviceType) error
	GetByID(ctx context.Context, id string) (*entity.DeviceType, error)
	GetByTypeCode(ctx context.Context, typeCode string) (*entity.DeviceType, error)
	List(ctx context.Context, pagination *common.Pagination) ([]*entity.DeviceType, int64, error)
	Update(ctx context.Context, deviceType *entity.DeviceType) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	Exists(ctx context.Context, typeCode string) (bool, error)
}
