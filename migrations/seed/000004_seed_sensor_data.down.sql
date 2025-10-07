-- Xóa seed data cho bảng sensor_data
DELETE FROM sensor_data WHERE device_id IN (
    SELECT id FROM iot_devices WHERE device_name IN (
        'Temp Sensor Zone A1',
        'Temp Sensor Zone A2',
        'Humidity Sensor Zone A1',
        'pH Sensor Zone A1',
        'Light Sensor Zone A1',
        'Soil Moisture Zone A1',
        'Water Pump Zone A1'
    )
);
