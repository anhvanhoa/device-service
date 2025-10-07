package sensor_data

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrSensorDataNotFound  = oops.New("Không tìm thấy dữ liệu cảm biến")
	ErrInvalidSensorType   = oops.New("Loại cảm biến không hợp lệ")
	ErrInvalidQualityScore = oops.New("Điểm chất lượng không hợp lệ (phải nằm trong khoảng 0 và 1)")
	ErrInvalidDateFormat   = oops.New("Định dạng ngày tháng không hợp lệ")
)
