package grpc_service

import (
	"device-service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/user_context"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	cacher cache.CacheI,
	deviceTypeServer proto_device_type.DeviceTypeServiceServer,
	iotDeviceServer proto_iot_device.IoTDeviceServiceServer,
	iotDeviceHistoryServer proto_iot_device_history.IoTDeviceHistoryServiceServer,
	sensorDataServer proto_sensor_data.SensorDataServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	middleware := grpc_service.NewMiddleware()
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_device_type.RegisterDeviceTypeServiceServer(server, deviceTypeServer)
			proto_iot_device.RegisterIoTDeviceServiceServer(server, iotDeviceServer)
			proto_iot_device_history.RegisterIoTDeviceHistoryServiceServer(server, iotDeviceHistoryServer)
			proto_sensor_data.RegisterSensorDataServiceServer(server, sensorDataServer)
		},
		middleware.AuthorizationInterceptor(
			env.SecretService,
			func(action string, resource string) bool {
				hasPermission, err := cacher.Get(resource + "." + action)
				if err != nil {
					return false
				}
				return hasPermission != nil && string(hasPermission) == "true"
			},
			func(id string) *user_context.UserContext {
				userData, err := cacher.Get(id)
				if err != nil || userData == nil {
					return nil
				}
				uCtx := user_context.NewUserContext()
				uCtx.FromBytes(userData)
				return uCtx
			},
		),
	)
}
