package DBManager

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/tot0p/Hackaton-25-Back/internal/models/DBModels"
)

// AddTravel adds a travel to the database
func (db *DBManager) AddTravel(travel DBModels.Travel) error {
	uuid := uuid.New().String()
	_, err := db.db.Exec("INSERT INTO travels (UUID, userUUID , StartLocation, EndLocation, Distance , Duration , CO2 , TransportType ) VALUES (?,?,?,?,?,?,?,?)", uuid, travel.UserUUID, travel.StartLocation, travel.EndLocation, travel.Distance, travel.Duration, travel.CO2, travel.TransportType)
	if err != nil {
		return err
	}
	return nil
}

// GetTravels returns the travels of the user
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

// GetCO2OfMonth returns the CO2 of the user of the current month
func (db *DBManager) GetCO2OfMonth(userUUID string) (float64, error) {
	var CO2 float64
	err := db.db.QueryRow("SELECT SUM(CO2) FROM travels WHERE userUUID = ? AND strftime('%m',Date) = strftime('%m',CURRENT_DATE)", userUUID).Scan(&CO2)
	if err != nil {
		if err.Error() == "sql: Scan error on column index 0, name \"SUM(CO2)\": converting NULL to float64 is unsupported" {
			return 0, nil
		}
		return 0, err
	}
	return CO2, nil
}

// GetCO2OfPrevMonth returns the CO2 of the user of the previous month
func (db *DBManager) GetCO2OfPrevMonth(userUUID string) (float64, error) {
	var CO2 float64
	err := db.db.QueryRow("SELECT SUM(CO2) FROM travels WHERE userUUID = ? AND CAST(strftime('%m',Date) AS INTEGER) = (CAST(strftime('%m',CURRENT_DATE) AS INTEGER) -1)", userUUID).Scan(&CO2)
	if err != nil {
		if err.Error() == "sql: Scan error on column index 0, name \"SUM(CO2)\": converting NULL to float64 is unsupported" {
			return -1, nil
		}
		return -1, err
	}
	return CO2, nil
}
