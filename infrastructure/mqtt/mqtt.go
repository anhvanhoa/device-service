package mqtt_service

import (
	"context"
	iot_device_mqtt "device-service/infrastructure/mqtt/iot_device"
	"device-service/infrastructure/repo"
	"encoding/json"
	"fmt"

	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
)

type MqttService interface {
	RunIoTDevice()
	RunSensorData()
}

type MqttServiceImpl struct {
	iotDeviceMQTT iot_device_mqtt.IoTDeviceMQTTService
	mq            mq.MQI
	log           *log.LogGRPCImpl
}

func NewMqttService(repo repo.Repositories, helper utils.Helper, mq mq.MQI, log *log.LogGRPCImpl) MqttService {
	iotDeviceMQTT := iot_device_mqtt.NewIoTDeviceMQTT(repo, helper, mq, log)
	return &MqttServiceImpl{iotDeviceMQTT: iotDeviceMQTT, mq: mq, log: log}
}

func (s *MqttServiceImpl) RunIoTDevice() {
	s.mq.Subscribe("iot_device/register", byte(0), func(message any) error {
		s.log.Info("RunIoTDevice subscribe iot_device/register")
		var jsonData map[string]any
		err := json.Unmarshal(message.([]byte), &jsonData)
		if err != nil {
			s.log.Error("RunIoTDevice subscribe iot_device/register")
			return err
		}
		fmt.Println("jsonData", jsonData)
		return s.iotDeviceMQTT.RegisterDevice(context.Background(), message)
	})
	s.mq.Subscribe("iot_device/update", byte(0), func(message any) error {
		s.log.Info("RunIoTDevice subscribe iot_device/update")
		return s.iotDeviceMQTT.UpdateDevice(context.Background(), message)
	})
}

func (s *MqttServiceImpl) RunSensorData() {
	s.mq.Subscribe("sensor/data", byte(0), func(message any) error {
		s.log.Info("RunSensorData subscribe sensor/data")
		return s.iotDeviceMQTT.ProcessSensorData(context.Background(), message)
	})
	s.mq.Subscribe("sensor/control", byte(0), func(message any) error {
		s.log.Info("RunSensorData subscribe sensor/control")
		return s.iotDeviceMQTT.ControlSensor(context.Background(), message)
	})
}
