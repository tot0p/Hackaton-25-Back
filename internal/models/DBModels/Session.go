package DBModels

// Session represents a session in the database
type Session struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"userUUID"`
	Device   string `json:"device"`
}
