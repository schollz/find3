package database

import (
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func (d *Database) MakeTables() (err error) {
	sqlStmt := `create table keystore (key text not null primary key, value text);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Error(err)
		return
	}
	sqlStmt = `create index keystore_idx on keystore(key);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Error(err)
		return
	}
	sqlStmt = `create table sensors (timestamp integer not null primary key, family text, user text, location text);`
	_, err = d.db.Exec(sqlStmt)
	if err != nil {
		err = errors.Wrap(err, "MakeTables")
		d.logger.Error(err)
		return
	}
	return
}

// Columns will list the columns
func (d *Database) Columns() (columns []string, err error) {
	rows, err := d.db.Query("select * from sensors limit 1")
	if err != nil {
		err = errors.Wrap(err, "Columns")
		return
	}
	columns, err = rows.Columns()
	rows.Close()
	if err != nil {
		err = errors.Wrap(err, "Columns")
		return
	}
	d.logger.Info("got columns")
	return
}

// Get will retrieve the value associated with a key.
func (d *Database) Get(key string, v interface{}) (err error) {
	stmt, err := d.db.Prepare("select value from keystore where key = ?")
	if err != nil {
		return errors.Wrap(err, "problem preparing SQL")
	}
	defer stmt.Close()
	var result string
	err = stmt.QueryRow(key).Scan(&result)
	if err != nil {
		return errors.Wrap(err, "problem getting key")
	}

	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		return
	}
	d.logger.Infof("Got %v from '%s'", v, key)
	return
}

// Set will set a value in the database, when using it like a keystore.
func (d *Database) Set(key string, value interface{}) (err error) {
	var b []byte
	b, err = json.Marshal(value)
	if err != nil {
		return err
	}
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Set")
	}
	stmt, err := tx.Prepare("insert or replace into keystore(key,value) values (?, ?)")
	if err != nil {
		return errors.Wrap(err, "Set")
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, string(b))
	if err != nil {
		return errors.Wrap(err, "Set")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Set")
	}

	d.logger.Infof("set '%s' to '%s'", key, string(b))
	return
}

// // AddSensor will insert a sensor data into the database
// func (d *Database) AddSensor(s sensor.Data) (err error) {
// 	tx, err := d.db.Begin()
// 	if err != nil {
// 		return errors.Wrap(err, "AddSensor")
// 	}
// 	stmt, err := tx.Prepare("insert into sensors(key,value) values (?, ?)")
// 	if err != nil {
// 		return errors.Wrap(err, "AddSensor")
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(key, string(b))
// 	if err != nil {
// 		return errors.Wrap(err, "AddSensor")
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return errors.Wrap(err, "AddSensor")
// 	}

// }
