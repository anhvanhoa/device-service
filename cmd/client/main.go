package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Nhập trang (mặc định 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Nhập số bản ghi mỗi trang (mặc định 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type DeviceServiceClient struct {
	deviceTypeClient       proto_device_type.DeviceTypeServiceClient
	iotDeviceClient        proto_iot_device.IoTDeviceServiceClient
	iotDeviceHistoryClient proto_iot_device_history.IoTDeviceHistoryServiceClient
	sensorDataClient       proto_sensor_data.SensorDataServiceClient
	conn                   *grpc.ClientConn
}

func NewDeviceServiceClient(address string) (*DeviceServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &DeviceServiceClient{
		deviceTypeClient:       proto_device_type.NewDeviceTypeServiceClient(conn),
		iotDeviceClient:        proto_iot_device.NewIoTDeviceServiceClient(conn),
		iotDeviceHistoryClient: proto_iot_device_history.NewIoTDeviceHistoryServiceClient(conn),
		sensorDataClient:       proto_sensor_data.NewSensorDataServiceClient(conn),
		conn:                   conn,
	}, nil
}

func (c *DeviceServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Device Service Tests ==================

// Device Type Service Tests
func (c *DeviceServiceClient) TestCreateDeviceType() {
	fmt.Println("\n=== Kiểm thử Tạo Loại Thiết bị ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập mã loại thiết bị: ")
	typeCode, _ := reader.ReadString('\n')
	typeCode = cleanInput(typeCode)

	fmt.Print("Nhập mô tả: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.deviceTypeClient.CreateDeviceType(ctx, &proto_device_type.CreateDeviceTypeRequest{
		TypeCode:    typeCode,
		Description: description,
	})
	if err != nil {
		fmt.Printf("Error calling CreateDeviceType: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo loại thiết bị:\n")
	fmt.Printf("ID: %s\n", resp.Id)
	fmt.Printf("Mã loại: %s\n", resp.TypeCode)
	fmt.Printf("Mô tả: %s\n", resp.Description)
	if resp.CreatedAt != nil {
		fmt.Printf("Ngày tạo: %s\n", resp.CreatedAt.AsTime().Format("2006-01-02 15:04:05"))
	}
}

func (c *DeviceServiceClient) TestGetDeviceType() {
	fmt.Println("\n=== Kiểm thử Lấy Loại Thiết bị ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID loại thiết bị: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.deviceTypeClient.GetDeviceType(ctx, &proto_device_type.GetDeviceTypeRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetDeviceType: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy loại thiết bị:\n")
	if resp.DeviceType != nil {
		fmt.Printf("ID: %s\n", resp.DeviceType.Id)
		fmt.Printf("Mã loại: %s\n", resp.DeviceType.TypeCode)
		fmt.Printf("Mô tả: %s\n", resp.DeviceType.Description)
		if resp.DeviceType.CreatedAt != nil {
			fmt.Printf("Ngày tạo: %s\n", resp.DeviceType.CreatedAt.AsTime().Format("2006-01-02 15:04:05"))
		}
		if resp.DeviceType.UpdatedAt != nil {
			fmt.Printf("Ngày cập nhật: %s\n", resp.DeviceType.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"))
		}
	}
}

func (c *DeviceServiceClient) TestListDeviceTypes() {
	fmt.Println("\n=== Kiểm thử Liệt kê Loại Thiết bị ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	fmt.Print("Nhập từ khóa tìm kiếm (để trống để bỏ qua): ")
	search, _ := reader.ReadString('\n')
	search = cleanInput(search)

	filters := &proto_device_type.DeviceTypeFilters{}
	if search != "" {
		filters.Search = search
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.deviceTypeClient.ListDeviceType(ctx, &proto_device_type.ListDeviceTypeRequest{
		Filters: filters,
		Pagination: &proto_common.PaginationRequest{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListDeviceType: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê loại thiết bị:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách loại thiết bị:\n")
	for i, deviceType := range resp.Data {
		fmt.Printf("  [%d] ID: %s, Mã: %s, Mô tả: %s\n",
			i+1, deviceType.Id, deviceType.TypeCode, deviceType.Description)
	}
}

func (c *DeviceServiceClient) TestUpdateDeviceType() {
	fmt.Println("\n=== Kiểm thử Cập nhật Loại Thiết bị ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID loại thiết bị cần cập nhật: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập mã loại thiết bị mới: ")
	typeCode, _ := reader.ReadString('\n')
	typeCode = cleanInput(typeCode)

	fmt.Print("Nhập mô tả mới: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.deviceTypeClient.UpdateDeviceType(ctx, &proto_device_type.UpdateDeviceTypeRequest{
		Id:          id,
		TypeCode:    typeCode,
		Description: description,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateDeviceType: %v\n", err)
		return
	}

	fmt.Printf("Kết quả cập nhật loại thiết bị:\n")
	if resp.DeviceType != nil {
		fmt.Printf("ID: %s\n", resp.DeviceType.Id)
		fmt.Printf("Mã loại: %s\n", resp.DeviceType.TypeCode)
		fmt.Printf("Mô tả: %s\n", resp.DeviceType.Description)
		if resp.DeviceType.UpdatedAt != nil {
			fmt.Printf("Ngày cập nhật: %s\n", resp.DeviceType.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"))
		}
	}
}

func (c *DeviceServiceClient) TestDeleteDeviceType() {
	fmt.Println("\n=== Kiểm thử Xóa Loại Thiết bị ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID loại thiết bị cần xóa: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.deviceTypeClient.DeleteDeviceType(ctx, &proto_device_type.DeleteDeviceTypeRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteDeviceType: %v\n", err)
		return
	}

	fmt.Printf("Kết quả xóa loại thiết bị:\n")
	fmt.Printf("Thành công: %t\n", resp.Success)
}

// IoT Device Service Tests
func (c *DeviceServiceClient) TestCreateIotDevice() {
	fmt.Println("\n=== Kiểm thử Tạo Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập tên thiết bị: ")
	deviceName, _ := reader.ReadString('\n')
	deviceName = cleanInput(deviceName)

	fmt.Print("Nhập ngày cài đặt (YYYY-MM-DD): ")
	installationDateStr, _ := reader.ReadString('\n')
	installationDateStr = cleanInput(installationDateStr)

	var installationDate *timestamppb.Timestamp
	if installationDateStr != "" {
		fmt.Printf("installationDateStr: %s\n", installationDateStr)
		if t, err := time.Parse(time.DateOnly, installationDateStr); err == nil {
			installationDate = timestamppb.New(t)
		} else {
			fmt.Printf("Error parsing installation date: %v\n", err)
			return
		}
	}

	fmt.Print("Nhập ID loại thiết bị: ")
	deviceTypeId, _ := reader.ReadString('\n')
	deviceTypeId = cleanInput(deviceTypeId)

	fmt.Print("Nhập model: ")
	model, _ := reader.ReadString('\n')
	model = cleanInput(model)

	fmt.Print("Nhập địa chỉ MAC: ")
	macAddress, _ := reader.ReadString('\n')
	macAddress = cleanInput(macAddress)

	fmt.Print("Nhập địa chỉ IP: ")
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = cleanInput(ipAddress)

	fmt.Print("Nhập ID nhà kính: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	fmt.Print("Nhập ID vùng trồng: ")
	growingZoneId, _ := reader.ReadString('\n')
	growingZoneId = cleanInput(growingZoneId)

	fmt.Print("Nhập setting mặc định (json): ")
	defaultConfigStr, _ := reader.ReadString('\n')
	defaultConfigStr = cleanInput(defaultConfigStr)

	var defaultConfig *structpb.Struct
	if defaultConfigStr != "" {
		if err := json.Unmarshal([]byte(defaultConfigStr), &defaultConfig); err != nil {
			fmt.Printf("Error parsing default config: %v\n", err)
			return
		}
	}

	fmt.Print("Nhập mức pin (0-100): ")
	batteryLevelStr, _ := reader.ReadString('\n')
	batteryLevelStr = cleanInput(batteryLevelStr)
	batteryLevel := int32(100)
	if batteryLevelStr != "" {
		if b, err := strconv.Atoi(batteryLevelStr); err == nil {
			batteryLevel = int32(b)
		}
	}

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_iot_device.CreateIoTDeviceRequest{
		DeviceName:       deviceName,
		DeviceTypeId:     deviceTypeId,
		Model:            model,
		MacAddress:       macAddress,
		IpAddress:        ipAddress,
		GreenhouseId:     greenhouseId,
		GrowingZoneId:    growingZoneId,
		InstallationDate: installationDate,
		DefaultConfig:    defaultConfig,
		BatteryLevel:     batteryLevel,
		CreatedBy:        createdBy,
	}
	fmt.Printf("Request: %+v\n", req)

	resp, err := c.iotDeviceClient.CreateIoTDevice(ctx, req)
	if err != nil {
		fmt.Printf("Error calling CreateIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo thiết bị IoT:\n")
	fmt.Printf("Tên thiết bị: %s\n", resp.DeviceName)
	fmt.Printf("Model: %s\n", resp.Model)
	fmt.Printf("Trạng thái: %s\n", resp.Status)
	if resp.CreatedAt != nil {
		fmt.Printf("Ngày tạo: %s\n", resp.CreatedAt.AsTime().Format("2006-01-02 15:04:05"))
	}
}

func (c *DeviceServiceClient) TestGetIotDevice() {
	fmt.Println("\n=== Kiểm thử Lấy Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị IoT: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceClient.GetIoTDevice(ctx, &proto_iot_device.GetIoTDeviceRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy thiết bị IoT:\n")
	if resp.Device != nil {
		fmt.Printf("ID: %s\n", resp.Device.Id)
		fmt.Printf("Tên thiết bị: %s\n", resp.Device.DeviceName)
		fmt.Printf("ID loại thiết bị: %s\n", resp.Device.DeviceTypeId)
		fmt.Printf("Model: %s\n", resp.Device.Model)
		fmt.Printf("Địa chỉ MAC: %s\n", resp.Device.MacAddress)
		fmt.Printf("Địa chỉ IP: %s\n", resp.Device.IpAddress)
		fmt.Printf("ID nhà kính: %s\n", resp.Device.GreenhouseId)
		fmt.Printf("ID vùng trồng: %s\n", resp.Device.GrowingZoneId)
		if resp.Device.InstallationDate != nil {
			fmt.Printf("Ngày cài đặt: %s\n", resp.Device.InstallationDate.AsTime().Format("2006-01-02 15:04:05"))
		}
		if resp.Device.LastMaintenanceDate != nil {
			fmt.Printf("Ngày bảo trì cuối: %s\n", resp.Device.LastMaintenanceDate.AsTime().Format("2006-01-02 15:04:05"))
		}
		fmt.Printf("Mức pin: %d%%\n", resp.Device.BatteryLevel)
		fmt.Printf("Trạng thái: %s\n", resp.Device.Status)
		fmt.Printf("Người tạo: %s\n", resp.Device.CreatedBy)
		if resp.Device.CreatedAt != nil {
			fmt.Printf("Ngày tạo: %s\n", resp.Device.CreatedAt.AsTime().Format("2006-01-02 15:04:05"))
		}
	}
}

func (c *DeviceServiceClient) TestListIotDevices() {
	fmt.Println("\n=== Kiểm thử Liệt kê Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	fmt.Print("Nhập ID loại thiết bị (để trống để bỏ qua): ")
	deviceTypeId, _ := reader.ReadString('\n')
	deviceTypeId = cleanInput(deviceTypeId)

	fmt.Print("Nhập trạng thái (active/inactive/maintenance/error/offline, để trống để bỏ qua): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	fmt.Print("Nhập ID nhà kính (để trống để bỏ qua): ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	fmt.Print("Nhập ID vùng trồng (để trống để bỏ qua): ")
	growingZoneId, _ := reader.ReadString('\n')
	growingZoneId = cleanInput(growingZoneId)

	fmt.Print("Nhập từ khóa tìm kiếm (để trống để bỏ qua): ")
	search, _ := reader.ReadString('\n')
	search = cleanInput(search)

	filters := &proto_iot_device.IoTDeviceFilters{}
	if deviceTypeId != "" {
		filters.DeviceTypeId = deviceTypeId
	}
	if status != "" {
		filters.Status = status
	}
	if greenhouseId != "" {
		filters.GreenhouseId = greenhouseId
	}
	if growingZoneId != "" {
		filters.GrowingZoneId = growingZoneId
	}
	if search != "" {
		filters.Search = search
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceClient.ListIoTDevice(ctx, &proto_iot_device.ListIoTDeviceRequest{
		Filters: filters,
		Pagination: &proto_common.PaginationRequest{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê thiết bị IoT:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách thiết bị IoT:\n")
	for i, device := range resp.Data {
		fmt.Printf("  [%d] ID: %s, Tên: %s, Model: %s, Trạng thái: %s\n",
			i+1, device.Id, device.DeviceName, device.Model, device.Status)
	}
}

func (c *DeviceServiceClient) TestUpdateIotDevice() {
	fmt.Println("\n=== Kiểm thử Cập nhật Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị IoT cần cập nhật: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập tên thiết bị mới: ")
	deviceName, _ := reader.ReadString('\n')
	deviceName = cleanInput(deviceName)

	fmt.Print("Nhập ID loại thiết bị mới: ")
	deviceTypeId, _ := reader.ReadString('\n')
	deviceTypeId = cleanInput(deviceTypeId)

	fmt.Print("Nhập model mới: ")
	model, _ := reader.ReadString('\n')
	model = cleanInput(model)

	fmt.Print("Nhập địa chỉ MAC mới: ")
	macAddress, _ := reader.ReadString('\n')
	macAddress = cleanInput(macAddress)

	fmt.Print("Nhập địa chỉ IP mới: ")
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = cleanInput(ipAddress)

	fmt.Print("Nhập ID nhà kính mới: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	fmt.Print("Nhập ID vùng trồng mới: ")
	growingZoneId, _ := reader.ReadString('\n')
	growingZoneId = cleanInput(growingZoneId)

	fmt.Print("Nhập ngày cài đặt mới (YYYY-MM-DD): ")
	installationDateStr, _ := reader.ReadString('\n')
	installationDateStr = cleanInput(installationDateStr)

	// Parse installation date
	var installationDate *timestamppb.Timestamp
	if installationDateStr != "" {
		if t, err := time.Parse(time.DateOnly, installationDateStr); err == nil {
			installationDate = timestamppb.New(t)
		} else {
			fmt.Printf("Error parsing installation date: %v\n", err)
			return
		}
	}

	fmt.Print("Nhập ngày bảo trì cuối mới (YYYY-MM-DD): ")
	lastMaintenanceDateStr, _ := reader.ReadString('\n')
	lastMaintenanceDateStr = cleanInput(lastMaintenanceDateStr)

	// Parse last maintenance date
	var lastMaintenanceDate *timestamppb.Timestamp
	if lastMaintenanceDateStr != "" {
		if t, err := time.Parse(time.DateOnly, lastMaintenanceDateStr); err == nil {
			lastMaintenanceDate = timestamppb.New(t)
		} else {
			fmt.Printf("Error parsing last maintenance date: %v\n", err)
			return
		}
	}

	fmt.Print("Nhập mức pin mới (0-100): ")
	batteryLevelStr, _ := reader.ReadString('\n')
	batteryLevelStr = cleanInput(batteryLevelStr)
	batteryLevel := int32(100)
	if batteryLevelStr != "" {
		if b, err := strconv.Atoi(batteryLevelStr); err == nil {
			batteryLevel = int32(b)
		}
	}

	fmt.Print("Nhập trạng thái mới (active/inactive/maintenance/error/offline): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_iot_device.UpdateIoTDeviceRequest{
		Id:                  id,
		DeviceName:          deviceName,
		DeviceTypeId:        deviceTypeId,
		Model:               model,
		MacAddress:          macAddress,
		IpAddress:           ipAddress,
		GreenhouseId:        greenhouseId,
		GrowingZoneId:       growingZoneId,
		InstallationDate:    installationDate,
		LastMaintenanceDate: lastMaintenanceDate,
		BatteryLevel:        batteryLevel,
		Status:              status,
	}

	fmt.Printf("Request: %+v\n", req)

	resp, err := c.iotDeviceClient.UpdateIoTDevice(ctx, req)
	if err != nil {
		fmt.Printf("Error calling UpdateIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả cập nhật thiết bị IoT:\n")
	if resp.Device != nil {
		fmt.Printf("ID: %s\n", resp.Device.Id)
		fmt.Printf("Tên thiết bị: %s\n", resp.Device.DeviceName)
		fmt.Printf("Model: %s\n", resp.Device.Model)
		fmt.Printf("Trạng thái: %s\n", resp.Device.Status)
		if resp.Device.UpdatedAt != nil {
			fmt.Printf("Ngày cập nhật: %s\n", resp.Device.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"))
		}
	}
}

func (c *DeviceServiceClient) TestDeleteIotDevice() {
	fmt.Println("\n=== Kiểm thử Xóa Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị IoT cần xóa: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceClient.DeleteIoTDevice(ctx, &proto_iot_device.DeleteIoTDeviceRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả xóa thiết bị IoT:\n")
	fmt.Printf("Thành công: %t\n", resp.Success)
}

func (c *DeviceServiceClient) TestControlIotDevice() {
	fmt.Println("\n=== Kiểm thử Điều khiển Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị IoT cần điều khiển: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập hành động (on/off/toggle/reset): ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceClient.ControlIoTDevice(ctx, &proto_iot_device.ControlIoTDeviceRequest{
		Id:     id,
		Action: action,
	})
	if err != nil {
		fmt.Printf("Error calling ControlIoTDevice: %v\n", err)
		return
	}

	fmt.Printf("Kết quả điều khiển thiết bị IoT:\n")
	fmt.Printf("ID: %s\n", resp.Id)
	fmt.Printf("Trạng thái: %s\n", resp.Status)
	fmt.Printf("Hành động: %s\n", resp.Action)
	fmt.Printf("Thông báo: %s\n", resp.Message)
}

// IoT Device History Service Tests
func (c *DeviceServiceClient) TestCreateIotDeviceHistory() {
	fmt.Println("\n=== Kiểm thử Tạo Lịch sử Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị: ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập hành động (install/update_config/firmware_update/relocate/maintenance/deactivate/reactivate): ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	fmt.Print("Nhập người thực hiện: ")
	performedBy, _ := reader.ReadString('\n')
	performedBy = cleanInput(performedBy)

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	// Create simple old_value and new_value structs
	oldValue, _ := structpb.NewStruct(map[string]any{
		"status":        "inactive",
		"battery_level": 50,
	})

	newValue, _ := structpb.NewStruct(map[string]any{
		"status":        "active",
		"battery_level": 80,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceHistoryClient.CreateIoTDeviceHistory(ctx, &proto_iot_device_history.CreateIoTDeviceHistoryRequest{
		DeviceId:    deviceId,
		Action:      action,
		OldValue:    oldValue,
		NewValue:    newValue,
		PerformedBy: performedBy,
		Notes:       notes,
	})
	if err != nil {
		fmt.Printf("Error calling CreateIoTDeviceHistory: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo lịch sử thiết bị IoT:\n")
	fmt.Printf("ID thiết bị: %s\n", resp.DeviceId)
	fmt.Printf("Hành động: %s\n", resp.Action)
	fmt.Printf("Ngày thực hiện: %s\n", resp.ActionDate)
	fmt.Printf("Người thực hiện: %s\n", resp.PerformedBy)
	fmt.Printf("Ghi chú: %s\n", resp.Notes)
}

func (c *DeviceServiceClient) TestGetIotDeviceHistory() {
	fmt.Println("\n=== Kiểm thử Lấy Lịch sử Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID lịch sử thiết bị IoT: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceHistoryClient.GetIoTDeviceHistory(ctx, &proto_iot_device_history.GetIoTDeviceHistoryRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetIoTDeviceHistory: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy lịch sử thiết bị IoT:\n")
	if resp.History != nil {
		fmt.Printf("ID: %s\n", resp.History.Id)
		fmt.Printf("ID thiết bị: %s\n", resp.History.DeviceId)
		fmt.Printf("Hành động: %s\n", resp.History.Action)
		fmt.Printf("Ngày thực hiện: %s\n", resp.History.ActionDate)
		fmt.Printf("Người thực hiện: %s\n", resp.History.PerformedBy)
		fmt.Printf("Ghi chú: %s\n", resp.History.Notes)
		fmt.Printf("Ngày tạo: %s\n", resp.History.CreatedAt.AsTime().Format(time.DateOnly))
	}
}

func (c *DeviceServiceClient) TestListIotDeviceHistories() {
	fmt.Println("\n=== Kiểm thử Liệt kê Lịch sử Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	fmt.Print("Nhập ID thiết bị (để trống để bỏ qua): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập hành động (install/update_config/firmware_update/relocate/maintenance/deactivate/reactivate, để trống để bỏ qua): ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	fmt.Print("Nhập người thực hiện (để trống để bỏ qua): ")
	performedBy, _ := reader.ReadString('\n')
	performedBy = cleanInput(performedBy)

	fmt.Print("Nhập ngày bắt đầu (YYYY-MM-DD, để trống để bỏ qua): ")
	startDate, _ := reader.ReadString('\n')
	startDate = cleanInput(startDate)

	fmt.Print("Nhập ngày kết thúc (YYYY-MM-DD, để trống để bỏ qua): ")
	endDate, _ := reader.ReadString('\n')
	endDate = cleanInput(endDate)

	filters := &proto_iot_device_history.IoTDeviceHistoryFilters{}
	if deviceId != "" {
		filters.DeviceId = deviceId
	}
	if action != "" {
		filters.Action = action
	}
	if performedBy != "" {
		filters.PerformedBy = performedBy
	}
	if startDate != "" {
		t, err := time.Parse(time.DateOnly, startDate)
		if err != nil {
			fmt.Printf("Error parsing start date: %v\n", err)
			return
		}
		filters.StartDate = timestamppb.New(t)
	}
	if endDate != "" {
		t, err := time.Parse(time.DateOnly, endDate)
		if err != nil {
			fmt.Printf("Error parsing end date: %v\n", err)
			return
		}
		filters.EndDate = timestamppb.New(t)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceHistoryClient.ListIoTDeviceHistory(ctx, &proto_iot_device_history.ListIoTDeviceHistoryRequest{
		Filters: filters,
		Pagination: &proto_common.PaginationRequest{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListIoTDeviceHistory: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê lịch sử thiết bị IoT:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách lịch sử thiết bị IoT:\n")
	for i, history := range resp.Data {
		fmt.Printf("  [%d] ID: %s, Thiết bị: %s, Hành động: %s, Ngày: %s, Ngày tạo: %s\n",
			i+1, history.Id, history.DeviceId, history.Action, history.ActionDate, history.CreatedAt.AsTime().Format(time.DateOnly))
	}
}

func (c *DeviceServiceClient) TestDeleteIotDeviceHistory() {
	fmt.Println("\n=== Kiểm thử Xóa Lịch sử Thiết bị IoT ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID lịch sử thiết bị IoT cần xóa: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.iotDeviceHistoryClient.DeleteIoTDeviceHistory(ctx, &proto_iot_device_history.DeleteIoTDeviceHistoryRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteIoTDeviceHistory: %v\n", err)
		return
	}

	fmt.Printf("Kết quả xóa lịch sử thiết bị IoT:\n")
	fmt.Printf("Thành công: %t\n", resp.Success)
}

// Sensor Data Service Tests
func (c *DeviceServiceClient) TestCreateSensorData() {
	fmt.Println("\n=== Kiểm thử Tạo Dữ liệu Cảm biến ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị: ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại cảm biến (temperature/humidity/ph/light_intensity/soil_moisture/co2/water_level): ")
	sensorType, _ := reader.ReadString('\n')
	sensorType = cleanInput(sensorType)

	fmt.Print("Nhập giá trị: ")
	valueStr, _ := reader.ReadString('\n')
	valueStr = cleanInput(valueStr)
	value := 0.0
	if valueStr != "" {
		if v, err := strconv.ParseFloat(valueStr, 64); err == nil {
			value = v
		}
	}

	fmt.Print("Nhập đơn vị (C/F/%/lux/ppm/pH): ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Có cảnh báo (true/false): ")
	isAlertStr, _ := reader.ReadString('\n')
	isAlertStr = cleanInput(isAlertStr)
	isAlert := false
	if isAlertStr == "true" {
		isAlert = true
	}

	fmt.Print("Nhập điểm chất lượng (0.0-1.0): ")
	qualityScoreStr, _ := reader.ReadString('\n')
	qualityScoreStr = cleanInput(qualityScoreStr)
	qualityScore := 1.0
	if qualityScoreStr != "" {
		if q, err := strconv.ParseFloat(qualityScoreStr, 64); err == nil {
			qualityScore = q
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.sensorDataClient.CreateSensorData(ctx, &proto_sensor_data.CreateSensorDataRequest{
		DeviceId:     deviceId,
		SensorType:   sensorType,
		Value:        value,
		Unit:         unit,
		IsAlert:      isAlert,
		QualityScore: qualityScore,
	})
	if err != nil {
		fmt.Printf("Error calling CreateSensorData: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo dữ liệu cảm biến:\n")
	fmt.Printf("ID thiết bị: %s\n", resp.DeviceId)
	fmt.Printf("Loại cảm biến: %s\n", resp.SensorType)
	fmt.Printf("Giá trị: %.2f %s\n", resp.Value, resp.Unit)
	fmt.Printf("Có cảnh báo: %t\n", resp.IsAlert)
	fmt.Printf("Điểm chất lượng: %.2f\n", resp.QualityScore)
	fmt.Printf("Ngày ghi nhận: %s\n", resp.RecordedAt.AsTime().Format(time.DateOnly))
	fmt.Printf("Ngày tạo: %s\n", resp.CreatedAt.AsTime().Format(time.DateOnly))
}

func (c *DeviceServiceClient) TestGetSensorData() {
	fmt.Println("\n=== Kiểm thử Lấy Dữ liệu Cảm biến ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID dữ liệu cảm biến: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.sensorDataClient.GetSensorData(ctx, &proto_sensor_data.GetSensorDataRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetSensorData: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy dữ liệu cảm biến:\n")
	if resp.SensorData != nil {
		fmt.Printf("ID: %s\n", resp.SensorData.Id)
		fmt.Printf("ID thiết bị: %s\n", resp.SensorData.DeviceId)
		fmt.Printf("Loại cảm biến: %s\n", resp.SensorData.SensorType)
		fmt.Printf("Giá trị: %.2f %s\n", resp.SensorData.Value, resp.SensorData.Unit)
		fmt.Printf("Ngày ghi nhận: %s\n", resp.SensorData.RecordedAt.AsTime().Format(time.DateOnly))
		fmt.Printf("Có cảnh báo: %t\n", resp.SensorData.IsAlert)
		fmt.Printf("Điểm chất lượng: %.2f\n", resp.SensorData.QualityScore)
		fmt.Printf("Ngày tạo: %s\n", resp.SensorData.CreatedAt.AsTime().Format(time.DateOnly))
	}
}

func (c *DeviceServiceClient) TestListSensorData() {
	fmt.Println("\n=== Kiểm thử Liệt kê Dữ liệu Cảm biến ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	fmt.Print("Nhập ID thiết bị (để trống để bỏ qua): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại cảm biến (temperature/humidity/ph/light_intensity/soil_moisture/co2/water_level, để trống để bỏ qua): ")
	sensorType, _ := reader.ReadString('\n')
	sensorType = cleanInput(sensorType)

	fmt.Print("Nhập đơn vị (C/F/%/lux/ppm/pH, để trống để bỏ qua): ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Có cảnh báo (true/false, để trống để bỏ qua): ")
	isAlertStr, _ := reader.ReadString('\n')
	isAlertStr = cleanInput(isAlertStr)
	var isAlert *bool
	if isAlertStr != "" {
		alert := isAlertStr == "true"
		isAlert = &alert
	}

	fmt.Print("Nhập điểm chất lượng tối thiểu (0.0-1.0, để trống để bỏ qua): ")
	minQualityStr, _ := reader.ReadString('\n')
	minQualityStr = cleanInput(minQualityStr)
	var minQuality float64
	if minQualityStr != "" {
		if q, err := strconv.ParseFloat(minQualityStr, 64); err == nil {
			minQuality = q
		}
	}

	fmt.Print("Nhập điểm chất lượng tối đa (0.0-1.0, để trống để bỏ qua): ")
	maxQualityStr, _ := reader.ReadString('\n')
	maxQualityStr = cleanInput(maxQualityStr)
	var maxQuality float64
	if maxQualityStr != "" {
		if q, err := strconv.ParseFloat(maxQualityStr, 64); err == nil {
			maxQuality = q
		}
	}

	fmt.Print("Nhập ngày bắt đầu (YYYY-MM-DD, để trống để bỏ qua): ")
	startDate, _ := reader.ReadString('\n')
	startDate = cleanInput(startDate)

	fmt.Print("Nhập ngày kết thúc (YYYY-MM-DD, để trống để bỏ qua): ")
	endDate, _ := reader.ReadString('\n')
	endDate = cleanInput(endDate)

	filters := &proto_sensor_data.SensorDataFilters{}
	if deviceId != "" {
		filters.DeviceId = deviceId
	}
	if sensorType != "" {
		filters.SensorType = sensorType
	}
	if unit != "" {
		filters.Unit = unit
	}
	if isAlert != nil {
		filters.IsAlert = *isAlert
	}
	if minQuality > 0 {
		filters.MinQualityScore = minQuality
	}
	if maxQuality > 0 {
		filters.MaxQualityScore = maxQuality
	}
	if startDate != "" {
		t, err := time.Parse(time.DateOnly, startDate)
		if err != nil {
			fmt.Printf("Error parsing start date: %v\n", err)
			return
		}
		filters.StartDate = timestamppb.New(t)
	}
	if endDate != "" {
		t, err := time.Parse(time.DateOnly, endDate)
		if err != nil {
			fmt.Printf("Error parsing end date: %v\n", err)
			return
		}
		filters.EndDate = timestamppb.New(t)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.sensorDataClient.ListSensorData(ctx, &proto_sensor_data.ListSensorDataRequest{
		Filters: filters,
		Pagination: &proto_common.PaginationRequest{
			Page:     offset,
			PageSize: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListSensorData: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê dữ liệu cảm biến:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách dữ liệu cảm biến:\n")
	for i, sensorData := range resp.Data {
		fmt.Printf("  [%d] ID: %s, Thiết bị: %s, Loại: %s, Giá trị: %.2f %s, Cảnh báo: %t, Chất lượng: %.2f, Ngày ghi nhận: %s, Ngày tạo: %s\n",
			i+1, sensorData.Id, sensorData.DeviceId, sensorData.SensorType,
			sensorData.Value, sensorData.Unit, sensorData.IsAlert, sensorData.QualityScore,
			sensorData.RecordedAt.AsTime().Format(time.DateOnly),
			sensorData.CreatedAt.AsTime().Format(time.DateOnly))
	}
}

func (c *DeviceServiceClient) TestDeleteSensorData() {
	fmt.Println("\n=== Kiểm thử Xóa Dữ liệu Cảm biến ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID dữ liệu cảm biến cần xóa: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.sensorDataClient.DeleteSensorData(ctx, &proto_sensor_data.DeleteSensorDataRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteSensorData: %v\n", err)
		return
	}

	fmt.Printf("Kết quả xóa dữ liệu cảm biến:\n")
	fmt.Printf("Thành công: %t\n", resp.Success)
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC Device Service ===")
	fmt.Println("1. Dịch vụ Loại Thiết bị")
	fmt.Println("2. Dịch vụ Thiết bị IoT")
	fmt.Println("3. Dịch vụ Lịch sử Thiết bị IoT")
	fmt.Println("4. Dịch vụ Dữ liệu Cảm biến")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printDeviceTypeMenu() {
	fmt.Println("\n=== Dịch vụ Loại Thiết bị ===")
	fmt.Println("1. Tạo loại thiết bị")
	fmt.Println("2. Lấy loại thiết bị")
	fmt.Println("3. Liệt kê loại thiết bị")
	fmt.Println("4. Cập nhật loại thiết bị")
	fmt.Println("5. Xóa loại thiết bị")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printIotDeviceMenu() {
	fmt.Println("\n=== Dịch vụ Thiết bị IoT ===")
	fmt.Println("1. Tạo thiết bị IoT")
	fmt.Println("2. Lấy thiết bị IoT")
	fmt.Println("3. Liệt kê thiết bị IoT")
	fmt.Println("4. Cập nhật thiết bị IoT")
	fmt.Println("5. Xóa thiết bị IoT")
	fmt.Println("6. Điều khiển thiết bị IoT")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printIotDeviceHistoryMenu() {
	fmt.Println("\n=== Dịch vụ Lịch sử Thiết bị IoT ===")
	fmt.Println("1. Tạo lịch sử thiết bị IoT")
	fmt.Println("2. Lấy lịch sử thiết bị IoT")
	fmt.Println("3. Liệt kê lịch sử thiết bị IoT")
	fmt.Println("4. Xóa lịch sử thiết bị IoT")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printSensorDataMenu() {
	fmt.Println("\n=== Dịch vụ Dữ liệu Cảm biến ===")
	fmt.Println("1. Tạo dữ liệu cảm biến")
	fmt.Println("2. Lấy dữ liệu cảm biến")
	fmt.Println("3. Liệt kê dữ liệu cảm biến")
	fmt.Println("4. Xóa dữ liệu cảm biến")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewDeviceServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Kết nối thành công!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Dịch vụ Loại Thiết bị
			for {
				printDeviceTypeMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateDeviceType()
				case "2":
					client.TestGetDeviceType()
				case "3":
					client.TestListDeviceTypes()
				case "4":
					client.TestUpdateDeviceType()
				case "5":
					client.TestDeleteDeviceType()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Dịch vụ Thiết bị IoT
			for {
				printIotDeviceMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateIotDevice()
				case "2":
					client.TestGetIotDevice()
				case "3":
					client.TestListIotDevices()
				case "4":
					client.TestUpdateIotDevice()
				case "5":
					client.TestDeleteIotDevice()
				case "6":
					client.TestControlIotDevice()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "3":
			// Dịch vụ Lịch sử Thiết bị IoT
			for {
				printIotDeviceHistoryMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateIotDeviceHistory()
				case "2":
					client.TestGetIotDeviceHistory()
				case "3":
					client.TestListIotDeviceHistories()
				case "4":
					client.TestDeleteIotDeviceHistory()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "4":
			// Dịch vụ Dữ liệu Cảm biến
			for {
				printSensorDataMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateSensorData()
				case "2":
					client.TestGetSensorData()
				case "3":
					client.TestListSensorData()
				case "4":
					client.TestDeleteSensorData()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}
