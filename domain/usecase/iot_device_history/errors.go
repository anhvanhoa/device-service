package iot_device_history

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrDeviceHistoryNotFound = oops.New("Không tìm thấy lịch sử thiết bị")
	ErrInvalidAction         = oops.New("Trạng thái thiết bị không hợp lệ")
	ErrInvalidDateFormat     = oops.New("Định dạng ngày tháng không hợp lệ")
)
