package APIOutput

type Login struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Device   string `json:"device"`
}
