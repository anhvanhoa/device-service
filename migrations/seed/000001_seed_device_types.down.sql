-- Xóa seed data cho bảng device_types
DELETE FROM device_types WHERE type_code IN (
    'temperature_sensor',
    'humidity_sensor', 
    'ph_sensor',
    'light_sensor',
    'soil_moisture_sensor',
    'co2_sensor',
    'water_level_sensor',
    'camera',
    'pump',
    'valve',
    'fan',
    'heater',
    'cooler',
    'led_grow_light',
    'irrigation_system',
    'nutrient_dispenser'
);
