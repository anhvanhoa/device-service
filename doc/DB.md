-- Bảng loại thiết bị
CREATE TABLE device_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type_code VARCHAR(50) UNIQUE,
    description TEXT
);

-- 5. BẢNG THIẾT BỊ IOT
CREATE TABLE iot_devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_name VARCHAR(255) NOT NULL,
    device_type_id UUID,
    model VARCHAR(255),
    mac_address VARCHAR(17) UNIQUE,
    ip_address VARCHAR(45),
    greenhouse_id UUID,
    growing_zone_id UUID,
    installation_date DATE,
    last_maintenance_date DATE,
    battery_level INTEGER CHECK (battery_level BETWEEN 0 AND 100),
    status VARCHAR(50) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'maintenance', 'error', 'offline')),
    configuration JSONB,
    default_config JSONB NOT NULL,
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (device_type_id) REFERENCES device_types(id) ON DELETE SET NULL,
    FOREIGN KEY (greenhouse_id) REFERENCES greenhouses(id) ON DELETE CASCADE,
    FOREIGN KEY (growing_zone_id) REFERENCES growing_zones(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_iot_devices_type_status ON iot_devices(device_type_id, status);
CREATE INDEX idx_iot_devices_ip ON iot_devices(ip_address);

COMMENT ON COLUMN iot_devices.device_type_id IS 'Tham chiếu đến bảng device_types (ví dụ: temperature_sensor, humidity_sensor, ph_sensor, camera, pump, valve, fan)';
COMMENT ON COLUMN iot_devices.battery_level IS 'Mức pin (%) cho thiết bị không dây';
COMMENT ON COLUMN iot_devices.status IS 'active, inactive, maintenance, error, offline';
COMMENT ON COLUMN iot_devices.configuration IS 'Cấu hình thiết bị dạng JSON';


-- 5b. BẢNG LỊCH SỬ THIẾT BỊ
CREATE TABLE iot_device_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID NOT NULL,
    action VARCHAR(50) CHECK (action IN ('install', 'update_config', 'firmware_update', 'relocate', 'maintenance', 'deactivate', 'reactivate')),
    old_value JSONB,
    new_value JSONB,
    action_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    performed_by UUID,
    notes TEXT,

    FOREIGN KEY (device_id) REFERENCES iot_devices(id) ON DELETE CASCADE,
    FOREIGN KEY (performed_by) REFERENCES users(id)
);

CREATE INDEX idx_device_history ON iot_device_history(device_id, action_date);
COMMENT ON COLUMN iot_device_history.action IS 'install, update_config, firmware_update, relocate, maintenance, deactivate, reactivate';


-- 6. BẢNG DỮ LIỆU CẢM BIẾN
CREATE TABLE sensor_data (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID NOT NULL,
    sensor_type VARCHAR(100),
    value DECIMAL(12,4),
    unit VARCHAR(20),
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_alert BOOLEAN DEFAULT FALSE,
    quality_score DECIMAL(3,2) CHECK (quality_score >= 0 AND quality_score <= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (device_id) REFERENCES iot_devices(id) ON DELETE CASCADE
);

CREATE INDEX idx_sensor_data_device_recorded ON sensor_data(device_id, recorded_at);
CREATE INDEX idx_sensor_data_type_recorded ON sensor_data(sensor_type, recorded_at);
CREATE INDEX idx_sensor_data_alert ON sensor_data(is_alert, recorded_at);

COMMENT ON COLUMN sensor_data.sensor_type IS 'temperature, humidity, ph, light_intensity, soil_moisture, co2, water_level';
COMMENT ON COLUMN sensor_data.unit IS 'C, F, %, lux, ppm, etc.';
COMMENT ON COLUMN sensor_data.quality_score IS 'Điểm chất lượng dữ liệu (0-1)';
