package entities

type DataHistoryTransaction struct {
	TransNo     string `json:"trans_no"`
	TransDate   string `json:"trans_date"`
	Status      string `json:"status"`
	ShipTo      string `json:"ship_to"`
	PODate      string `json:"po_date"`
	TotalPrice  string `json:"total_price"`
	ArticleDesc string `json:"article_desc"`
	Qty         string `json:"qty"`
	Note        string `json:"note"`
	UrlImage    string `json:"url_image"`
}

type DataHistoryTransactionDetail struct {
	DataHistoryTransactionProducts interface{} `json:"history_transaction_products"`
	DataHistoryTransactionDelivery interface{} `json:"history_transaction_delivery"`
}

type DataHistoryTransactionProducts struct {
	ArticleDesc string `json:"article_desc"`
	Qty         string `json:"qty"`
	Uom         string `json:"uom"`
	UnitPrice   string `json:"unit_price"`
	TotalPrice  string `json:"total_price"`
	Image       string `json:"url_image"`
	WhId        string `json:"wh_id"`
	SiteName    string `json:"site_name"`
}

type DataHistoryTransactionDelivery struct {
	TransDate       string `json:"trans_date"`
	Salesman        string `json:"salesman"`
	Status          string `json:"status"`
	RequestRecvDate string `json:"request_recv_date"`
	ReceiveDate     string `json:"receive_date"`
	ShipTo          string `json:"ship_to"`
	ShipToAddress   string `json:"ship_to_address"`
	CustomerPoNo    string `json:"cust_po_no"`
	AttacmentPO     string `json:"attachment_po"`
	PaymentTerm     string `json:"payment_term"`
	PaymentDue      string `json:"payment_due"`
	PoExpDate       string `json:"po_exp_date"`
	BillTo          string `json:"bill_to"`
	BillToAddress   string `json:"bill_to_address"`
	Amount          string `json:"amount"`
	Disc            string `json:"disc"`
	Tax             string `json:"tax"`
	TotalPrice      string `json:"total_price"`
	Reason          string `json:"reason"`
	Attachment      string `json:"attachment"`
}
