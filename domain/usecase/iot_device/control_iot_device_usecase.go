package iot_device

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"fmt"

	"github.com/anhvanhoa/service-core/domain/mq"
)

type ControlIoTDeviceUsecase struct {
	iotDeviceRepo repository.IoTDeviceRepository
	mq            mq.MQI
}

func NewControlIoTDeviceUsecase(iotDeviceRepo repository.IoTDeviceRepository, mq mq.MQI) *ControlIoTDeviceUsecase {
	return &ControlIoTDeviceUsecase{
		iotDeviceRepo: iotDeviceRepo,
		mq:            mq,
	}
}

type ControlIoTDeviceRequest struct {
	DeviceID string
	Action   string
}

type ControlIoTDeviceResponse struct {
	DeviceID string
	Status   entity.DeviceStatus
	Action   string
	Message  string
}

func (u *ControlIoTDeviceUsecase) Execute(ctx context.Context, req *ControlIoTDeviceRequest) (*ControlIoTDeviceResponse, error) {
	if err := validateControlRequest(req); err != nil {
		return nil, err
	}

	device, err := u.iotDeviceRepo.GetByID(ctx, req.DeviceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %v", err)
	}

	if device == nil {
		return nil, ErrDeviceNotFound
	}

	var newStatus entity.DeviceStatus
	var message string

	switch req.Action {
	case "on":
		if device.Status == entity.DeviceStatusActive {
			return nil, ErrDeviceAlreadyOn
		}
		newStatus = entity.DeviceStatusActive
		message = "Thiết bị đã được bật"
	case "off":
		if device.Status == entity.DeviceStatusInactive {
			return nil, ErrDeviceAlreadyOff
		}
		newStatus = entity.DeviceStatusInactive
		message = "Thiết bị đã được tắt"
	case "toggle":
		if device.Status == entity.DeviceStatusActive {
			newStatus = entity.DeviceStatusInactive
			message = "Thiết bị đã được tắt"
		} else {
			newStatus = entity.DeviceStatusActive
			message = "Thiết bị đã được bật"
		}
	case "reset":
		newStatus = entity.DeviceStatusActive
		message = "Thiết bị đã được reset"
	default:
		return nil, ErrInvalidAction
	}

	err = u.iotDeviceRepo.UpdateStatus(ctx, req.DeviceID, newStatus)
	if err != nil {
		return nil, err
	}

	token := u.mq.Publish("iot_device/control", 0, req)
	if token.Wait() && token.Error() != nil {
		return nil, ErrPublishFailed
	}

	return &ControlIoTDeviceResponse{
		DeviceID: req.DeviceID,
		Status:   newStatus,
		Action:   req.Action,
		Message:  message,
	}, nil
}

func validateControlRequest(req *ControlIoTDeviceRequest) error {
	if req.DeviceID == "" {
		return ErrDeviceIDRequired
	}

	if req.Action == "" {
		return ErrActionRequired
	}

	validActions := map[string]bool{
		"on":     true,
		"off":    true,
		"toggle": true,
		"reset":  true,
	}

	if !validActions[req.Action] {
		return ErrInvalidAction
	}

	return nil
}
