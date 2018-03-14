package nb2

import (
	"errors"
	"math"
	"sort"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
)

// Algorithm defines the basic structure
type Algorithm struct {
	Data     map[string]map[string]float64
	isLoaded bool
}

// New returns new algorithm
func New() *Algorithm {
	n := new(Algorithm)
	n.Data = make(map[string]map[string]float64)
	n.isLoaded = false
	return n
}

// Fit will take the data and learn it
func (a *Algorithm) Fit(datas []models.SensorData) (err error) {
	if len(datas) == 0 {
		err = errors.New("no data")
		return
	}
	a.Data = make(map[string]map[string]float64)
	locationTotals := make(map[string]float64)
	for _, data := range datas {
		if _, ok := a.Data[data.Location]; !ok {
			a.Data[data.Location] = make(map[string]float64)
			locationTotals[data.Location] = float64(0)
		}
		locationTotals[data.Location]++
		for sensorType := range data.Sensors {
			for sensor := range data.Sensors[sensorType] {
				mac := sensorType + "-" + sensor
				if _, ok := a.Data[data.Location][mac]; !ok {
					a.Data[data.Location][mac] = float64(0)
				}
				a.Data[data.Location][mac]++
			}
		}
	}
	// normalize each location
	for loc := range a.Data {
		for mac := range a.Data[loc] {
			a.Data[loc][mac] = a.Data[loc][mac] / locationTotals[loc]
		}
	}
	db, err := database.Open(datas[0].Family)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Set("NB2", a.Data)
	return
}

// Classify will classify the specified data
func (a *Algorithm) Classify(data models.SensorData) (pl PairList, err error) {
	// load data if not already
	if !a.isLoaded {
		db, err2 := database.Open(data.Family, true)
		if err2 != nil {
			err = err2
			return
		}
		err = db.Get("NB2", &a.Data)
		db.Close()
		if err != nil {
			return
		}
		a.isLoaded = true
	}
	if len(a.Data) == 0 {
		err = errors.New("need to fit first")
		return
	}

	numLocations := float64(len(a.Data))
	NA := 1 / numLocations
	NnotA := 1 - NA
	Ps := make(map[string][]float64)
	for location := range a.Data {
		Ps[location] = []float64{}
	}
	for sensorType := range data.Sensors {
		for name := range data.Sensors[sensorType] {
			mac := sensorType + "-" + name
			val := int(data.Sensors[sensorType][name].(float64))
			for location := range Ps {
				PA := a.probMacGivenLocation(mac, val, location, true)
				PnotA := a.probMacGivenLocation(mac, val, location, false)
				P := PA * NA / (PA*NA + PnotA*NnotA)
				Ps[location] = append(Ps[location], math.Log(P))
			}
		}
	}
	PsumTotal := float64(0)
	Psum := make(map[string]float64)
	for location := range Ps {
		Psum[location] = float64(0)
		for _, v := range Ps[location] {
			Psum[location] += v
		}
		Psum[location] = math.Exp(Psum[location])
		PsumTotal += Psum[location]
	}
	for location := range Psum {
		Psum[location] = Psum[location] / PsumTotal
	}

	pl = make(PairList, len(Psum))
	i := 0
	for k, v := range Psum {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (a *Algorithm) probMacGivenLocation(mac string, val int, loc string, positive bool) (P float64) {
	P = 0.005

	numerator := float64(0)
	if positive {
		// positive: find count of mac in loc
		if v, ok := a.Data[loc][mac]; ok {
			numerator = float64(v)
		}
	} else {
		// NOT positive: find count of mac NOT in loc
		for locX := range a.Data {
			if locX != loc {
				if v, ok := a.Data[locX][mac]; ok {
					numerator += float64(v)
				}
			}
		}
	}
	// find total count of mac
	denominator := float64(0)
	for locX := range a.Data {
		if v, ok := a.Data[locX][mac]; ok {
			denominator += float64(v)
		}
	}

	if denominator > 0 && numerator > 0 {
		P = numerator / denominator
	}

	// TODO: cache it

	return
}
