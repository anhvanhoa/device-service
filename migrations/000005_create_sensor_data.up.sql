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

CREATE TRIGGER update_sensor_data_updated_at 
    BEFORE UPDATE ON sensor_data 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

COMMENT ON COLUMN sensor_data.sensor_type IS 'temperature, humidity, ph, light_intensity, soil_moisture, co2, water_level';
COMMENT ON COLUMN sensor_data.unit IS 'C, F, %, lux, ppm, etc.';
COMMENT ON COLUMN sensor_data.quality_score IS 'Điểm chất lượng dữ liệu (0-1)';
