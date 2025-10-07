package iot_device_history

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
)

type CreateIoTDeviceHistoryRequest struct {
	DeviceID    string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	PerformedBy string
	Notes       *string
}

type CreateIoTDeviceHistoryResponse struct {
	ID          string
	DeviceID    string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	ActionDate  string
	PerformedBy string
	Notes       *string
}

type CreateIoTDeviceHistoryUsecase struct {
	deviceHistoryRepo repository.IoTDeviceHistoryRepository
}

func NewCreateIoTDeviceHistoryUsecase(deviceHistoryRepo repository.IoTDeviceHistoryRepository) *CreateIoTDeviceHistoryUsecase {
	return &CreateIoTDeviceHistoryUsecase{
		deviceHistoryRepo: deviceHistoryRepo,
	}
}

func (u *CreateIoTDeviceHistoryUsecase) Execute(ctx context.Context, req *CreateIoTDeviceHistoryRequest) (*CreateIoTDeviceHistoryResponse, error) {
	action := entity.DeviceAction(req.Action)
	if action != entity.DeviceActionInstall && action != entity.DeviceActionUpdateConfig &&
		action != entity.DeviceActionFirmwareUpdate && action != entity.DeviceActionRelocate &&
		action != entity.DeviceActionMaintenance && action != entity.DeviceActionDeactivate &&
		action != entity.DeviceActionReactivate {
		return nil, ErrInvalidAction
	}

	history := &entity.IoTDeviceHistory{
		DeviceID:    req.DeviceID,
		Action:      action,
		OldValue:    entity.JSONB(req.OldValue),
		NewValue:    entity.JSONB(req.NewValue),
		PerformedBy: req.PerformedBy,
		Notes:       req.Notes,
	}

	err := u.deviceHistoryRepo.Create(ctx, history)
	if err != nil {
		return nil, err
	}

	return &CreateIoTDeviceHistoryResponse{
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
