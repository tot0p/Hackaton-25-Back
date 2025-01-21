package DBManager

import (
	"database/sql"
	"os"
)

// DBManager manages the database
type DBManager struct {
	fileName string
	db       *sql.DB
	tx       *sql.Tx
}

// Open opens the database
func (db *DBManager) begin() error {
	if db.tx == nil {
		tx, err := db.db.Begin()
		if err != nil {
			return err
		}
		db.tx = tx
	}
	return nil
}

// NewDBManager creates a new DBManager
func NewDBManager(fileName string) *DBManager {
	//check if file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//create file
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()
	}

	return &DBManager{
		fileName: fileName,
	}
}
