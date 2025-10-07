package device_type

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrDeviceTypeNotFound      = oops.New("Không tìm thấy loại thiết bị")
	ErrDeviceTypeAlreadyExists = oops.New("Loại thiết bị đã tồn tại")
)
