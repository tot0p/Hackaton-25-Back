package APIInput

type Login struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Device   string `json:"device" xml:"device" form:"device"`
}
