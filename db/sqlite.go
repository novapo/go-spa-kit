package db

import (
	"database/sql"
	"os/user"

	"gopkg.in/gorp.v1"
	// needed for Open()
	_ "github.com/mattn/go-sqlite3"
)

var dbase *gorp.DbMap

// initDb creates the sqlite file for users sessions etc and indexes those
func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		// error
		return nil
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTable(user.User{}).SetKeys(true, "ID")

	return dbmap
}
