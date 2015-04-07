package db

import (
	"database/sql"
	"log"
	// needed for Open()
	_ "github.com/mattn/go-sqlite3"
)

var db *DatabaseManager

// DatabaseManager has filedescriptor for database
type DatabaseManager struct {
	Db *sql.DB
}

// Model implements all Funcs needed to insert, delete or update database
type Model interface {
	GetTableName() string
	ColumnsToString() string
	ValuesToString() string
}

// Open opens db & returns filedescriptor for db
func Open() (*DatabaseManager, error) {
	if db == nil {
		db = &DatabaseManager{}
		var err error
		db.Db, err = sql.Open("sqlite3", "./database.db")
		if err != nil {
			return nil, err
		}

		if err = db.init(); err != nil {
			db = nil
			return nil, err
		}
	}
	return db, nil
}

// init creates the sqlite file for users sessions etc and indexes those
func (dm *DatabaseManager) init() error {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (id integer not null primary key AUTOINCREMENT, name text, password text, deleted_at datetime);
		CREATE INDEX IF NOT EXISTS index_users_on_deleted_at ON users (deleted_at);
		CREATE TABLE IF NOT EXISTS sessions(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL );`
	_, err := dm.Db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (dm *DatabaseManager) insert(model Model) error {
	dbase, err := Open()
	// TODO Fehlerbehandlung
	sqlStmt := "INSERT INTO " + model.GetTableName() + "(" + model.ColumnsToString() + ") VALUES(" + model.ValuesToString() + ")"
	_, err = dbase.Db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DatabaseManager) update(model Model) error {
	dbase, err := Open()
	// TODO Fehlerbehandlung
	sqlStmt := "UPDATE " + model.GetTableName() + "(" + model.ColumnsToString() + ") VALUES(" + model.ValuesToString() + ")"
	_, err = dbase.Db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
