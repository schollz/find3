package models

type LocationAnalysis struct {
	LocationNames map[string]string `json:"location_names"`
	Predictions   []struct {
		Locations     []string  `json:"locations"`
		Name          string    `json:"name"`
		Probabilities []float64 `json:"probabilities"`
	} `json:"predictions"`
}
