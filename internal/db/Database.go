package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Database struct {
	db *sql.DB
}

func (db *Database) Close() {
	err := db.db.Close()
	if err != nil {
		return
	}
}

// CreateDatabase creates a new database connection
func CreateDatabase(filename string) *Database {
	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// create file
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		err = file.Close()
		if err != nil {
			return nil
		}
	}

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	return &Database{db}
}
