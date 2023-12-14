package entities

type InfoCustomerFreezer struct {
	CustomerId        string `json:"customer_id"`
	CustomerAddressId string `json:"customer_address_id"`
	FreezerOrigin     string `json:"freezer_origin"`
	FreezerType       string `json:"freezer_type"`
	RequestedAmount   string `json:"requested_amount"`
	CreatedBy         string `json:"created_by"`
	CreatedByIp       string `json:"created_by_ip"`
}
