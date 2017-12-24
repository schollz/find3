package sensor

type Data struct {
	Time           int                           `json:"t"`
	Family         string                        `json:"f"`
	User           string                        `json:"u"`
	Location       string                        `json:"l"`
	Authentication string                        `json:"a"`
	Sensors        map[string]map[string]float64 `json:"s"`
}
