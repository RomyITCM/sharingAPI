package entities

type DataDrafts struct {
	ShipTo           string `json:"ship_to"`
	TotalPrice       string `json:"total_price"`
	ArticleDesc      string `json:"article_desc"`
	Qty              string `json:"qty"`
	Note             string `json:"note"`
	UrlImage         string `json:"url_image"`
	CustomerNo       string `json:"customer_no"`
	CustomerName     string `json:"customer_name"`
	CustomerAddress  string `json:"customer_address"`
	CustomerType     string `json:"customer_type"`
	CustomerStatus   bool   `json:"customer_status"`
	CustomerTypeDesc string `json:"customer_type_desc"`
	ShipToId         string `json:"ship_to_id"`
	ShipToAddress    string `json:"ship_to_address"`
	PriceZone        string `json:"price_zone"`
	SiteId           string `json:"site_id"`
	SiteName         string `json:"site_name"`
	SiteAddress      string `json:"site_address"`
	ShipToStatus     bool   `json:"ship_to_status"`
	TripZone         string `json:"trip_zone"`
}
