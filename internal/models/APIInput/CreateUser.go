package APIInput

// CreateUser represents the input for creating a user
type CreateUser struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
