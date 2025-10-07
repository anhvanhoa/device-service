package iot_device_history

import (
	"context"
	"device-service/domain/repository"
)

type GetIoTDeviceHistoryRequest struct {
	ID string
}

type GetIoTDeviceHistoryResponse struct {
	ID          string
	DeviceID    string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	ActionDate  string
	PerformedBy string
	Notes       *string
}

type GetIoTDeviceHistoryUsecase struct {
	deviceHistoryRepo repository.IoTDeviceHistoryRepository
}

func NewGetIoTDeviceHistoryUsecase(deviceHistoryRepo repository.IoTDeviceHistoryRepository) *GetIoTDeviceHistoryUsecase {
	return &GetIoTDeviceHistoryUsecase{
		deviceHistoryRepo: deviceHistoryRepo,
	}
}

func (u *GetIoTDeviceHistoryUsecase) Execute(ctx context.Context, req *GetIoTDeviceHistoryRequest) (*GetIoTDeviceHistoryResponse, error) {
	history, err := u.deviceHistoryRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if history == nil {
		return nil, ErrDeviceHistoryNotFound
	}

	return &GetIoTDeviceHistoryResponse{
		ID:          history.ID,
		DeviceID:    history.DeviceID,
		Action:      string(history.Action),
		OldValue:    map[string]any(history.OldValue),
		NewValue:    map[string]any(history.NewValue),
		ActionDate:  history.ActionDate.Format("2006-01-02T15:04:05Z07:00"),
		PerformedBy: history.PerformedBy,
		Notes:       history.Notes,
	}, nil
}
