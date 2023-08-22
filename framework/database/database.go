// framework/database/database.go

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "C:/sqlite/tasks.db")
	if err != nil {
		panic(err)
	}
	return db
}
