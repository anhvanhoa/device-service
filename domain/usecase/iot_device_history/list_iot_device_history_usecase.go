package iot_device_history

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListIoTDeviceHistoryRequest struct {
	Filters    IoTDeviceHistoryFilters
	Pagination *common.Pagination
}

type ListIoTDeviceHistoryResponse common.PaginationResult[IoTDeviceHistoryItem]

type IoTDeviceHistoryItem struct {
	ID          string
	DeviceID    string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	ActionDate  string
	PerformedBy string
	Notes       *string
}

type IoTDeviceHistoryFilters struct {
	DeviceID  *string
	Action    *string
	StartDate *string
	EndDate   *string
}

type ListIoTDeviceHistoryUsecase struct {
	deviceHistoryRepo repository.IoTDeviceHistoryRepository
	helper            utils.Helper
}

func NewListIoTDeviceHistoryUsecase(deviceHistoryRepo repository.IoTDeviceHistoryRepository, helper utils.Helper) *ListIoTDeviceHistoryUsecase {
	return &ListIoTDeviceHistoryUsecase{
		deviceHistoryRepo: deviceHistoryRepo,
		helper:            helper,
	}
}

func (u *ListIoTDeviceHistoryUsecase) Execute(ctx context.Context, req *ListIoTDeviceHistoryRequest) (*ListIoTDeviceHistoryResponse, error) {
	pagination := req.Pagination
	if pagination == nil {
		pagination = &common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}

	filters := repository.IoTDeviceHistoryFilters{
		DeviceID: req.Filters.DeviceID,
	}

	if req.Filters.Action != nil {
		action := entity.DeviceAction(*req.Filters.Action)
		filters.Action = &action
	}

	if req.Filters.StartDate != nil {
		startDate, err := time.Parse("2006-01-02", *req.Filters.StartDate)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}
		filters.StartDate = &startDate
	}

	if req.Filters.EndDate != nil {
		endDate, err := time.Parse("2006-01-02", *req.Filters.EndDate)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}
		filters.EndDate = &endDate
	}

	histories, total, err := u.deviceHistoryRepo.List(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}

	items := make([]IoTDeviceHistoryItem, len(histories))
	for i, history := range histories {
		item := IoTDeviceHistoryItem{
			ID:          history.ID,
			DeviceID:    history.DeviceID,
			Action:      string(history.Action),
			OldValue:    map[string]any(history.OldValue),
			NewValue:    map[string]any(history.NewValue),
			ActionDate:  history.ActionDate.Format("2006-01-02T15:04:05Z07:00"),
			PerformedBy: history.PerformedBy,
			Notes:       history.Notes,
		}

		items[i] = item
	}
	return &ListIoTDeviceHistoryResponse{
		Data:       items,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(int64(total), int64(pagination.PageSize)),
	}, nil
}
