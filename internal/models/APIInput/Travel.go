package APIInput

// Travel represents the input for creating a travel
type Travel struct {
	StartLocation string  `json:"startLocation" xml:"startLocation" form:"startLocation"`
	EndLocation   string  `json:"endLocation" xml:"endLocation" form:"endLocation"`
	Distance      float64 `json:"distance" xml:"distance" form:"distance"`
	Duration      int64   `json:"duration" xml:"duration" form:"duration"`
	CO2           float64 `json:"co2" xml:"co2" form:"co2"`
	TransportType string  `json:"transportType" xml:"transportType" form:"transportType"`
}
