package APIOutput

// Login represents the output for logging in
type Login struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Device   string `json:"device"`
}
