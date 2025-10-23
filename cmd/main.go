package main

import (
	"context"
	"device-service/bootstrap"
	"device-service/infrastructure/grpc_client"
	"device-service/infrastructure/grpc_service"
	device_type_service "device-service/infrastructure/grpc_service/device_type"
	iot_device_service "device-service/infrastructure/grpc_service/iot_device"
	iot_device_history_service "device-service/infrastructure/grpc_service/iot_device_history"
	sensor_data_service "device-service/infrastructure/grpc_service/sensor_data"
	mqtt_service "device-service/infrastructure/mqtt"

	"github.com/anhvanhoa/service-core/domain/discovery"
	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

	deviceTypeServer := device_type_service.NewDeviceTypeService(app.Repo.DeviceType(), app.Helper)
	iotDeviceServer := iot_device_service.NewIoTDeviceService(app.Repo.IoTDevice(), app.Helper, app.MQ)
	iotDeviceHistoryServer := iot_device_history_service.NewIoTDeviceHistoryService(app.Repo.IoTDeviceHistory(), app.Helper)
	sensorDataServer := sensor_data_service.NewSensorDataService(app.Repo.SensorData(), app.Helper)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log, app.Cache,
		deviceTypeServer,
		iotDeviceServer,
		iotDeviceHistoryServer,
		sensorDataServer,
	)

	mqttService := mqtt_service.NewMqttService(app.Repo, app.Helper, app.MQ, log)
	mqttService.RunIoTDevice()
	mqttService.RunSensorData()

	ctx, cancel := context.WithCancel(context.Background())
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
