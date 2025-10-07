package device_type

import (
	"context"
	"device-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListDeviceTypeRequest struct {
	Pagination *common.Pagination
}

type ListDeviceTypeResponse common.PaginationResult[DeviceTypeItem]

type DeviceTypeItem struct {
	ID          string
	TypeCode    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type ListDeviceTypeUsecase struct {
	deviceTypeRepo repository.DeviceTypeRepository
	helper         utils.Helper
}

func NewListDeviceTypeUsecase(deviceTypeRepo repository.DeviceTypeRepository, helper utils.Helper) *ListDeviceTypeUsecase {
	return &ListDeviceTypeUsecase{
		deviceTypeRepo: deviceTypeRepo,
		helper:         helper,
	}
}

func (u *ListDeviceTypeUsecase) Execute(ctx context.Context, req *ListDeviceTypeRequest) (*ListDeviceTypeResponse, error) {
	pagination := req.Pagination
	if pagination == nil {
		pagination = &common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}

	deviceTypes, total, err := u.deviceTypeRepo.List(ctx, pagination)
	if err != nil {
		return nil, err
	}

	items := make([]DeviceTypeItem, len(deviceTypes))
	for i, deviceType := range deviceTypes {
		item := DeviceTypeItem{
			ID:          deviceType.ID,
			TypeCode:    deviceType.TypeCode,
			Description: deviceType.Description,
			CreatedAt:   deviceType.CreatedAt,
		}

		if deviceType.UpdatedAt != nil {
			item.UpdatedAt = deviceType.UpdatedAt
		}

		items[i] = item
	}

	return &ListDeviceTypeResponse{
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(int64(total), int64(pagination.PageSize)),
		Data:       items,
	}, nil
}
