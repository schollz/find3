package models

import "time"

type ByLocationDevice struct {
	Device      string    `json:"device"`
	Timestamp   time.Time `json:"timestamp"`
	Probability float64   `json:"probability"`
	Randomized  bool      `json:"random_mac"`
	NumScanners int       `json:"num_scanners"`
}

type ByLocation struct {
	Devices  []ByLocationDevice `json:"device"`
	Location string             `json:"location"`
	Total    int                `json:"total"`
}
