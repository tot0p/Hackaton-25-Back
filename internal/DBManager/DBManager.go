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

func (db *DBManager) Commit() error {
	if db.tx != nil {
		err := db.tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBManager) Rollback() error {
	if db.tx != nil {
		err := db.tx.Rollback()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBManager) Validate() {
	db.tx = nil
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
