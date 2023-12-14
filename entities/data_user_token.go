package entities

type UserToken struct {
	UserId      string `json:"user_id"`
	Token       string `json:"token"`
	UserName    string `json:"user_name"`
	DeviceName  string `json:"device_name"`
	CreatedByIp string `json:"created_by_ip"`
}

type UserTokenLogout struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
