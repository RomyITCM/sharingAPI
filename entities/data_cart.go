package entities

type DataCart struct {
	CartProducts   interface{} `json:"cart_products"`
	CartDeliveries interface{} `json:"cart_deliveries"`
	CartBills      interface{} `json:"cart_bills"`
	CartPromo      interface{} `json:"cart_promo"`
}

type CartProduct struct {
	ArticleDesc string `json:"article_desc"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
	Uom         string `json:"uom"`
	Total       string `json:"total"`
	Image       string `json:"image"`
}

type CartDelivery struct {
	VendorNo      string `json:"vendor_no"`
	VendorName    string `json:"vendor_name"`
	ShipTo        string `json:"ship_to"`
	ShipToAddress string `json:"ship_to_address"`
	PaymentTermId string `json:"payment_term_id"`
	PaymentTerm   string `json:"payment_term"`
}

type CartDeliveryEdit struct {
	CustPoNo      string `json:"cust_po_no"`
	PoDate        string `json:"po_date"`
	ShipTo        string `json:"ship_to"`
	ShipToAddress string `json:"ship_to_address"`
	PaymentTermId string `json:"payment_term_id"`
	PaymentTerm   string `json:"payment_term"`
	BillTo        string `json:"bill_to"`
	BillToAddress string `json:"bill_to_address"`
	Attachment    string `json:"attachment"`
	DeliveryFrom  string `json:"delivery_from"`
	CreatedDate   string `json:"created_date"`
	DelvDate      string `json:"delv_date"`
	PoExpDate     string `json:"po_exp_date"`
	CustType      string `json:"cust_type"`
}

type CartBill struct {
	TotalAmount string `json:"total_amount"`
	Disc        string `json:"discount"`
	Tax         string `json:"tax"`
	TotalBill   string `json:"total_bill"`
}

type CartPromo struct {
	PromoCode  string `json:"promo_code"`
	PromoName  string `json:"promo_name"`
	ResultDesc string `json:"result_desc"`
	AmountDisc string `json:"amount_disc"`
	IsPrimary  string `json:"is_primary"`
}

type CartPaymentTerm struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CartProcess struct {
	WHID        string `json:"wh_id"`
	CustomerNo  string `json:"customer_no"`
	BillTo      string `json:"bill_to"`
	ShipTo      string `json:"ship_to"`
	CustPoNo    string `json:"cust_po_no"`
	CustPoDate  string `json:"cust_po_date"`
	ExpPoDate   string `json:"exp_po_date"`
	DelvDate    string `json:"delv_date"`
	Attachment  string `json:"attachment"`
	PaymentTerm string `json:"payment_term"`
	CreatedBy   string `json:"created_by"`
	CreatedByIp string `json:"created_by_ip"`
}

type CartProcessEdit struct {
	TransNo     string `json:"trans_no"`
	CustPoNo    string `json:"cust_po_no"`
	DelvDate    string `json:"delv_date"`
	ExpPoDate   string `json:"exp_po_date"`
	EditFile    bool   `json:"edit_file"`
	Attachment  string `json:"attachment"`
	CreatedBy   string `json:"created_by"`
	CreatedByIp string `json:"created_by_ip"`
}

type CartResult struct {
	TransNo      string `json:"trans_no"`
	PreSignedURL string `json:"pre_signed_url"`
	FileName     string `json:"file_name"`
	MsgType      string `json:"msg_type"`
	MsgError     string `json:"msg_error"`
}
