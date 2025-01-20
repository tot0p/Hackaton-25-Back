package DBModels

type Travel struct {
	UUID          string  `json:"uuid"`
	UserUUID      string  `json:"userUUID"`
	Date          string  `json:"date"`
	StartLocation string  `json:"startLocation"`
	EndLocation   string  `json:"endLocation"`
	Distance      float64 `json:"distance"`
	Duration      int64   `json:"duration"`
	CO2           float64 `json:"co2"`
	TransportType string  `json:"transportType"`
}
