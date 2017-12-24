package sensor

// Data is the typical data structure for storing sensor data.
type Data struct {
	Timestamp int                               `json:"t"`
	Family    string                            `json:"f"`
	Device    string                            `json:"d"`
	Sensors   map[string]map[string]interface{} `json:"s"`
}
