syntax = "proto3";

package device_type.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/device_type/v1;proto_device_type";

import "common/v1/common.proto";
import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

message DeviceType {
  string id = 1;
  string type_code = 2;
  string description = 3;
  google.protobuf.Timestamp created_at = 4;
  google.protobuf.Timestamp updated_at = 5;
}

message CreateDeviceTypeRequest {
  string type_code = 1 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string description = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 500
  ];
}

message CreateDeviceTypeResponse {
  string id = 1 [(buf.validate.field).string.min_len = 1];
  string type_code = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string description = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 500
  ];
  google.protobuf.Timestamp created_at = 4;
}

message GetDeviceTypeRequest {
  string id = 1 [(buf.validate.field).string.min_len = 1];
}

message GetDeviceTypeResponse {
  DeviceType device_type = 1;
}

message UpdateDeviceTypeRequest {
  string id = 1 [(buf.validate.field).string.min_len = 1];
  string type_code = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string description = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 500
  ];
}

message UpdateDeviceTypeResponse {
  DeviceType device_type = 1;
}

message DeleteDeviceTypeRequest {
  string id = 1 [(buf.validate.field).string.min_len = 1];
}

message DeleteDeviceTypeResponse {
  bool success = 1;
}

message ListDeviceTypeRequest {
  DeviceTypeFilters filters = 1;
  common.PaginationRequest pagination = 2;
}

message ListDeviceTypeResponse {
  repeated DeviceType data = 1;
  common.PaginationResponse pagination = 2;
}

message DeviceTypeFilters {
  string search = 1 [(buf.validate.field).string.max_len = 100];
}

service DeviceTypeService {
  rpc CreateDeviceType(CreateDeviceTypeRequest) returns (CreateDeviceTypeResponse);
  rpc GetDeviceType(GetDeviceTypeRequest) returns (GetDeviceTypeResponse);
  rpc UpdateDeviceType(UpdateDeviceTypeRequest) returns (UpdateDeviceTypeResponse);
  rpc DeleteDeviceType(DeleteDeviceTypeRequest) returns (DeleteDeviceTypeResponse);
  rpc ListDeviceType(ListDeviceTypeRequest) returns (ListDeviceTypeResponse);
}

syntax = "proto3";

package iot_device.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/iot_device/v1;proto_iot_device";

import "common/v1/common.proto";
import "google/protobuf/struct.proto";
import "buf/validate/validate.proto";
import "google/protobuf/timestamp.proto";

message IoTDevice {
  string id = 1;
  string device_name = 2;
  string device_type_id = 3;
  string model = 4;
  string mac_address = 5;
  string ip_address = 6;
  string greenhouse_id = 7;
  string growing_zone_id = 8;
  google.protobuf.Timestamp installation_date = 9;
  google.protobuf.Timestamp last_maintenance_date = 10;
  int32 battery_level = 11;
  string status = 12;
  google.protobuf.Struct configuration = 13;
  google.protobuf.Struct default_config = 14;
  string created_by = 15;
  google.protobuf.Timestamp created_at = 16;
  google.protobuf.Timestamp updated_at = 17;
}

message CreateIoTDeviceRequest {
  string device_name = 1 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string device_type_id = 2 [(buf.validate.field).string.min_len = 1];
  string model = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string mac_address = 4 [
    (buf.validate.field).string.min_len = 17,
    (buf.validate.field).string.max_len = 17
  ];
  string ip_address = 5 [
    (buf.validate.field).string.min_len = 7,
    (buf.validate.field).string.max_len = 45
  ];
  string greenhouse_id = 6 [(buf.validate.field).string.min_len = 1];
  string growing_zone_id = 7 [(buf.validate.field).string.min_len = 1];
  google.protobuf.Timestamp installation_date = 8 [(buf.validate.field).string.min_len = 1];
  int32 battery_level = 9 [
    (buf.validate.field).int32.gte = 0,
    (buf.validate.field).int32.lte = 100
  ];
  google.protobuf.Struct configuration = 10;
  google.protobuf.Struct default_config = 11;
  string created_by = 12 [(buf.validate.field).string.min_len = 1];
}

message CreateIoTDeviceResponse {
  string device_name = 1 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string device_type_id = 2 [(buf.validate.field).string.min_len = 1];
  string model = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string mac_address = 4 [
    (buf.validate.field).string.min_len = 17,
    (buf.validate.field).string.max_len = 17
  ];
  string ip_address = 5 [
    (buf.validate.field).string.min_len = 7,
    (buf.validate.field).string.max_len = 45
  ];
  string greenhouse_id = 6 [(buf.validate.field).string.min_len = 1];
  string growing_zone_id = 7 [(buf.validate.field).string.min_len = 1];
  google.protobuf.Timestamp installation_date = 8 [(buf.validate.field).string.min_len = 1];
  int32 battery_level = 9 [
    (buf.validate.field).int32.gte = 0,
    (buf.validate.field).int32.lte = 100
  ];
  string status = 10 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "error",
    (buf.validate.field).string.in = "offline"
  ];
  google.protobuf.Struct configuration = 11;
  google.protobuf.Struct default_config = 12;
  string created_by = 13 [(buf.validate.field).string.min_len = 1];
  google.protobuf.Timestamp created_at = 14 [(buf.validate.field).string.min_len = 1];
}

message GetIoTDeviceRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message GetIoTDeviceResponse {
  IoTDevice device = 1;
}

message UpdateIoTDeviceRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
  string device_name = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string device_type_id = 3 [(buf.validate.field).string.min_len = 1];
  string model = 4 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string mac_address = 5 [
    (buf.validate.field).string.min_len = 17,
    (buf.validate.field).string.max_len = 17
  ];
  string ip_address = 6 [
    (buf.validate.field).string.min_len = 7,
    (buf.validate.field).string.max_len = 45
  ];
  string greenhouse_id = 7 [(buf.validate.field).string.min_len = 1];
  string growing_zone_id = 8 [(buf.validate.field).string.min_len = 1];
  google.protobuf.Timestamp installation_date = 9 [(buf.validate.field).string.min_len = 1];
  google.protobuf.Timestamp last_maintenance_date = 10 [(buf.validate.field).string.min_len = 1];
  int32 battery_level = 11 [
    (buf.validate.field).int32.gte = 0,
    (buf.validate.field).int32.lte = 100
  ];
  string status = 12 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "error",
    (buf.validate.field).string.in = "offline"
  ];
  google.protobuf.Struct configuration = 13;
  google.protobuf.Struct default_config = 14;
}

message UpdateIoTDeviceResponse {
  IoTDevice device = 1;
}

message DeleteIoTDeviceRequest {
  string id = 1 [(buf.validate.field).string.min_len = 1];
}

message DeleteIoTDeviceResponse {
  bool success = 1;
}

message ListIoTDeviceRequest {
  IoTDeviceFilters filters = 1;
  common.PaginationRequest pagination = 2;
}

message ListIoTDeviceResponse {
  repeated IoTDevice data = 1;
  common.PaginationResponse pagination = 2;
}

message IoTDeviceFilters {
  string device_type_id = 1 [(buf.validate.field).string.min_len = 1];
  string status = 2 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "error",
    (buf.validate.field).string.in = "offline"
  ];
  string greenhouse_id = 3 [(buf.validate.field).string.min_len = 1];
  string growing_zone_id = 4 [(buf.validate.field).string.min_len = 1];
  string search = 5 [(buf.validate.field).string.max_len = 100];
}

service IoTDeviceService {
  rpc CreateIoTDevice(CreateIoTDeviceRequest) returns (CreateIoTDeviceResponse);
  rpc GetIoTDevice(GetIoTDeviceRequest) returns (GetIoTDeviceResponse);
  rpc UpdateIoTDevice(UpdateIoTDeviceRequest) returns (UpdateIoTDeviceResponse);
  rpc DeleteIoTDevice(DeleteIoTDeviceRequest) returns (DeleteIoTDeviceResponse);
  rpc ListIoTDevice(ListIoTDeviceRequest) returns (ListIoTDeviceResponse);
}


syntax = "proto3";

package iot_device_history.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1;proto_iot_device_history";

import "common/v1/common.proto";
import "buf/validate/validate.proto";
import "google/protobuf/struct.proto";
import "google/protobuf/timestamp.proto";

message IoTDeviceHistory {
  string id = 1;
  string device_id = 2;
  string action = 3;
  google.protobuf.Struct old_value = 4;
  google.protobuf.Struct new_value = 5;
  string action_date = 6;
  string performed_by = 7;
  string notes = 8;
  google.protobuf.Timestamp created_at = 9;
  google.protobuf.Timestamp updated_at = 10;
}

message CreateIoTDeviceHistoryRequest {
  string device_id = 1 [(buf.validate.field).string.min_len = 1];
  string action = 2 [
    (buf.validate.field).string.in = "install",
    (buf.validate.field).string.in = "update_config",
    (buf.validate.field).string.in = "firmware_update",
    (buf.validate.field).string.in = "relocate",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "deactivate",
    (buf.validate.field).string.in = "reactivate"
  ];
  google.protobuf.Struct old_value = 3;
  google.protobuf.Struct new_value = 4;
  string performed_by = 5 [(buf.validate.field).string.min_len = 1];
  string notes = 6 [(buf.validate.field).string.max_len = 1000];
}

message CreateIoTDeviceHistoryResponse {
  string device_id = 1 [(buf.validate.field).string.min_len = 1];
  string action = 2 [
    (buf.validate.field).string.in = "install",
    (buf.validate.field).string.in = "update_config",
    (buf.validate.field).string.in = "firmware_update",
    (buf.validate.field).string.in = "relocate",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "deactivate",
    (buf.validate.field).string.in = "reactivate"
  ];
  google.protobuf.Struct old_value = 3;
  google.protobuf.Struct new_value = 4;
  string action_date = 5 [(buf.validate.field).string.min_len = 1];
  string performed_by = 6 [(buf.validate.field).string.min_len = 1];
  string notes = 7 [(buf.validate.field).string.max_len = 1000];
}

message GetIoTDeviceHistoryRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message GetIoTDeviceHistoryResponse {
  IoTDeviceHistory history = 1;
}

message DeleteIoTDeviceHistoryRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message DeleteIoTDeviceHistoryResponse {
  bool success = 1;
}

message ListIoTDeviceHistoryRequest {
  IoTDeviceHistoryFilters filters = 1;
  common.PaginationRequest pagination = 2;
}

message ListIoTDeviceHistoryResponse {
  repeated IoTDeviceHistory data = 1;
  common.PaginationResponse pagination = 2;
}

message IoTDeviceHistoryFilters {
  string device_id = 1 [(buf.validate.field).string.min_len = 1];
  string action = 2 [
    (buf.validate.field).string.in = "install",
    (buf.validate.field).string.in = "update_config",
    (buf.validate.field).string.in = "firmware_update",
    (buf.validate.field).string.in = "relocate",
    (buf.validate.field).string.in = "maintenance",
    (buf.validate.field).string.in = "deactivate",
    (buf.validate.field).string.in = "reactivate"
  ];
  string performed_by = 3 [(buf.validate.field).string.min_len = 1];
  string start_date = 4 [(buf.validate.field).string.min_len = 1];
  string end_date = 5 [(buf.validate.field).string.min_len = 1];
}

service IoTDeviceHistoryService {
  rpc CreateIoTDeviceHistory(CreateIoTDeviceHistoryRequest) returns (CreateIoTDeviceHistoryResponse);
  rpc GetIoTDeviceHistory(GetIoTDeviceHistoryRequest) returns (GetIoTDeviceHistoryResponse);
  rpc DeleteIoTDeviceHistory(DeleteIoTDeviceHistoryRequest) returns (DeleteIoTDeviceHistoryResponse);
  rpc ListIoTDeviceHistory(ListIoTDeviceHistoryRequest) returns (ListIoTDeviceHistoryResponse);
}


syntax = "proto3";

package sensor_data.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1;proto_sensor_data";

import "common/v1/common.proto";
import "buf/validate/validate.proto";

// SensorType (temperature, humidity, ph, light_intensity, soil_moisture, co2, water_level)
// SensorUnit (C, F, %, lux, ppm, pH)

message SensorData {
  string id = 1;
  string device_id = 2;
  string sensor_type = 3;
  double value = 4;
  string unit = 5;
  string recorded_at = 6;
  bool is_alert = 7;
  double quality_score = 8;
  string created_at = 9;
}

message CreateSensorDataRequest {
  string device_id = 1 [(buf.validate.field).string.min_len = 1];
  string sensor_type = 2 [
    (buf.validate.field).string.in = "temperature",
    (buf.validate.field).string.in = "humidity",
    (buf.validate.field).string.in = "ph",
    (buf.validate.field).string.in = "light_intensity",
    (buf.validate.field).string.in = "soil_moisture",
    (buf.validate.field).string.in = "co2",
    (buf.validate.field).string.in = "water_level"
  ];
  double value = 3 [
    (buf.validate.field).double.gte = -999999.99,
    (buf.validate.field).double.lte = 999999.99
  ];
  string unit = 4 [
    (buf.validate.field).string.in = "C",
    (buf.validate.field).string.in = "F",
    (buf.validate.field).string.in = "%",
    (buf.validate.field).string.in = "lux",
    (buf.validate.field).string.in = "ppm",
    (buf.validate.field).string.in = "pH"
  ];
  bool is_alert = 5;
  double quality_score = 6 [
    (buf.validate.field).double.gte = 0.0,
    (buf.validate.field).double.lte = 1.0
  ];
}

message CreateSensorDataResponse {
  string device_id = 1 [(buf.validate.field).string.min_len = 1];
  string sensor_type = 2 [
    (buf.validate.field).string.in = "temperature",
    (buf.validate.field).string.in = "humidity",
    (buf.validate.field).string.in = "ph",
    (buf.validate.field).string.in = "light_intensity",
    (buf.validate.field).string.in = "soil_moisture",
    (buf.validate.field).string.in = "co2",
    (buf.validate.field).string.in = "water_level"
  ];
  double value = 3 [
    (buf.validate.field).double.gte = -999999.99,
    (buf.validate.field).double.lte = 999999.99
  ];
  string unit = 4 [
    (buf.validate.field).string.in = "C",
    (buf.validate.field).string.in = "F",
    (buf.validate.field).string.in = "%",
    (buf.validate.field).string.in = "lux",
    (buf.validate.field).string.in = "ppm",
    (buf.validate.field).string.in = "pH"
  ];
  string recorded_at = 5 [(buf.validate.field).string.min_len = 1];
  bool is_alert = 6;
  double quality_score = 7 [
    (buf.validate.field).double.gte = 0.0,
    (buf.validate.field).double.lte = 1.0
  ];
  string created_at = 8 [(buf.validate.field).string.min_len = 1];
}

message GetSensorDataRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message GetSensorDataResponse {
  SensorData sensor_data = 1;
}

message DeleteSensorDataRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message DeleteSensorDataResponse {
  bool success = 1;
}

message ListSensorDataRequest {
  SensorDataFilters filters = 1;
  common.PaginationRequest pagination = 2;
}

message ListSensorDataResponse {
  repeated SensorData data = 1;
  common.PaginationResponse pagination = 2;
}

message SensorDataFilters {
  string device_id = 1 [(buf.validate.field).string.uuid = true];
  string sensor_type = 2 [
    (buf.validate.field).string.in = "temperature",
    (buf.validate.field).string.in = "humidity",
    (buf.validate.field).string.in = "ph",
    (buf.validate.field).string.in = "light_intensity",
    (buf.validate.field).string.in = "soil_moisture",
    (buf.validate.field).string.in = "co2",
    (buf.validate.field).string.in = "water_level"
  ];
  string unit = 3 [
    (buf.validate.field).string.in = "C",
    (buf.validate.field).string.in = "F",
    (buf.validate.field).string.in = "%",
    (buf.validate.field).string.in = "lux",
    (buf.validate.field).string.in = "ppm",
    (buf.validate.field).string.in = "pH"
  ];
  bool is_alert = 4;
  double min_quality_score = 5 [
    (buf.validate.field).double.gte = 0.0,
    (buf.validate.field).double.lte = 1.0
  ];
  double max_quality_score = 6 [
    (buf.validate.field).double.gte = 0.0,
    (buf.validate.field).double.lte = 1.0
  ];
  string start_date = 7 [(buf.validate.field).string.min_len = 1];
  string end_date = 8 [(buf.validate.field).string.min_len = 1];
}

service SensorDataService {
  rpc CreateSensorData(CreateSensorDataRequest) returns (CreateSensorDataResponse);
  rpc GetSensorData(GetSensorDataRequest) returns (GetSensorDataResponse);
  rpc DeleteSensorData(DeleteSensorDataRequest) returns (DeleteSensorDataResponse);
  rpc ListSensorData(ListSensorDataRequest) returns (ListSensorDataResponse);
}
