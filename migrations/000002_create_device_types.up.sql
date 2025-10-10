-- Bảng loại thiết bị
CREATE TABLE device_types (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
    type_code VARCHAR(50) UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_device_types_type_code ON device_types(type_code);

CREATE TRIGGER update_device_types_updated_at 
    BEFORE UPDATE ON device_types 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();