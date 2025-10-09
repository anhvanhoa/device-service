package iot_device_mqtt

import (
	"context"
	"device-service/domain/usecase/iot_device"
	"device-service/infrastructure/repo"
	"fmt"

	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
)

type IoTDeviceMQTTService interface {
	RegisterDevice(ctx context.Context, msg any) error
	UpdateDevice(ctx context.Context, msg any) error
}

type IoTDeviceMQTT struct {
	iotDeviceUsecase iot_device.IoTDeviceUsecase
	log              *log.LogGRPCImpl
}

func NewIoTDeviceMQTT(repo repo.Repositories, helper utils.Helper, mq mq.MQI, log *log.LogGRPCImpl) IoTDeviceMQTTService {
	return &IoTDeviceMQTT{
		iotDeviceUsecase: iot_device.NewIoTDeviceUsecase(repo.IoTDevice(), helper, mq),
		log:              log,
	}
}

func (s *IoTDeviceMQTT) RegisterDevice(ctx context.Context, msg any) error {
	if err := s.checkMessage(msg); err != nil {
		s.log.Error("RegisterDevice dữ liệu không hợp lệ")
		return err
	}
	req, err := s.getMessageRegister(msg)
	if err != nil {
		s.log.Error("RegisterDevice dữ liệu không hợp lệ")
		return err
	}

	_, err = s.iotDeviceUsecase.Create(ctx, &req)
	if err != nil {
		s.log.Error(fmt.Sprintf("RegisterDevice lỗi: %v", err))
		return err
	}

	return nil
}

func (s *IoTDeviceMQTT) UpdateDevice(ctx context.Context, msg any) error {
	if err := s.checkMessage(msg); err != nil {
		s.log.Error("UpdateDevice dữ liệu không hợp lệ")
		return err
	}

	req, err := s.getMessageUpdate(msg)
	if err != nil {
		s.log.Error(fmt.Sprintf("UpdateDevice lỗi: %v", err))
		return err
	}

	_, err = s.iotDeviceUsecase.Update(ctx, &req)
	if err != nil {
		s.log.Error(fmt.Sprintf("UpdateDevice lỗi: %v", err))
		return err
	}
	return nil
}

func (s *IoTDeviceMQTT) checkMessage(message any) error {
	if message == nil {
		return ErrMessageNil
	}
	return nil
}

func (s *IoTDeviceMQTT) getMessageRegister(message any) (iot_device.CreateIoTDeviceRequest, error) {
	if _, ok := message.(iot_device.CreateIoTDeviceRequest); !ok {
		return iot_device.CreateIoTDeviceRequest{}, ErrMessageNil
	}
	return message.(iot_device.CreateIoTDeviceRequest), nil
}

func (s *IoTDeviceMQTT) getMessageUpdate(message any) (iot_device.UpdateIoTDeviceRequest, error) {
	if _, ok := message.(iot_device.UpdateIoTDeviceRequest); !ok {
		return iot_device.UpdateIoTDeviceRequest{}, ErrMessageNil
	}
	return message.(iot_device.UpdateIoTDeviceRequest), nil
}
