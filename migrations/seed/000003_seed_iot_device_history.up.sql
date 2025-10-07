-- Seed data cho bảng iot_device_history
-- Lưu ý: Cần có dữ liệu trong bảng iot_devices trước
INSERT INTO iot_device_history (
    device_id,
    action,
    old_value,
    new_value,
    action_date,
    performed_by,
    notes
) VALUES
-- Lịch sử cài đặt thiết bị
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'install', NULL, 
 '{"device_name": "Temp Sensor Zone A1", "model": "DHT22", "mac_address": "AA:BB:CC:DD:EE:01", "ip_address": "192.168.1.101"}',
 '2024-01-15 09:00:00', NULL, 'Cài đặt cảm biến nhiệt độ cho Zone A1'),

((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'install', NULL,
 '{"device_name": "Humidity Sensor Zone A1", "model": "DHT22", "mac_address": "AA:BB:CC:DD:EE:03", "ip_address": "192.168.1.103"}',
 '2024-01-15 10:30:00', NULL, 'Cài đặt cảm biến độ ẩm cho Zone A1'),

((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'install', NULL,
 '{"device_name": "pH Sensor Zone A1", "model": "pH-2000", "mac_address": "AA:BB:CC:DD:EE:04", "ip_address": "192.168.1.104"}',
 '2024-01-20 14:00:00', NULL, 'Cài đặt cảm biến pH cho Zone A1'),

-- Lịch sử cập nhật cấu hình
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'update_config',
 '{"sampling_rate": 30, "alert_threshold_high": 35, "alert_threshold_low": 15}',
 '{"sampling_rate": 60, "alert_threshold_high": 40, "alert_threshold_low": 10}',
 '2024-01-25 11:00:00', NULL, 'Tăng tần suất lấy mẫu và điều chỉnh ngưỡng cảnh báo'),

((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'update_config',
 '{"sampling_rate": 30, "alert_threshold_high": 80, "alert_threshold_low": 40}',
 '{"sampling_rate": 45, "alert_threshold_high": 85, "alert_threshold_low": 35}',
 '2024-01-25 11:15:00', NULL, 'Điều chỉnh ngưỡng cảnh báo độ ẩm'),

-- Lịch sử bảo trì
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'maintenance',
 '{"calibration_date": "2024-01-20"}',
 '{"calibration_date": "2024-01-25", "calibration_ph_4": 4.01, "calibration_ph_7": 7.00, "calibration_ph_10": 10.01}',
 '2024-01-25 15:30:00', NULL, 'Hiệu chuẩn lại cảm biến pH'),

((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'maintenance',
 '{"last_maintenance": "2024-01-15"}',
 '{"last_maintenance": "2024-01-25", "cleaned_filter": true, "checked_pressure": "2.5bar", "lubricated_motor": true}',
 '2024-01-25 16:00:00', NULL, 'Bảo trì định kỳ máy bơm nước'),

-- Lịch sử di chuyển thiết bị
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A2'), 'relocate',
 '{"location": "Zone A1", "ip_address": "192.168.1.101"}',
 '{"location": "Zone A2", "ip_address": "192.168.1.102"}',
 '2024-01-22 13:45:00', NULL, 'Di chuyển cảm biến nhiệt độ từ Zone A1 sang Zone A2'),

-- Lịch sử cập nhật firmware
((SELECT id FROM iot_devices WHERE device_name = 'Camera Zone A1'), 'firmware_update',
 '{"firmware_version": "v1.0.0", "build_date": "2024-01-01"}',
 '{"firmware_version": "v1.2.0", "build_date": "2024-01-20", "new_features": ["motion_detection_v2", "night_vision_enhanced"]}',
 '2024-01-26 09:30:00', NULL, 'Cập nhật firmware camera với tính năng mới'),

-- Lịch sử tạm ngưng hoạt động
((SELECT id FROM iot_devices WHERE device_name = 'LED Grow Light Zone A1'), 'deactivate',
 '{"status": "active", "timer_schedule": "06:00-18:00"}',
 '{"status": "inactive", "reason": "maintenance", "deactivated_at": "2024-01-27 08:00:00"}',
 '2024-01-27 08:00:00', NULL, 'Tạm ngưng đèn LED để bảo trì hệ thống điện'),

-- Lịch sử kích hoạt lại
((SELECT id FROM iot_devices WHERE device_name = 'LED Grow Light Zone A1'), 'reactivate',
 '{"status": "inactive", "reason": "maintenance"}',
 '{"status": "active", "timer_schedule": "06:00-18:00", "reactivated_at": "2024-01-27 18:00:00"}',
 '2024-01-27 18:00:00', NULL, 'Kích hoạt lại đèn LED sau khi hoàn thành bảo trì');
