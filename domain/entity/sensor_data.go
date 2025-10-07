package entity

import (
	"time"
)

type SensorType string

const (
	SensorTypeTemperature    SensorType = "temperature"
	SensorTypeHumidity       SensorType = "humidity"
	SensorTypePH             SensorType = "ph"
	SensorTypeLightIntensity SensorType = "light_intensity"
	SensorTypeSoilMoisture   SensorType = "soil_moisture"
	SensorTypeCO2            SensorType = "co2"
	SensorTypeWaterLevel     SensorType = "water_level"
)

type SensorUnit string

const (
	UnitCelsius    SensorUnit = "C"
	UnitFahrenheit SensorUnit = "F"
	UnitPercent    SensorUnit = "%"
	UnitLux        SensorUnit = "lux"
	SensorUnitPPM  SensorUnit = "ppm"
	SensorUnitPH   SensorUnit = "pH"
)

type SensorData struct {
	tableName    struct{} `pg:"sensor_data"`
	ID           string
	DeviceID     string
	SensorType   string
	Value        float64
	Unit         string
	RecordedAt   time.Time
	IsAlert      bool
	QualityScore float64
	CreatedAt    time.Time
}

func (s *SensorData) TableName() any {
	return s.tableName
}

func (s *SensorData) SetAlert(isAlert bool) {
	s.IsAlert = isAlert
}

func (s *SensorData) SetQualityScore(score float64) {
	if score >= 0 && score <= 1 {
		s.QualityScore = score
	}
}

func (s *SensorData) IsHighQuality() bool {
	return s.QualityScore >= 0.8
}

func (s *SensorData) IsLowQuality() bool {
	return s.QualityScore < 0.5
}

func (s *SensorData) GetSensorTypeEnum() SensorType {
	return SensorType(s.SensorType)
}

func (s *SensorData) GetUnitEnum() SensorUnit {
	return SensorUnit(s.Unit)
}
