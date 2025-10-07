package iot_device

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrIoTDeviceNotFound       = oops.New("Không tìm thấy thiết bị IoT")
	ErrMacAddressAlreadyExists = oops.New("Địa chỉ MAC đã tồn tại")
	ErrInvalidDateFormat       = oops.New("Định dạng ngày tháng không hợp lệ")
	ErrInvalidStatus           = oops.New("Trạng thái thiết bị không hợp lệ")
)
