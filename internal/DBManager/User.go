package DBManager

import "github.com/google/uuid"

func (db *DBManager) CheckIfUserExist(username string) (bool, error) {
	var count int
	err := db.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *DBManager) CreateUser(username string, password string) error {
	uuid := uuid.New().String()
	_, err := db.db.Exec("INSERT INTO users (uuid,username, password) VALUES (?,?, ?)", uuid, username, password)
	if err != nil {
		return err
	}
	return nil
}
