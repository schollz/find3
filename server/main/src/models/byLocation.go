package models

import "time"

type ByLocationDevice struct {
	Device      string    `json:"device"`
	Vendor      string    `json:"vendor,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
	Probability float64   `json:"probability"`
	Randomized  bool      `json:"randomized"`
	NumScanners int       `json:"num_scanners"`
	ActiveMins  int       `json:"active_mins"`
	FirstSeen   time.Time `json:"first_seen"`
}

type ByLocation struct {
	Devices  []ByLocationDevice `json:"devices"`
	Location string             `json:"location"`
	Total    int                `json:"total"`
}
