package entities

type DataNewCustomerFreezer struct {
	FreezerId         string `json:"freezer_id"`
	OutletId          string `json:"outlet_id"`
	CustomerId        string `json:"customer_id"`
	CustomerAddressId string `json:"customer_address_id"`
	FreezerOrigin     string `json:"freezer_origin"`
	FreezerType       string `json:"freezer_type"`
	RequestedAmount   string `json:"requested_amount"`
}

type InfoNewCustomerFreezer struct {
	CustomerId      string `json:"customer_id"`
	OutletId        string `json:"outlet_id"`
	FreezerOrigin   string `json:"freezer_origin"`
	FreezerType     string `json:"freezer_type"`
	RequestedAmount string `json:"requested_amount"`
	CreatedBy       string `json:"created_by"`
	CreatedByIp     string `json:"created_by_ip"`
}

type InfoNewCustomerFreezerUpdate struct {
	FreezerId       string `json:"freezer_id"`
	FreezerOrigin   string `json:"freezer_origin"`
	FreezerType     string `json:"freezer_type"`
	RequestedAmount string `json:"requested_amount"`
	CreatedBy       string `json:"created_by"`
	CreatedByIp     string `json:"created_by_ip"`
}

type FreezerId struct {
	FreezerId string `json:"freezer_id"`
}

type DataBenfarmFreezer struct {
	FreezerId       string `json:"freezer_id"`
	CustomerId      string `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	OutletId        string `json:"outlet_id"`
	OutletName      string `json:"outlet_name"`
	FreezerOrigin   string `json:"freezer_origin"`
	FreezerType     string `json:"freezer_type"`
	RequestedAmount string `json:"requested_amount"`
}

type DataSendEmailBenfarm struct {
	Recipient    string `json:"recipient"`
	CustomerId   string `json:"customer_id"`
	OutletId     string `json:"outlet_id"`
	CustomerName string `json:"customer_name"`
	HtmlBody     string `json:"html_body"`
}
