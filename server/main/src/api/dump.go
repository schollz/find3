package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/schollz/find3/server/main/src/models"

	"github.com/schollz/find3/server/main/src/database"
)

func Dump(family string) (err error) {
	defer logger.Log.Flush()
	// gather the data
	db, err := database.Open(family, true)
	if err != nil {
		return
	}
	defer db.Close()
	datasLearn, err := db.GetAllForClassification()
	if err != nil {
		return
	}
	datasTrack, err := db.GetAllNotForClassification()
	if err != nil {
		return
	}

	if len(datasLearn) == 0 && len(datasTrack) == 0 {
		err = errors.New("no data to dump for " + family)
	}
	if len(datasLearn) > 0 {
		err = writeDatas(family, "learn", datasLearn)
		if err != nil {
			return
		}
	}
	if len(datasTrack) > 0 {
		err = writeDatas(family, "track", datasTrack)
		if err != nil {
			return
		}

	}

	return
}

func writeDatas(family string, name string, datas []models.SensorData) (err error) {
	fname := fmt.Sprintf("%s.%s.%d.jsons", family, name, datas[len(datas)-1].Timestamp)
	os.Remove(fname)
	f, err := os.Create(fname)
	if err != nil {
		return
	}
	defer f.Close()
	for _, data := range datas {
		bData, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return errMarshal
		}
		f.Write(bData)
		f.Write([]byte("\n"))
	}
	f.Sync()
	logger.Log.Infof("dumped data to %s", fname)
	return
}
