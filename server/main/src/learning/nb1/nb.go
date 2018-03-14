package nb1

import (
	"errors"
	"math"
	"sort"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
)

// Algorithm defines the basic structure
type Algorithm struct {
	Data     map[string]map[string]map[int]int
	isLoaded bool
}

// New returns new algorithm
func New() *Algorithm {
	n := new(Algorithm)
	n.Data = make(map[string]map[string]map[int]int)
	n.isLoaded = false
	return n
}

// Fit will take the data and learn it
func (a *Algorithm) Fit(datas []models.SensorData) (err error) {
	if len(datas) == 0 {
		err = errors.New("no data")
		return
	}
	a.Data = make(map[string]map[string]map[int]int)
	for _, data := range datas {
		if _, ok := a.Data[data.Location]; !ok {
			a.Data[data.Location] = make(map[string]map[int]int)
		}
		for sensorType := range data.Sensors {
			for sensor := range data.Sensors[sensorType] {
				mac := sensorType + "-" + sensor
				val := int(data.Sensors[sensorType][sensor].(float64))
				if _, ok := a.Data[data.Location][mac]; !ok {
					a.Data[data.Location][mac] = make(map[int]int)
				}
				if _, ok := a.Data[data.Location][mac][val]; !ok {
					a.Data[data.Location][mac][val] = 0
				}
				a.Data[data.Location][mac][val]++
			}
		}
	}
	db, err := database.Open(datas[0].Family)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Set("NB1", a.Data)
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
		err = db.Get("NB1", &a.Data)
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
	valToCount := make(map[int]int)
	newValToCount := make(map[int]int)
	// positive: find val,count where loc = X and mac = X
	// not positive: find val,count where loc != X and mac = X
	for locX := range a.Data {
		if positive {
			if locX != loc {
				continue
			}
		} else {
			if locX == loc {
				continue
			}
		}
		for macX := range a.Data[locX] {
			if macX != mac {
				continue
			}
			for valX := range a.Data[locX][macX] {
				valToCount[valX] = a.Data[locX][macX][valX]
				newValToCount[valX] = a.Data[locX][macX][valX]
			}
		}
	}

	// apply gaussian filter
	width := 3
	gaussRange := []int{}
	widthCubed := int(math.Pow(float64(width), 3))
	for i := -1 * widthCubed; i <= widthCubed; i++ {
		gaussRange = append(gaussRange, i)
	}
	for _, v := range valToCount {
		for _, x := range gaussRange {
			addend := int(round(normPDF(0, float64(x), float64(width))))
			if addend <= 0 {
				continue
			}
			if _, ok := newValToCount[v+x]; !ok {
				newValToCount[v+x] = 0
			}
			newValToCount[v+x] += addend
		}
	}

	// normalize
	total := 0
	for v := range newValToCount {
		total += newValToCount[v]
	}
	probs := make(map[int]float64)
	for v := range newValToCount {
		probs[v] = float64(newValToCount[v]) / float64(total)
	}

	// return probability
	if v, ok := probs[val]; ok {
		P = v
	}

	// TODO: cache it

	return
}

func normPDF(mean, x, sd float64) float64 {
	m := sd * math.Sqrt(2*math.Pi)
	e := math.Exp(-math.Pow(x-mean, 2) / (2 * math.Pow(sd, 2)))
	return e / m
}

// https://play.golang.org/p/BkdofAFOJRh
func round(val float64) (newVal float64) {
	roundOn := 0.5
	places := 0
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
