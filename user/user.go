package user

import (
	"time"

	"code.google.com/p/go-sqlite/go1/sqlite3"
)

// User repressents a user in db, User is 1-indexed
type User struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	PasswordHash string    `db:"password"`
	DeletedAt    time.Time `db:"deleted_at"`
}

// Save saves User u in db
func (u *User) Save() error {
	return nil
}

// Delete deletes User u from db
func (u User) Delete(Sql *sqlite3.Conn) error {
	return nil
}
