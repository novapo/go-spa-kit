package db

import (
	"log"
	"os"
)

// CreateDatabase creates the sqlite file for users sessions etc
func CreateDatabase(filename string) {
	Sqlite, err := os.Open("./db.sqlite") // For read access.
	if err != nil {
		os.Create("./db.sqlite")
	}
	defer Sqlite.Close()
	sqlStmt := `
    create table user (id integer not null primary key, name text, password text);
    create table session(id integer not null primary key, userId integer not null foreign key);`
	_, err = Sqlite.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
