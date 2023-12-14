package entities

type DataProduct struct {
	ArticleNumber      string `json:"article_number"`
	UrlImage           string `json:"url_image"`
	ArticleDescription string `json:"article_description"`
	SalesPrice         string `json:"sales_price"`
	Stock              string `json:"stock"`
	QtyBook            string `json:"qty_book"`
}
