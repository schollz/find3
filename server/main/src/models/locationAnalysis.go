package models

type LocationAnalysis struct {
	IsUnknown     bool                  `json:"is_unknown,omitempty"`
	LocationNames map[string]string     `json:"location_names"`
	Predictions   []AlgorithmPrediction `json:"predictions"`
	Guesses       []LocationPrediction  `json:"guesses,omitempty"`
}

type AlgorithmPrediction struct {
	Locations     []string  `json:"locations"`
	Name          string    `json:"name"`
	Probabilities []float64 `json:"probabilities"`
}

type LocationPrediction struct {
	Location    string  `json:"location,omitempty"`
	Probability float64 `json:"probability,omitempty"`
}
