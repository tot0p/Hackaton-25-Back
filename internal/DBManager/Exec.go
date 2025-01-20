package DBManager

import "database/sql"

// Exec executes a query
func (db *DBManager) Exec(query string, args ...any) (sql.Result, error) {
	err := db.begin()
	if err != nil {
		return nil, err
	}
	return db.tx.Exec(query, args...)
}
