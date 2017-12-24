package sensor

type Data struct {
	Timestamp int                               `json:"t"`
	Family    string                            `json:"f"`
	User      string                            `json:"u"`
	Sensors   map[string]map[string]interface{} `json:"s"`
}
