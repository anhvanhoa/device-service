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
    FOREIGN KEY (device_type_id) REFERENCES device_types(id) ON DELETE SET NULL
);

CREATE INDEX idx_iot_devices_type_status ON iot_devices(device_type_id, status);
CREATE INDEX idx_iot_devices_ip ON iot_devices(ip_address);

CREATE TRIGGER update_iot_devices_updated_at 
    BEFORE UPDATE ON iot_devices 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();