package entity

import "time"

type DeviceType struct {
	tableName   struct{} `pg:"device_types"`
	ID          string
	TypeCode    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (d *DeviceType) TableName() any {
	return d.tableName
}
