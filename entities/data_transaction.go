package entities

type DataTransaction struct {
	ShipTo    string `json:"ship_to"`
	ArticleNo string `json:"article_no"`
	Qty       string `json:"qty"`
	SalesMan  string `json:"sales_man"`
}

type QtyCart struct {
	QtyCart int `json:"qty_cart"`
}

type DataTransactionPromo struct {
	ShipTo      string `json:"ship_to"`
	SalesMan    string `json:"sales_man"`
	PaymentTerm string `json:"payment_term"`
}

type UpdateQtyCart struct {
	TransNo string           `json:"trans_no"`
	Detail  []DetailItemCart `json:"detail"`
}

type DetailItemCart struct {
	ArticleNo string `json:"article_no"`
	Qty       int    `json:"qty"`
}
