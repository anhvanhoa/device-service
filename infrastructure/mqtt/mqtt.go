package mqtt_service

import (
	"context"
	iot_device_mqtt "device-service/infrastructure/mqtt/iot_device"
	"device-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
)

type MqttService interface {
	RunIoTDevice()
}

type MqttServiceImpl struct {
	iotDeviceMQTT iot_device_mqtt.IoTDeviceMQTTService
	mq            mq.MQI
	log           *log.LogGRPCImpl
}

func NewMqttService(repo repo.Repositories, helper utils.Helper, mq mq.MQI, log *log.LogGRPCImpl) MqttService {
	iotDeviceMQTT := iot_device_mqtt.NewIoTDeviceMQTT(repo, helper, log)
	return &MqttServiceImpl{iotDeviceMQTT: iotDeviceMQTT, mq: mq, log: log}
}

func (s *MqttServiceImpl) RunIoTDevice() {
	s.mq.Subscribe("iot_device/register", byte(0), func(message any) error {
		s.log.Info("RunIoTDevice subscribe iot_device/register")
		return s.iotDeviceMQTT.RegisterDevice(context.Background(), message)
	})
	s.mq.Subscribe("iot_device/update", byte(0), func(message any) error {
		s.log.Info("RunIoTDevice subscribe iot_device/update")
		return s.iotDeviceMQTT.UpdateDevice(context.Background(), message)
	})
}
