package entities

type DataNotification struct {
	MessageTitle string `json:"message_title"`
	MessageBody  string `json:"message_body"`
	TransNo      string `json:"trans_no"`
	Token        string `json:"token"`
}

type DataUserNotification struct {
	TransNo      string `json:"trans_no"`
	MessageTitle string `json:"message_tittle"`
	MessageBody  string `json:"message_body"`
	IsRead       bool   `json:"is_read"`
	Type         string `json:"type"`
	Salesman     string `json:"salesman"`
}

type NotificationDelete struct {
	UserId  string `json:"user_id"`
	TransNo string `json:"trans_no"`
}

type NotificationRead struct {
	UserId   string `json:"user_id"`
	TransNo  string `json:"trans_no"`
	ReaderIp string `json:"read_by_ip"`
}

type NotificationSend struct {
	TransNo string `json:"trans_no"`
}
