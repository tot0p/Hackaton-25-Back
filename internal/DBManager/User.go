package DBManager

import (
	"github.com/google/uuid"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
	"github.com/tot0p/Hackaton-25-Back/internal/utils"
)

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

func (db *DBManager) Login(username, password, plateform string) (DBModels.Session, bool, error) {
	var session DBModels.Session
	var User DBModels.User
	err := db.db.QueryRow("SELECT uuid, username, password FROM users WHERE username = ?", username).Scan(&User.UUID, &User.Username, &User.Password)
	if err != nil {
		return DBModels.Session{}, false, err
	}
	ok := utils.CheckPasswordHash(password, User.Password)
	if !ok {
		return DBModels.Session{}, false, nil
	}

	session.UUID = uuid.New().String()
	session.UserUUID = User.UUID
	session.Device = plateform
	_, err = db.db.Exec("INSERT INTO sessions (uuid, userUUID, device) VALUES (?,?, ?)", session.UUID, session.UserUUID, session.Device)
	if err != nil {
		return DBModels.Session{}, false, err
	}
	return session, true, nil
}
