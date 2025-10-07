package entity

import (
	"time"
)

type DeviceAction string

const (
	DeviceActionInstall        DeviceAction = "install"
	DeviceActionUpdateConfig   DeviceAction = "update_config"
	DeviceActionFirmwareUpdate DeviceAction = "firmware_update"
	DeviceActionRelocate       DeviceAction = "relocate"
	DeviceActionMaintenance    DeviceAction = "maintenance"
	DeviceActionDeactivate     DeviceAction = "deactivate"
	DeviceActionReactivate     DeviceAction = "reactivate"
)

type IoTDeviceHistory struct {
	tableName   struct{} `pg:"iot_device_history"`
	ID          string
	DeviceID    string
	Action      DeviceAction
	OldValue    JSONB
	NewValue    JSONB
	ActionDate  time.Time
	PerformedBy string
	Notes       *string
}

func (i *IoTDeviceHistory) TableName() any {
	return i.tableName
}
