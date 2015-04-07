package db

import "testing"

func TestOpen(t *testing.T) {
	db := Open()

	err := db.Db.Ping()
	if err != nil {
		t.Error(err)
	}
}
