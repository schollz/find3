package models

// GPS is the typical data structure for storing GPS data.
type GPS struct {
	// Timestamp is the unique identifier, the time in milliseconds
	Timestamp int64   `json:"t"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Altitude  float64 `json:"alt"`
	Mac       string  `json:"mac"`
}

type FingerprintWithGPS struct {
	GPS         GPS        `json:"gps"`
	Fingerprint SensorData `json:"fingerprint"`
}
