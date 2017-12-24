package sensor

// Data is the typical data structure for storing sensor data.
type Data struct {
	Timestamp float64                           `json:"t"`
	Family    string                            `json:"f"`
	Device    string                            `json:"d"`
	Sensors   map[string]map[string]interface{} `json:"s"`
}

// FINDFingerprint is the prototypical information from the fingerprinting device
type FINDFingerprint struct {
	Group           string   `json:"group"`
	Username        string   `json:"username"`
	Location        string   `json:"location"`
	Timestamp       int64    `json:"timestamp"`
	WifiFingerprint []Router `json:"wifi-fingerprint"`
}

// Router is the router information for each invdividual mac address
type Router struct {
	Mac  string `json:"mac"`
	Rssi int    `json:"rssi"`
}

// Convert will convert a FINDFingerprint into the new type of data,
// for backwards compatibility with FIND.
func (f FINDFingerprint) Convert() (d Data) {
	d = Data{
		Timestamp: float64(f.Timestamp),
		Family:    f.Group,
		Device:    f.Username,
		Sensors:   make(map[string]map[string]interface{}),
	}
	if f.Location != "" {
		d.Sensors["location"] = make(map[string]interface{})
		d.Sensors["location"][f.Location] = 1.0
	}
	if len(f.WifiFingerprint) > 0 {
		d.Sensors["wifi"] = make(map[string]interface{})
		for _, fingerprint := range f.WifiFingerprint {
			d.Sensors["wifi"][fingerprint.Mac] = float64(fingerprint.Rssi)
		}
	}
	return
}
