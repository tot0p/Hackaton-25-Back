package APIInput

// Login represents the input for logging in
type Login struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
