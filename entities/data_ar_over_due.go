package entities

// type ParamCustArOverDueGetrows struct {
// 	Search string `json:"search"`
// }

// type ParamArOverDueGetrows struct {
// 	CustomerNo string `json:"customer_no"`
// 	BillTo     string `json:"bill_to"`
// 	BankCode   string `json:"bank_code"`
// 	Search     string `json:"search"`
// }

type DataCustArOverDueGetrows struct {
	CustomerNo             string `json:"customer_no"`
	CustomerName           string `json:"customer_name"`
	BillTo                 string `json:"bill_to"`
	BillToName             string `json:"bill_to_name"`
	Dpp                    string `json:"dpp"`
	Ppn                    string `json:"ppn"`
	TotalAmount            string `json:"total_amount"`
	PaymentAmount          string `json:"payment_amount"`
	TotalAmountOutstanding string `json:"total_amount_outstanding"`
	OneToFourteen          string `json:"onetofourteen"`
	FifteenToTwentyOne     string `json:"fifteentotwentyone"`
	TwentyTwoToThirty      string `json:"twentytwotothirty"`
	ThirtyoneToSixty       string `json:"thirtyonetosixty"`
	SixtyoneToninety       string `json:"sixtyonetoninety"`
	MoreThanNinety         string `json:"morethanninety"`
}

type DataArOverDueGetrows struct {
	TransNo                  string `json:"trans_no"`
	DocNo                    string `json:"doc_no"`
	TransDate                string `json:"trans_date"`
	JvDate                   string `json:"jv_date"`
	CustomerName             string `json:"customer_name"`
	BillTo                   string `json:"bill_to"`
	BillToName               string `json:"bill_to_name"`
	ShipTo                   string `json:"ship_to"`
	ShipToName               string `json:"ship_to_name"`
	TotalInvoice             string `json:"total_invoice"`
	Ppn                      string `json:"ppn"`
	TotalAmount              string `json:"total_amount"`
	PaymentAmount            string `json:"payment_amount"`
	TotalAmountOutstanding   string `json:"total_amount_outstanding"`
	DueDate                  string `json:"duedate"`
	JatuhTempo               string `json:"jatuh_tempo"`
	OneToFourteenDetail      string `json:"onetofourteen"`
	FifteenToTwentyOneDetail string `json:"fifteentotwentyone"`
	TwentyTwoToThirtyDetail  string `json:"twentytwotothirty"`
	ThirtyoneToSixtyDetail   string `json:"thirtyonetosixty"`
	SixtyoneToninetyDetail   string `json:"sixtyonetoninety"`
	MoreThanNinetyDetail     string `json:"morethanninety"`
}

type DataCustArOverdueGetrows struct {
	DocNo       string `json:"doc_no"`
	CustPONo    string `json:"cust_po_no"`
	Overdue     string `json:"overdue"`
	Outstanding string `json:"outstanding"`
}

type DataCustArOverdueGetrow struct {
	Customer       string `json:"customer"`
	CustomerName   string `json:"customer_name"`
	BillTo         string `json:"bill_to"`
	ShipTo         string `json:"ship_to"`
	TotalAROverdue string `json:"total_ar_overdue"`
}

type DataCustArOverdue struct {
	AROverdueHeader interface{} `json:"ar_overdue_header"`
	AROverdueDetail interface{} `json:"ar_overdue_detail"`
}
