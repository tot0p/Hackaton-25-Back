package DBManager

// Close closes the database
func (db *DBManager) Close() {
	err := db.db.Close()
	if err != nil {
		panic(err)
	}
}
