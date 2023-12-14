package entities

type DataTransactionSuccess struct {
	DataTransactionSuccessHeader interface{} `json:"header"`
	DataTransactionSuccessDetail interface{} `json:"detail"`
	DataTransactionSuccessPromo  interface{} `json:"promo"`
}

type DataTransactionSuccessHeader struct {
	TransNo       string `json:"trans_no"`
	TransDate     string `json:"trans_date"`
	CustPoNo      string `json:"cust_po_no"`
	CustPoDate    string `json:"cust_po_date"`
	PoExpDate     string `json:"po_exp_date"`
	ReqDate       string `json:"req_date"`
	PaymentTerm   string `json:"payment_term"`
	Amount        string `json:"amount"`
	Disc          string `json:"disc"`
	Vat           string `json:"vat"`
	TotalAmount   string `json:"total_amount"`
	BillTo        string `json:"bill_to"`
	BillToAddress string `json:"bill_to_address"`
	ShipTo        string `json:"ship_to"`
	ShipToAddress string `json:"ship_to_address"`
	TotalCarton   string `json:"total_carton"`
	Status        bool   `json:"status"`
	Attachment    string `json:"attachment"`
	CustType      string `json:"cust_type"`
}

type DataTransactionSuccessDetail struct {
	ArticleDesc string `json:"article_desc"`
	Qty         string `json:"qty"`
	Uom         string `json:"uom"`
	Price       string `json:"price"`
	TotalPrice  string `json:"total_price"`
}

type DataTransactionSuccessPromo struct {
	PromoCode  string `json:"promo_code"`
	PromoName  string `json:"promo_name"`
	ResultDesc string `json:"result_desc"`
	AmountDisc string `json:"amount_disc"`
}
