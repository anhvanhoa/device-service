-- Seed data cho bảng iot_devices
-- Lưu ý: Cần có dữ liệu trong bảng device_types, greenhouses, growing_zones, users trước
INSERT INTO iot_devices (
    device_name, 
    device_type_id, 
    model, 
    mac_address, 
    ip_address, 
    greenhouse_id, 
    growing_zone_id, 
    installation_date, 
    battery_level, 
    status, 
    default_config, 
    created_by
) VALUES
-- Cảm biến nhiệt độ
('Temp Sensor Zone A1', (SELECT id FROM device_types WHERE type_code = 'temperature_sensor'), 'DHT22', 'AA:BB:CC:DD:EE:01', '192.168.1.101', NULL, NULL, '2024-01-15', 85, 'active', 
 '{"sampling_rate": 30, "alert_threshold_high": 35, "alert_threshold_low": 15}', 
 NULL),

('Temp Sensor Zone A2', (SELECT id FROM device_types WHERE type_code = 'temperature_sensor'), 'DHT22', 'AA:BB:CC:DD:EE:02', '192.168.1.102', NULL, NULL, '2024-01-15', 92, 'active',
 '{"sampling_rate": 30, "alert_threshold_high": 35, "alert_threshold_low": 15}',
 NULL),

-- Cảm biến độ ẩm
('Humidity Sensor Zone A1', (SELECT id FROM device_types WHERE type_code = 'humidity_sensor'), 'DHT22', 'AA:BB:CC:DD:EE:03', '192.168.1.103', NULL, NULL, '2024-01-15', 78, 'active',
 '{"sampling_rate": 30, "alert_threshold_high": 80, "alert_threshold_low": 40}',
 NULL),

-- Cảm biến pH
('pH Sensor Zone A1', (SELECT id FROM device_types WHERE type_code = 'ph_sensor'), 'pH-2000', 'AA:BB:CC:DD:EE:04', '192.168.1.104', NULL, NULL, '2024-01-20', 95, 'active',
 '{"sampling_rate": 60, "calibration_date": "2024-01-20", "alert_threshold_high": 7.5, "alert_threshold_low": 5.5}',
 NULL),

-- Cảm biến ánh sáng
('Light Sensor Zone A1', (SELECT id FROM device_types WHERE type_code = 'light_sensor'), 'BH1750', 'AA:BB:CC:DD:EE:05', '192.168.1.105', NULL, NULL, '2024-01-20', 88, 'active',
 '{"sampling_rate": 60, "alert_threshold_low": 1000}',
 NULL),

-- Cảm biến độ ẩm đất
('Soil Moisture Zone A1', (SELECT id FROM device_types WHERE type_code = 'soil_moisture_sensor'), 'FC-28', 'AA:BB:CC:DD:EE:06', '192.168.1.106', NULL, NULL, '2024-01-20', 82, 'active',
 '{"sampling_rate": 120, "alert_threshold_low": 30}',
 NULL),

-- Camera giám sát
('Camera Zone A1', (SELECT id FROM device_types WHERE type_code = 'camera'), 'IP-CAM-001', 'AA:BB:CC:DD:EE:07', '192.168.1.107', NULL, NULL, '2024-01-25', 100, 'active',
 '{"resolution": "1920x1080", "fps": 30, "night_vision": true, "motion_detection": true}',
 NULL),

-- Máy bơm nước
('Water Pump Zone A1', (SELECT id FROM device_types WHERE type_code = 'pump'), 'PUMP-500W', 'AA:BB:CC:DD:EE:08', '192.168.1.108', NULL, NULL, '2024-01-25', 100, 'active',
 '{"flow_rate": "500L/h", "pressure": "2.5bar", "auto_mode": true}',
 NULL),

-- Van điều khiển
('Control Valve Zone A1', (SELECT id FROM device_types WHERE type_code = 'valve'), 'VALVE-001', 'AA:BB:CC:DD:EE:09', '192.168.1.109', NULL, NULL, '2024-01-25', 100, 'active',
 '{"valve_type": "solenoid", "voltage": "12V", "flow_control": true}',
 NULL),

-- Quạt thông gió
('Ventilation Fan Zone A1', (SELECT id FROM device_types WHERE type_code = 'fan'), 'FAN-1200RPM', 'AA:BB:CC:DD:EE:10', '192.168.1.110', NULL, NULL, '2024-01-25', 100, 'active',
 '{"speed_levels": 5, "auto_speed": true, "timer_mode": true}',
 NULL),

-- Đèn LED trồng cây
('LED Grow Light Zone A1', (SELECT id FROM device_types WHERE type_code = 'led_grow_light'), 'LED-GROW-100W', 'AA:BB:CC:DD:EE:11', '192.168.1.111', NULL, NULL, '2024-01-25', 100, 'active',
 '{"wattage": 100, "color_spectrum": "full_spectrum", "timer_schedule": "06:00-18:00"}',
 NULL),

-- Hệ thống tưới tiêu
('Irrigation System Zone A1', (SELECT id FROM device_types WHERE type_code = 'irrigation_system'), 'IRRIG-001', 'AA:BB:CC:DD:EE:12', '192.168.1.112', NULL, NULL, '2024-01-25', 100, 'active',
 '{"zones": 4, "timer_schedule": "06:00,12:00,18:00", "duration": 300}',
 NULL);
