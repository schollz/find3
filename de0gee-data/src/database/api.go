package database

import (
	_ "github.com/mattn/go-sqlite3"
)

func (d *Database) makeTables() (err error) {
	sqlStmt := `
create table keystore (key text not null primary key, value text);
`
	_, err = d.db.Exec(sqlStmt)
	return
}
