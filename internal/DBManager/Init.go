package DBManager

import (
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

// Init initializes the database
func (db *DBManager) Init() {
	content, err := utils.ReadFile("./sql/init.sql")
	if err != nil {
		panic(err)
	}
	_, err = db.db.Exec(content)
	if err != nil {
		panic(err)
	}
}
