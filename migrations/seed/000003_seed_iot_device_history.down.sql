-- Xóa seed data cho bảng iot_device_history
DELETE FROM iot_device_history WHERE device_id IN (
    SELECT id FROM iot_devices WHERE device_name IN (
        'Temp Sensor Zone A1',
        'Humidity Sensor Zone A1',
        'pH Sensor Zone A1',
        'Temp Sensor Zone A2',
        'Camera Zone A1',
        'Water Pump Zone A1',
        'LED Grow Light Zone A1'
    )
);
