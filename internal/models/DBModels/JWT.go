package DBModels

type JWT struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"userUUID"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
