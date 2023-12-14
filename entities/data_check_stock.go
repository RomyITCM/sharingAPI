package entities

type DataCheckStockGetrows struct {
	TransNo            string `json:"trans_no"`
	ArticleNumber      string `json:"article_number"`
	UrlImage           string `json:"url_image"`
	ArticleDescription string `json:"article_description"`
	SalesPrice         string `json:"sales_price"`
	Stock              string `json:"stock"`
	Checked            string `json:"checked"`
}

type DataCheckStockGetrow struct {
	TransNo             string `json:"trans_no"`
	ArticleNo           string `json:"article_no"`
	Type                string `json:"type"`
	StockAvailable      string `json:"stock_available"`
	Price               string `json:"price"`
	PromoAvailable      string `json:"promo_available"`
	MediaPromoAvailable string `json:"media_promo_available"`
	NotePromo           string `json:"note_promo"`
	StartDatePromo      string `json:"start_date_promo"`
	EndDatePromo        string `json:"end_date_promo"`
	ImageType           string `json:"image_type"`
	ImagesName          string `json:"images_name"`
	GoodStock           string `json:"good_stock"`
	BadStock            string `json:"bad_stock"`
	ExpiredStock        string `json:"expired_stock"`
	TotalStock          string `json:"total_stock"`
}

type DataCheckStockDetailGetrow struct {
	ID            string      `json:"id"`
	TransNo       string      `json:"trans_no"`
	ArticleNumber string      `json:"article_number"`
	ExpiredDate   string      `json:"expired_date"`
	QtyStock      string      `json:"qty_stock"`
	Condition     string      `json:"condition"`
	ImageStock    interface{} `json:"image_stocks"`
	// ImageStock      []DataCheckStockDetailImageGetrow
	NoteStock       string `json:"note_stock"`
	SalesReturnNo   string `json:"sales_return_no"`
	NoteSalesReturn string `json:"note_sales_return"`
}

type DataCheckStockDetailImageGetrow struct {
	ImageStock string `json:"image_stock"`
}

type DataCheckStockCompetitorGetrows struct {
	ArticleNumber      string `json:"article_number"`
	UrlImage           string `json:"url_image"`
	ArticleDescription string `json:"article_description"`
	SalesPrice         string `json:"sales_price"`
	Stock              string `json:"stock"`
	Checked            string `json:"checked"`
}

type DataCheckStockInsert struct {
	CustomerNo     string `json:"customer_no"`
	BillTo         string `json:"bill_to"`
	ShipTo         string `json:"ship_to"`
	StockAvailable int    `json:"stock_available"`
	CreatedBy      string `json:"created_by"`
	CreatedByIP    string `json:"created_by_ip"`
	ArticleNo      string `json:"article_no"`
	Type           string `json:"type"`
	Price          string `json:"price"`
	PriceImages    []struct {
		Base64Image string `json:"base64Image"`
	} `json:"price_images"`
	PromoAvailable      int `json:"promo_available"`
	MediaPromoAvailable int `json:"media_promo_available"`
	PromoImages         []struct {
		Base64Image string `json:"base64Image"`
	} `json:"promo_images"`
	NotePromo      string `json:"note_promo"`
	StartDatePromo string `json:"start_date_promo"`
	EndDatePromo   string `json:"end_date_promo"`
	DetailSKU      []struct {
		ArticleNo       string `json:"article_no"`
		ExpDate         string `json:"exp_date"`
		Stock           string `json:"stock"`
		Condition       string `json:"condition"`
		NoteStock       string `json:"note_stock"`
		SalesReturnNo   string `json:"sales_return_no"`
		NoteSalesReturn string `json:"note_sales_return"`
		BadStockImages  []struct {
			Base64Image string `json:"base64Image"`
		} `json:"bad_stock_image"`
	} `json:"detail_sku"`
}

type DataCheckStockOutputInsert struct {
	TransNo string `json:"trans_no"`
	ReffNo  string `json:"reff_no"`
}

type DataCheckStockDetailOutputInsert struct {
	DetailID string `json:"detail_id"`
}

type DataCheckStockInsertOutput struct {
	PriceImages      interface{} `json:"price_images"`
	MediaPromoImages interface{} `json:"media_promo_images"`
	ConditionImages  interface{} `json:"condition_images"`
}

type DataCheckStockInsertOutputURL struct {
	FileName   string `json:"file_name"`
	PreSignURL string `json:"presign_url"`
}
