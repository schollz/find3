package api

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/schollz/find3/server/main/src/database"
	"github.com/schollz/find3/server/main/src/models"
	"github.com/stretchr/testify/assert"
)

// for testing purposes
var j = `{
	"t":1514034330040,
	"f":"familyname",
	"d":"devicename",
	"l":"bathroom",
	"s":{
		 "wifi":{
				"aa:bb:cc:dd:ee":-20,
				"ff:gg:hh:ii:jj":-80
		 },
		 "bluetooth":{
				"aa:00:cc:11:ee":-42,
				"ff:22:hh:33:jj":-50        
		 },
		 "temperature":{
				"sensor1":12,
				"sensor2":20       
		 },
		 "accelerometer":{
				"x":-1.11,
				"y":2.111,
				"z":1.23   
		 }      
	}
}`

// for testing purposes
var j2 = `{
	"t":1514034335555,
	"f":"familyname",
	"d":"devicename",
	"l":"kitchen",
	"s":{
		 "wifi":{
				"22:bb:cc:dd:ee":-19,
				"ff:gg:hh:ii:jj":-77
		 },
		 "bluetooth":{
				"aa:00:cc:11:ee":-40,
				"ff:22:hh:33:jj":-45        
		 },
		 "temperature":{
				"sensor1":10,
				"sensor2":32       
		 },
		 "accelerometer":{
				"x":-2.11,
				"y":4.111,
				"z":0.23   
		 }      
	}
}`

func BenchmarkDumpToCSV(b *testing.B) {
	var s models.SensorData
	db, _ := database.Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	db.AddSensor(s)
	json.Unmarshal([]byte(j2), &s)
	db.AddSensor(s)
	ss, _ := db.GetAllForClassification()

	db.Debug(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dumpSensorsToCSV(ss, "test.csv")
	}

}

func TestDumpSensorsToCSV(t *testing.T) {
	var s models.SensorData
	db, _ := database.Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	db.AddSensor(s)
	json.Unmarshal([]byte(j2), &s)
	db.AddSensor(s)
	ss, _ := db.GetAllForClassification()

	db.Debug(false)
	err := dumpSensorsToCSV(ss, "test.csv")
	assert.Nil(t, err)
}

func TestRisingEfficacy(t *testing.T) {
	DataFolder, _ = filepath.Abs("../../data")
	database.DataFolder = DataFolder

	db, err := database.Open("pike5")
	assert.Nil(t, err)
	datas, err := db.GetAllForClassification()
	assert.Nil(t, err)
	db.Close()
	datas = datas[:2000]
	fmt.Println(len(datas))

	datasLearn, datasTest, err := splitDataForLearning(datas, true)
	assert.Nil(t, err)
	fmt.Println(len(datasLearn))
	fmt.Println(len(datasTest))

	err = learnFromData("pike5", datasLearn)
	assert.Nil(t, err)

	algorithmEfficacy, err := findBestAlgorithm(datasTest)
	assert.Nil(t, err)
	// bA, _ := json.MarshalIndent(algorithmEfficacy, "", " ")
	// fmt.Println(string(bA))
	bestInformedness := make(map[string][]float64)
	for alg := range algorithmEfficacy {
		for loc := range algorithmEfficacy[alg] {
			if _, ok := bestInformedness[loc]; !ok {
				bestInformedness[loc] = []float64{}
			}
			bestInformedness[loc] = append(bestInformedness[loc], algorithmEfficacy[alg][loc].Informedness)
		}
	}
	for loc := range bestInformedness {
		fmt.Println(loc, Max(bestInformedness[loc]))
	}
}

// Max returns the maximum value in the input slice. If the slice is empty, Max will panic.
func Max(s []float64) float64 {
	return s[MaxIdx(s)]
}

// MaxIdx returns the index of the maximum value in the input slice. If several
// entries have the maximum value, the first such index is returned. If the slice
// is empty, MaxIdx will panic.
func MaxIdx(s []float64) int {
	if len(s) == 0 {
		panic("floats: zero slice length")
	}
	max := s[0]
	var ind int
	for i, v := range s {
		if v > max {
			max = v
			ind = i
		}
	}
	return ind
}

func TestNB(t *testing.T) {
	DataFolder, _ = filepath.Abs("../../data")
	database.DataFolder = DataFolder

	d, err := database.Open("schollz")
	assert.Nil(t, err)
	datas, err := d.GetAllForClassification()
	assert.Nil(t, err)
	d.Close()

	fmt.Println(datas)

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return
	}
	sqlStmt := `CREATE TABLE data (loc TEXT, mac TEXT, val INTEGER, count INTEGER);`
	db.Exec(sqlStmt)
	nbmap := make(map[string]map[string]map[int]int)
	for _, data := range datas {
		if _, ok := nbmap[data.Location]; !ok {
			nbmap[data.Location] = make(map[string]map[int]int)
		}
		for sensorType := range data.Sensors {
			for sensor := range data.Sensors[sensorType] {
				mac := sensorType + "-" + sensor
				val := int(data.Sensors[sensorType][sensor].(float64))
				if _, ok := nbmap[data.Location][mac]; !ok {
					nbmap[data.Location][mac] = make(map[int]int)
				}
				if _, ok := nbmap[data.Location][mac][val]; !ok {
					nbmap[data.Location][mac][val] = 0
				}
				nbmap[data.Location][mac][val]++
			}
		}
	}

	fmt.Println(nbmap)
	tx, err := db.Begin()
	assert.Nil(t, err)
	stmt, err := tx.Prepare("INSERT INTO data(loc,mac,val,count) VALUES(?,?,?,?)")
	assert.Nil(t, err)
	for loc := range nbmap {
		for mac := range nbmap[loc] {
			for val := range nbmap[loc][mac] {
				_, err = stmt.Exec(loc, mac, val, nbmap[loc][mac][val])
				assert.Nil(t, err)
			}
		}
	}
	err = stmt.Close()
	assert.Nil(t, err)
	err = tx.Commit()
	assert.Nil(t, err)

	var b bytes.Buffer
	out := bufio.NewWriter(&b)
	err = dumpDB(db, out)
	assert.Nil(t, err)
	out.Flush()
	fmt.Println(string(b.Bytes()))
	db.Close()

	os.Remove("foo.db")
	db, err = sql.Open("sqlite3", "foo.db")
	assert.Nil(t, err)
	tx, err = db.Begin()
	assert.Nil(t, err)
	for _, line := range bytes.Split(b.Bytes(), []byte("\n")) {
		sqlStmt := string(line)
		if sqlStmt == "COMMIT;" {
			continue
		}
		if len(sqlStmt) > 0 {
			stmt, err := tx.Prepare(sqlStmt)
			assert.Nil(t, err)
			_, err = stmt.Exec()
			assert.Nil(t, err)
		}
	}
	err = tx.Commit()
	assert.Nil(t, err)

	db.Close()
	bNB, _ := json.Marshal(nbmap)
	fmt.Println(len(bNB))

}

func dumpDB(db *sql.DB, out io.Writer) (err error) {
	// sqlite_master table contains the SQL CREATE statements for the database.
	schemas, err := getSchemas(db, `
		SELECT "name", "type", "sql"
		FROM "sqlite_master"
				WHERE "sql" NOT NULL AND
				"type" == 'table'
				ORDER BY "name"
`)
	if err != nil {
		return
	}

	for _, schema := range schemas {
		if schema.Name == "sqlite_sequence" {
			out.Write([]byte(`DELETE FROM "sqlite_sequence";` + "\n"))
		} else if schema.Name == "sqlite3_stat1" {
			out.Write([]byte(`ANALYZE "sqlite_master";` + "\n"))
		} else if strings.HasPrefix(schema.Name, "sqlite_") {
			continue
			// # NOTE: Virtual table support not implemented
			// #elif sql.startswith('CREATE VIRTUAL TABLE'):
			// #    qtable = table_name.replace("'", "''")
			// #    yield("INSERT INTO sqlite_master(type,name,tbl_name,rootpage,sql)"\
			// #        "VALUES('table','{0}','{0}',0,'{1}');".format(
			// #        qtable,
			// #        sql.replace("''")))
		} else {
			out.Write([]byte(fmt.Sprintf("%s;\n", schema.SQL)))
		}

		// Build the insert statement for each row of the current table
		schema.Name = strings.Replace(schema.Name, `"`, `""`, -1)
		var inserts []string
		inserts, err = getTableRows(db, schema.Name)
		if err != nil {
			return
		}
		for _, insert := range inserts {
			out.Write([]byte(fmt.Sprintf("%s;\n", insert)))
		}
	}

	// Now when the type is 'index', 'trigger', or 'view'
	schemas, err = getSchemas(db, `
SELECT "name", "type", "sql"
		FROM "sqlite_master"
				WHERE "sql" NOT NULL AND
				"type" IN ('index', 'trigger', 'view')
`)
	if err != nil {
		return
	}
	for _, schema := range schemas {
		out.Write([]byte(fmt.Sprintf("%s;\n", schema.SQL)))
	}

	out.Write([]byte("COMMIT;\n"))
	return
}

func getTableRows(db *sql.DB, tableName string) (inserts []string, err error) {
	// first get the column names
	columnNames, err := pragmaTableInfo(db, tableName)
	if err != nil {
		return
	}

	// sqlite_master table contains the SQL CREATE statements for the database.
	columnSelects := make([]string, len(columnNames))
	for i, c := range columnNames {
		columnSelects[i] = fmt.Sprintf(`'||quote("%s")||'`, strings.Replace(c, `"`, `""`, -1))
	}

	q := fmt.Sprintf(`
SELECT 'INSERT INTO "%s" VALUES(%s)' FROM "%s";
`,
		tableName,
		strings.Join(columnSelects, ","),
		tableName,
	)

	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	inserts = []string{}
	for rows.Next() {
		var insert string
		err = rows.Scan(&insert)
		if err != nil {
			return
		}
		inserts = append(inserts, insert)
	}
	err = rows.Err()
	return
}

func pragmaTableInfo(db *sql.DB, tableName string) (columnNames []string, err error) {
	// sqlite_master table contains the SQL CREATE statements for the database.
	q := `
		PRAGMA table_info("` + tableName + `")
`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	columnNames = []string{}
	for rows.Next() {
		var arr []interface{}
		for i := 0; i < 6; i++ {
			arr = append(arr, new(interface{}))
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}
		columnNames = append(columnNames, string((*arr[1].(*interface{})).([]uint8)))
	}
	err = rows.Err()
	return
}

type schema struct {
	Name string
	Type string
	SQL  string
}

func getSchemas(db *sql.DB, q string) (schemas []schema, err error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	schemas = []schema{}
	for rows.Next() {
		s := schema{}
		err = rows.Scan(&s.Name, &s.Type, &s.SQL)
		if err != nil {
			return
		}
		schemas = append(schemas, s)
	}
	err = rows.Err()
	return
}
