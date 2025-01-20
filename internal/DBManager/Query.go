package DBManager

import "database/sql"

// Query executes a query
func (db *DBManager) Query(query string, args ...any) (*sql.Rows, error) {
	err := db.begin()
	if err != nil {
		return nil, err
	}
	return db.tx.Query(query, args...)
}
