package user

import (
	"errors"
	"log"
	"strconv"

	"github.com/novapo/go-spa-kit/db"

	"code.google.com/p/go-sqlite/go1/sqlite3"
)

// User repressents a user in db, User is 1-indexed
type User struct {
	ID           int
	Name         string
	PasswordHash string
}

// Save saves User u in db
func (u *User) Save() error {
	if u.ID > 0 {
		//update query
		db, err := db.Open()
		if err != nil {
			return err
		}
		db.Update(u)
		// id := u.ID
		// sqlStmt := "update user SET name = " + u.Name + ", password = " + u.PasswordHash + "where id = " + strconv.Itoa(id)
		// err := Sql.Exec(sqlStmt)
		// if err != nil {
		// 	log.Printf("%q: %s\n", err, sqlStmt)
		// }
		return nil
	}
	// insert query
	idCount++
	id := idCount
	sqlStmt := "insert into user (name, password) values(" + u.Name + ", " + u.PasswordHash + ")"
	err := Sql.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

// Delete deletes User u from db
func (u User) Delete(Sql *sqlite3.Conn) error {
	if u.ID < 1 {
		return errors.New("id not found")
	}
	sqlStmt := "update user SET name = NULL, password = NULL, where id = " + strconv.Itoa(u.ID)
	err := Sql.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
