package iot_device_mqtt

import (
	"context"
	"device-service/domain/usecase/iot_device"
	"device-service/domain/usecase/sensor_data"
	"device-service/infrastructure/repo"
	"encoding/json"
	"fmt"

	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
)

type IoTDeviceMQTTService interface {
	RegisterDevice(ctx context.Context, msg any) error
	UpdateDevice(ctx context.Context, msg any) error
	ProcessSensorData(ctx context.Context, msg any) error
	ControlSensor(ctx context.Context, msg any) error
}

type IoTDeviceMQTT struct {
	iotDeviceUsecase  iot_device.IoTDeviceUsecase
	sensorDataUsecase sensor_data.SensorDataUsecase
	log               *log.LogGRPCImpl
}

func NewIoTDeviceMQTT(repo repo.Repositories, helper utils.Helper, mq mq.MQI, log *log.LogGRPCImpl) IoTDeviceMQTTService {
	return &IoTDeviceMQTT{
		iotDeviceUsecase:  iot_device.NewIoTDeviceUsecase(repo.IoTDevice(), helper, mq),
		sensorDataUsecase: sensor_data.NewSensorDataUsecase(repo.SensorData(), helper),
		log:               log,
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
	var jsonData iot_device.CreateIoTDeviceRequest
	err := json.Unmarshal(message.([]byte), &jsonData)
	if err != nil {
		fmt.Println("err", err)
		return iot_device.CreateIoTDeviceRequest{}, ErrMessageNil
	}
	return jsonData, nil
}

func (s *IoTDeviceMQTT) getMessageUpdate(message any) (iot_device.UpdateIoTDeviceRequest, error) {
	var jsonData iot_device.UpdateIoTDeviceRequest
	err := json.Unmarshal(message.([]byte), &jsonData)
	if err != nil {
		return iot_device.UpdateIoTDeviceRequest{}, ErrMessageNil
	}
	return jsonData, nil
}

func (s *IoTDeviceMQTT) ProcessSensorData(ctx context.Context, msg any) error {
	if err := s.checkMessage(msg); err != nil {
		s.log.Error("ProcessSensorData dữ liệu không hợp lệ")
		return err
	}

	// Parse JSON message
	jsonData, ok := msg.([]byte)
	if !ok {
		return fmt.Errorf("không thể convert message thành []byte")
	}

	req, err := sensor_data.ParseSensorDataFromJSON(jsonData)
	if err != nil {
		s.log.Error(fmt.Sprintf("ProcessSensorData parse JSON lỗi: %v", err))
		return err
	}

	// Xử lý dữ liệu cảm biến
	_, err = s.sensorDataUsecase.ProcessSensorData(ctx, req)
	if err != nil {
		s.log.Error(fmt.Sprintf("ProcessSensorData lỗi: %v", err))
		return err
	}

	s.log.Info("Đã xử lý dữ liệu cảm biến thành công")
	return nil
}

func (s *IoTDeviceMQTT) ControlSensor(ctx context.Context, msg any) error {
	if err := s.checkMessage(msg); err != nil {
		s.log.Error("ControlSensor dữ liệu không hợp lệ")
		return err
	}

	// Parse JSON message
	jsonData, ok := msg.([]byte)
	if !ok {
		return fmt.Errorf("không thể convert message thành []byte")
	}

	req, err := sensor_data.ParseControlSensorFromJSON(jsonData)
	if err != nil {
		s.log.Error(fmt.Sprintf("ControlSensor parse JSON lỗi: %v", err))
		return err
	}

	// Xử lý điều khiển cảm biến
	_, err = s.sensorDataUsecase.ControlSensor(ctx, req)
	if err != nil {
		s.log.Error(fmt.Sprintf("ControlSensor lỗi: %v", err))
		return err
	}

	s.log.Info("Đã xử lý điều khiển cảm biến thành công")
	return nil
}
