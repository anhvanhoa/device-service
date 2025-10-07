-- Xóa seed data cho bảng iot_devices
DELETE FROM iot_devices WHERE device_name IN (
    'Temp Sensor Zone A1',
    'Temp Sensor Zone A2', 
    'Humidity Sensor Zone A1',
    'pH Sensor Zone A1',
    'Light Sensor Zone A1',
    'Soil Moisture Zone A1',
    'Camera Zone A1',
    'Water Pump Zone A1',
    'Control Valve Zone A1',
    'Ventilation Fan Zone A1',
    'LED Grow Light Zone A1',
    'Irrigation System Zone A1'
);
