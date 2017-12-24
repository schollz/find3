package sensor

type Data struct {
	Time           int64                         `json:"t"`
	Group          string                        `json:"g"`
	User           string                        `json:"u"`
	Authentication string                        `json:"a"`
	Sensors        map[string]map[string]float64 `json:"s"`
}
