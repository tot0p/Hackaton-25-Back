package DBModels

// User represents a user
type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
