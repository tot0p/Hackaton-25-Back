package DBManager

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
)

// CheckIfUserExist checks if the user exists
func (db *DBManager) CheckIfUserExist(username string) (bool, error) {
	var count int
	err := db.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateUser creates a user
func (db *DBManager) CreateUser(username string, password string) error {
	uuid := uuid.New().String()
	_, err := db.db.Exec("INSERT INTO users (uuid,username, password) VALUES (?,?, ?)", uuid, username, password)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUUID returns the user by UUID
func (db *DBManager) GetUserByUUID(UUID string) (DBModels.User, error) {
	var user DBModels.User
	err := db.db.QueryRow("SELECT uuid, username FROM users WHERE uuid = ?", UUID).Scan(&user.UUID, &user.Username)
	if err != nil {
		return DBModels.User{}, err
	}
	return user, nil
}

// GetUserByUsernameWithPass returns the
func (db *DBManager) GetUserByUsernameWithPass(username string) (*DBModels.User, error) {
	var user DBModels.User
	err := db.db.QueryRow("SELECT uuid, username ,password FROM users WHERE username = ?", username).Scan(&user.UUID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername returns the user by username
func (db *DBManager) GetUserByUsername(username string) (*DBModels.User, error) {
	var user DBModels.User
	err := db.db.QueryRow("SELECT uuid, username FROM users WHERE username = ?", username).Scan(&user.UUID, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
