package entities

type InfoCustomerAsset struct {
	ShipTo          string `json:"ship_to"`
	FreezerName     string `json:"freezer_name"`
	Location        string `json:"location"`
	Usage           string `json:"usage"`
	CreateBy        string `json:"create_by"`
	CreateDate      string `json:"create_date"`
	CreateIpAddress string `json:"create_ip_address"`
}
