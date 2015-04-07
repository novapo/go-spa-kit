package user

import (
	"log"
	"strconv"

	"code.google.com/p/go-sqlite/go1/sqlite3"
)

// User repressents a user in db, User is 1-indexed
type User struct {
	ID           int
	Name         string
	PasswordHash string
}

var idCount = 0

// Save saves User u in db
func (u User) Save(Sql *sqlite3.Conn) {
	if idCount > u.ID && u.ID > 0 {
		//update query
		id := u.ID
		sqlStmt := "update user SET name = " + u.Name + ", password = " + u.PasswordHash + "where id = " + strconv.Itoa(id)
		err := Sql.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
		}
		return
	}
	// insert query
	idCount++
	id := idCount
	sqlStmt := "insert into user values(" + strconv.Itoa(id) + ", " + u.Name + ", " + u.PasswordHash + ")"
	err := Sql.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

// Delete deletes User u from db
func (u User) Delete() {

}
