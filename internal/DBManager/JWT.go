package DBManager

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

func (db *DBManager) JWTRegister(token string, userUUID string, expireAt int64) error {
	uuid := uuid.New().String()
	_, err := db.db.Exec("INSERT INTO JWT (UUID ,token, userUUID, ExpiresAt) VALUES (?,?,?, ?)", uuid, userUUID, token, expireAt)
	return err
}

func (db *DBManager) JWTGetByUserUUIDIfNotExpired(userUUID string) (string, error) {
	var token string
	exp := time.Now().Unix()
	err := db.db.QueryRow("SELECT token FROM JWT WHERE userUUID = ? AND ExpiresAt > ?", userUUID, exp).Scan(&token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.db.Exec("DELETE FROM JWT WHERE ExpiresAt < ?", exp)
			if err != nil {
				return "", err
			}
		}
		return "", err
	}
	return token, nil
}
