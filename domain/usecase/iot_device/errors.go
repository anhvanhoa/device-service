package iot_device

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrIoTDeviceNotFound       = oops.New("Không tìm thấy thiết bị IoT")
	ErrMacAddressAlreadyExists = oops.New("Địa chỉ MAC đã tồn tại")
	ErrInvalidDateFormat       = oops.New("Định dạng ngày tháng không hợp lệ")
	ErrInvalidStatus           = oops.New("Trạng thái thiết bị không hợp lệ")

	// Control device errors
	ErrDeviceNotFound   = oops.New("Không tìm thấy thiết bị")
	ErrDeviceIDRequired = oops.New("ID thiết bị là bắt buộc")
	ErrActionRequired   = oops.New("Hành động là bắt buộc")
	ErrInvalidAction    = oops.New("Hành động không hợp lệ")
	ErrDeviceAlreadyOn  = oops.New("Thiết bị đã được bật")
	ErrDeviceAlreadyOff = oops.New("Thiết bị đã được tắt")
	ErrPublishFailed    = oops.New("Lỗi khi gửi thông tin điều khiển đến MQTT")
)
