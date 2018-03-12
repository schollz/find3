package models

import (
	"time"
)

type ReverseRollingData struct {
	HasData        bool
	Family         string
	Datas          []SensorData
	Timestamp      time.Time
	TimeBlock      time.Duration
	MinimumPassive int
	DeviceLocation map[string]string // Device -> Location for learning
	DeviceGPS      map[string]GPS    // Device -> GPS for learning
}
