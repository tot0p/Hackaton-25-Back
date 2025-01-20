package DBManager

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Open opens the database
func (db *DBManager) Open() {
	//open file
	dbsql, err := sql.Open("sqlite3", db.fileName)
	if err != nil {
		panic(err)
	}
	db.db = dbsql
}
