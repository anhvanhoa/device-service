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

    FOREIGN KEY (device_id) REFERENCES iot_devices(id) ON DELETE CASCADE
);

CREATE TRIGGER update_iot_device_history_updated_at 
    BEFORE UPDATE ON iot_device_history 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

CREATE INDEX idx_device_history ON iot_device_history(device_id, action_date);
COMMENT ON COLUMN iot_device_history.action IS 'install, update_config, firmware_update, relocate, maintenance, deactivate, reactivate';
