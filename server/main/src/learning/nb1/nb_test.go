package nb1

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	database.DataFolder, _ = filepath.Abs("../../../data")

	d, err := database.Open("schollz")
	assert.Nil(t, err)
	datas, err := d.GetAllForClassification()
	assert.Nil(t, err)
	d.Close()

	nb1 := New()
	err = nb1.Fit(datas[1:])
	assert.Nil(t, err)

	pl, err := nb1.Classify(datas[0])
	assert.Nil(t, err)
	fmt.Println(datas[0].Location)
	fmt.Println(pl)

	pl, err = nb1.Classify(datas[1])
	assert.Nil(t, err)
	fmt.Println(datas[1].Location)
	fmt.Println(pl)
}
