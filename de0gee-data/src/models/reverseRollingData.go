package models

import (
	"time"
)

type ReverseRollingData struct {
	HasData        bool
	Family         string
	Datas          []SensorData
	Timestamp      time.Time
	DeviceLocation map[string]string // Device -> Location for learning
}
