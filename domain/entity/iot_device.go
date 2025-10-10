package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type DeviceStatus string

const (
	DeviceStatusActive      DeviceStatus = "active"
	DeviceStatusInactive    DeviceStatus = "inactive"
	DeviceStatusMaintenance DeviceStatus = "maintenance"
	DeviceStatusError       DeviceStatus = "error"
	DeviceStatusOffline     DeviceStatus = "offline"
)

type JSONB map[string]any

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (j *JSONB) Scan(value any) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), j)
	}
	return json.Unmarshal(bytes, j)
}

type IoTDevice struct {
	tableName           struct{} `pg:"iot_devices"`
	ID                  string
	DeviceName          string
	DeviceTypeID        string
	Model               string
	MacAddress          string
	IPAddress           string
	GreenhouseID        string
	GrowingZoneID       string
	InstallationDate    *time.Time
	LastMaintenanceDate *time.Time
	BatteryLevel        int
	Status              DeviceStatus
	DefaultConfig       JSONB
	ReadInterval        int
	AlertEnabled        bool
	AlertThresholdHigh  float64
	AlertThresholdLow   float64
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

func (i *IoTDevice) TableName() any {
	return i.tableName
}

func (d *IoTDevice) IsBatteryLow() bool {
	if d.BatteryLevel == 0 {
		return false
	}
	return d.BatteryLevel < 20
}

func (d *IoTDevice) IsOnline() bool {
	return d.Status == DeviceStatusActive
}

func (d *IoTDevice) UpdateBatteryLevel(level int) {
	if level >= 0 && level <= 100 {
		d.BatteryLevel = level
		now := time.Now()
		d.UpdatedAt = &now
	}
}

func (d *IoTDevice) UpdateStatus(status DeviceStatus) {
	d.Status = status
	now := time.Now()
	d.UpdatedAt = &now
}
