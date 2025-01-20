package DBManager

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
)

func (db *DBManager) AddTravel(travel DBModels.Travel) error {
	uuid := uuid.New().String()
	_, err := db.db.Exec("INSERT INTO travels (UUID, userUUID , StartLocation, EndLocation, Distance , Duration , CO2 , TransportType ) VALUES (?,?,?,?,?,?,?,?)", uuid, travel.UserUUID, travel.StartLocation, travel.EndLocation, travel.Distance, travel.Duration, travel.CO2, travel.TransportType)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBManager) GetTravels(userUUID string) ([]DBModels.Travel, error) {
	var travels []DBModels.Travel
	rows, err := db.db.Query("SELECT UUID,Date, StartLocation, EndLocation, Distance , Duration , CO2 , TransportType FROM travels WHERE userUUID = ?", userUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []DBModels.Travel{}, nil
		}
		return nil, err
	}
	for rows.Next() {
		var travel DBModels.Travel
		err = rows.Scan(&travel.UUID, &travel.Date, &travel.StartLocation, &travel.EndLocation, &travel.Distance, &travel.Duration, &travel.CO2, &travel.TransportType)
		if err != nil {
			return nil, err
		}
		travel.UserUUID = userUUID
		travels = append(travels, travel)
	}
	return travels, nil
}
