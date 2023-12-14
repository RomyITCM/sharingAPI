package entities

type DataPriceZone struct {
	ZoneId   string `json:"zone_id"`
	ZoneName string `json:"zone_name"`
}

type DataArticlePrice struct {
	ZoneId             string `json:"zone_id"`
	ArticleNo          string `json:"article_no"`
	ArticleDescription string `json:"article_description"`
	Uom                string `json:"uom"`
	SalesPrice         string `json:"sales_price"`
}
