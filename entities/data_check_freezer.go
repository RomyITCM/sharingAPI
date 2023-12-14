package entities

type DataCheckFreezerGetrows struct {
	TransNo            string `json:"trans_no"`
	SerialNo           string `json:"serial_no"`
	CustomerNo         string `json:"customer_no"`
	BillTo             string `json:"location_cust_bill_to"`
	ShipTo             string `json:"location_cust_ship_to"`
	ArticleNo          string `json:"article_no"`
	ArticleDescription string `json:"article_description"`
	FreezerAvailable   string `json:"freezer_available"`
	FreezerFrom        string `json:"freezer_from"`
	FreezerUse         string `json:"freezer_use"`
	FreezerLocation    string `json:"freezer_location"`
	FreezerChecked     string `json:"freezer_checked"`
}

type DataCheckFreezerGetrow struct {
	TransNo            string `json:"trans_no"`
	Mode               string `json:"mode"`
	ArticleDescription string `json:"article_description"`
	SerialNo           string `json:"serial_no"`
	ImagesType         string `json:"images_type"`
	ImagesName         string `json:"images_name"`
	FreezerAvailable   string `json:"freezer_available"`
	AvailableNote      string `json:"available_note"`
	Brand              string `json:"brand"`
	NoteBrand          string `json:"note_brand"`
	Type               string `json:"type"`
	Capacity           string `json:"capacity"`
	Status             string `json:"status"`
	NoteStatus         string `json:"note_status"`
	Location           string `json:"location"`
	NoteLocation       string `json:"note_location"`
	IceThickness       string `json:"ice_thickness"`
	Temperature        string `json:"temperature"`
	FreezerUse         string `json:"freezer_use"`
	NoteFreezerUse     string `json:"note_freezer_use"`
	CreatedBy          string `json:"created_by"`
	CreatedDate        string `json:"created_date"`
	CreatedByIP        string `json:"created_by_ip"`
}

type DataCheckFreezerInsert struct {
	CustomerNo      string `json:"customer_no"`
	BillTo          string `json:"bill_to"`
	ShipTo          string `json:"ship_to"`
	FrezerAvailable int    `json:"frezer_available"`
	CreatedBy       string `json:"created_by"`
	CreatedByIP     string `json:"created_by_ip"`

	AvailableNote      string `json:"available_note"`
	Mode               string `json:"mode"`
	ArticleDescription string `json:"article_description"`
	SerialNo           string `json:"serial_no"`
	SerialNoImages     []struct {
		Base64Image string `json:"base64Image"`
	} `json:"serial_no_images"`
	Brand          int    `json:"brand"`
	NoteBrand      string `json:"note_brand"`
	Type           string `json:"type"`
	Capacity       int    `json:"capacity"`
	Location       int    `json:"location"`
	NoteLocation   string `json:"note_location"`
	LocationImages []struct {
		Base64Image string `json:"base64Image"`
	} `json:"location_images"`
	Status       int    `json:"status"`
	NoteStatus   string `json:"note_status"`
	StatusImages []struct {
		Base64Image string `json:"base64Image"`
	} `json:"status_images"`
	IceThickness      int `json:"ice_thickness"`
	Temperature       int `json:"temperature"`
	TemperatureImages []struct {
		Base64Image string `json:"base64Image"`
	} `json:"temperature_images"`
	FreezerUse     string `json:"freezer_use"`
	NoteFreezerUse string `json:"note_freezer_use"`
	UsageImages    []struct {
		Base64Image string `json:"base64Image"`
	} `json:"usage_images"`
}

type DataCheckFreezerOutputInsertResult struct {
	DetailID string `json:"detail_id"`
}

type DataCheckFreezerOutputInsert struct {
	TransNo  string `json:"trans_no"`
	DetailID string `json:"detail_id"`
}

type DataCheckFreezerStatus struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type DataCheckFreezerInsertOutput struct {
	SerialNoImages    interface{} `json:"serial_no_images"`
	LocationImages    interface{} `json:"location_images"`
	StatusImages      interface{} `json:"status_images"`
	TemperatureImages interface{} `json:"temperature_images"`
	UsageImages       interface{} `json:"usage_images"`
}

type DataCheckFreezerInsertOutputURL struct {
	FileName   string `json:"file_name"`
	PreSignURL string `json:"presign_url"`
}
