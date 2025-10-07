-- Seed data cho bảng sensor_data
-- Lưu ý: Cần có dữ liệu trong bảng iot_devices trước
INSERT INTO sensor_data (
    device_id,
    sensor_type,
    value,
    unit,
    recorded_at,
    is_alert,
    quality_score
) VALUES
-- Dữ liệu cảm biến nhiệt độ
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'temperature', 25.3, 'C', '2024-01-28 08:00:00', false, 0.95),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'temperature', 26.1, 'C', '2024-01-28 08:30:00', false, 0.92),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'temperature', 27.8, 'C', '2024-01-28 09:00:00', false, 0.88),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'temperature', 29.2, 'C', '2024-01-28 09:30:00', false, 0.91),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'temperature', 31.5, 'C', '2024-01-28 10:00:00', true, 0.89),

((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A2'), 'temperature', 24.8, 'C', '2024-01-28 08:00:00', false, 0.94),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A2'), 'temperature', 25.5, 'C', '2024-01-28 08:30:00', false, 0.96),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A2'), 'temperature', 26.9, 'C', '2024-01-28 09:00:00', false, 0.93),

-- Dữ liệu cảm biến độ ẩm
((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'humidity', 65.2, '%', '2024-01-28 08:00:00', false, 0.97),
((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'humidity', 68.5, '%', '2024-01-28 08:30:00', false, 0.95),
((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'humidity', 72.1, '%', '2024-01-28 09:00:00', false, 0.93),
((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'humidity', 75.8, '%', '2024-01-28 09:30:00', false, 0.91),
((SELECT id FROM iot_devices WHERE device_name = 'Humidity Sensor Zone A1'), 'humidity', 82.3, '%', '2024-01-28 10:00:00', true, 0.88),

-- Dữ liệu cảm biến pH
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'ph', 6.8, 'pH', '2024-01-28 08:00:00', false, 0.98),
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'ph', 6.9, 'pH', '2024-01-28 09:00:00', false, 0.96),
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'ph', 7.1, 'pH', '2024-01-28 10:00:00', false, 0.94),
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'ph', 7.3, 'pH', '2024-01-28 11:00:00', false, 0.92),
((SELECT id FROM iot_devices WHERE device_name = 'pH Sensor Zone A1'), 'ph', 7.8, 'pH', '2024-01-28 12:00:00', true, 0.89),

-- Dữ liệu cảm biến ánh sáng
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 1200, 'lux', '2024-01-28 08:00:00', false, 0.95),
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 1500, 'lux', '2024-01-28 09:00:00', false, 0.93),
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 1800, 'lux', '2024-01-28 10:00:00', false, 0.91),
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 2200, 'lux', '2024-01-28 11:00:00', false, 0.89),
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 2500, 'lux', '2024-01-28 12:00:00', false, 0.87),
((SELECT id FROM iot_devices WHERE device_name = 'Light Sensor Zone A1'), 'light_intensity', 800, 'lux', '2024-01-28 18:00:00', true, 0.85),

-- Dữ liệu cảm biến độ ẩm đất
((SELECT id FROM iot_devices WHERE device_name = 'Soil Moisture Zone A1'), 'soil_moisture', 45.2, '%', '2024-01-28 08:00:00', false, 0.92),
((SELECT id FROM iot_devices WHERE device_name = 'Soil Moisture Zone A1'), 'soil_moisture', 42.8, '%', '2024-01-28 09:00:00', false, 0.94),
((SELECT id FROM iot_devices WHERE device_name = 'Soil Moisture Zone A1'), 'soil_moisture', 38.5, '%', '2024-01-28 10:00:00', false, 0.91),
((SELECT id FROM iot_devices WHERE device_name = 'Soil Moisture Zone A1'), 'soil_moisture', 35.1, '%', '2024-01-28 11:00:00', false, 0.88),
((SELECT id FROM iot_devices WHERE device_name = 'Soil Moisture Zone A1'), 'soil_moisture', 28.3, '%', '2024-01-28 12:00:00', true, 0.86),

-- Dữ liệu cảm biến CO2 (giả lập)
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'co2', 420, 'ppm', '2024-01-28 08:00:00', false, 0.96),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'co2', 435, 'ppm', '2024-01-28 09:00:00', false, 0.94),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'co2', 450, 'ppm', '2024-01-28 10:00:00', false, 0.92),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'co2', 480, 'ppm', '2024-01-28 11:00:00', false, 0.90),
((SELECT id FROM iot_devices WHERE device_name = 'Temp Sensor Zone A1'), 'co2', 520, 'ppm', '2024-01-28 12:00:00', true, 0.88),

-- Dữ liệu cảm biến mực nước (giả lập)
((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'water_level', 85.5, '%', '2024-01-28 08:00:00', false, 0.98),
((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'water_level', 82.3, '%', '2024-01-28 09:00:00', false, 0.96),
((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'water_level', 78.8, '%', '2024-01-28 10:00:00', false, 0.94),
((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'water_level', 75.2, '%', '2024-01-28 11:00:00', false, 0.92),
((SELECT id FROM iot_devices WHERE device_name = 'Water Pump Zone A1'), 'water_level', 25.1, '%', '2024-01-28 12:00:00', true, 0.90);
