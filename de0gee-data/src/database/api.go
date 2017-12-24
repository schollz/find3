package database

import (
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func (d *Database) makeTables() (err error) {
	sqlStmt := `
create table keystore (key text not null primary key, value text);
`
	_, err = d.db.Exec(sqlStmt)
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
		return errors.Wrap(err, "problem with set")
	}
	stmt, err := tx.Prepare("insert or replace into keystore(key,value) values (?, ?)")
	if err != nil {
		return errors.Wrap(err, "problem preparing SQL")
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, string(b))
	if err != nil {
		return errors.Wrap(err, "problem executing SQL")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "problem committing")
	}

	d.logger.Infof("set '%s' to '%s'", key, string(b))
	return
}
