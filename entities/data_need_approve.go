package entities

type DataNeedApproveList struct {
	TransNo     string `json:"trans_no"`
	TransDate   string `json:"trans_date"`
	ShipTo      string `json:"ship_to"`
	ReqDelvDate string `json:"req_delv_date"`
	TotalPrice  string `json:"total_price"`
	ArticleDesc string `json:"article_desc"`
	Qty         string `json:"qty"`
	Note        string `json:"note"`
	UrlImage    string `json:"url_image"`
	UserName    string `json:"user_name"`
}

type DataNeedApprove struct {
	TransNo      string `json:"trans_no"`
	NextStatus   string `json:"next_status"`
	CreatedBy    string `json:"created_by"`
	CreaterdByIp string `json:"creted_by_ip"`
}

type DataNeedApproveReject struct {
	TransNo     string `json:"trans_no"`
	Reason      string `json:"reason"`
	CancelledBy string `json:"cancelled_by"`
}

type DataSendEmail struct {
	Message string `json:"message"`
}

type DataUserApprove struct {
	UserId      string `json:"user_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type DataCustomerNumbers struct {
	CustomerRequestNo string `json:"customer_request_no"`
	CustomerNo        string `json:"customer_no"`
}
