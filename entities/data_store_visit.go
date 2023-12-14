package entities

type DataStartEndDay struct {
	Type             string `json:"type"`
	Pic              string `json:"pic"`
	VisitDate        string `json:"visit_date"`
	Kilometer        int    `json:"kilometer"`
	PicKilometer     string `json:"pic_kilometer"`
	Vehicle          string `json:"vehicle"`
	PicVehiclePlate  string `json:"pic_vehicle_plate"`
	Address          string `json:"address"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
	CreatedBy        string `json:"created_by"`
	CreatedByIp      string `json:"created_by_ip"`
}

type DataStartDayGetrow struct {
	Id               string `json:"id"`
	Pic              string `json:"pic"`
	VisitDate        string `json:"visit_date"`
	Kilometer        int    `json:"kilometer"`
	PicKilometer     string `json:"pic_kilometer"`
	Vehicle          string `json:"vehicle"`
	PicVehiclePlate  string `json:"pic_vehicle_plate"`
	Address          string `json:"address"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
}

type DataEndDay struct {
	Pic              string `json:"pic"`
	VisitDate        string `json:"visit_date"`
	Kilometer        int    `json:"kilometer"`
	PicKilometer     string `json:"pic_kilometer"`
	Vehicle          string `json:"vehicle"`
	PicVehiclePlate  string `json:"pic_vehicle_plate"`
	Address          string `json:"address"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
	CreatedBy        string `json:"created_by"`
	CreatedByIp      string `json:"created_by_ip"`
}

type DataStoreVisit struct {
	Id                    int    `json:"id"`
	DetailID              int    `json:"detail_id"`
	StartDayAddress       string `json:"start_day_address"`
	StartDayTime          string `json:"start_day_time"`
	EndDayAddress         string `json:"end_day_address"`
	EndDayTime            string `json:"end_day_time"`
	ShipTo                string `json:"ship_to"`
	CustomerNo            string `json:"customer_no"`
	CustomerName          string `json:"customer_name"`
	CustomerAddress       string `json:"customer_address"`
	CustomerType          string `json:"customer_type"`
	CustomerStatus        bool   `json:"customer_status"`
	CustomerTypeDesc      string `json:"customer_type_desc"`
	StoreAdddress         string `json:"store_address"`
	StoreName             string `json:"store_name"`
	ShipToName            string `json:"ship_to_name"`
	CheckInTime           string `json:"check_in_time"`
	CheckInAddress        string `json:"check_in_address"`
	CheckOutTime          string `json:"check_out_time"`
	CheckOutAddress       string `json:"check_out_address"`
	ShipToID              string `json:"ship_to_id"`
	PriceZone             string `json:"price_zone"`
	SiteId                string `json:"site_id"`
	SiteName              string `json:"site_name"`
	SiteAddress           string `json:"site_address"`
	TripZone              string `json:"trip_zone"`
	ShipToStatus          bool   `json:"ship_to_status"`
	Latitude              string `json:"latitude"`
	Longitude             string `json:"longitude"`
	TotalEfectiveCall     string `json:"total_effective_call"`
	TotalSales            string `json:"total_sales"`
	TotalAvailableStore   string `json:"total_available_store"`
	TotalUnAvailableStore string `json:"total_unavailable_store"`
}

type DataCheckInOut struct {
	VisitID          string `json:"visit_id"`
	Type             string `json:"type"`
	CheckTime        string `json:"check_time"`
	CustomerNo       string `json:"customer_no"`
	ShipTo           string `json:"ship_to"`
	Address          string `json:"address"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
	CreatedBy        string `json:"created_by"`
	CreatedByIp      string `json:"created_by_ip"`
	Images           []struct {
		Base64Image string `json:"base64Image"`
	} `json:"images"`
}

type DataCheckInOutID struct {
	DetailID string `json:"detail_id"`
}

type DataCheckInOutGetrow struct {
	ID                string `json:"id"`
	VisitID           string `json:"visit_id"`
	CheckInTime       string `json:"check_in_time"`
	CheckInAddress    string `json:"check_in_address"`
	CheckInLatitude   string `json:"check_in_latitude"`
	CheckInLongitude  string `json:"check_in_longitude"`
	CheckOutTime      string `json:"check_out_time"`
	CheckOutAddress   string `json:"check_out_address"`
	CheckOutLatitude  string `json:"check_out_latitude"`
	CheckOutLongitude string `json:"check_out_longitude"`
	PicCheckInOut     string `json:"pic_check_in_out"`
}

type DataStoreList struct {
	ShipToName       string `json:"ship_to_name"`
	ShipToNo         string `json:"ship_to_no"`
	CustomerName     string `json:"customer_name"`
	CustomerNo       string `json:"customer_no"`
	CustomerAddress  string `json:"customer_address"`
	CustomerType     string `json:"customer_type"`
	CustomerStatus   bool   `json:"customer_status"`
	CustomerTypeDesc string `json:"customer_type_desc"`
	BillToName       string `json:"bill_to_name"`
	BillToNo         string `json:"bill_to_no"`
	ShipToAddress    string `json:"ship_to_address"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	LastReport       string `json:"last_report"`
	Status           string `json:"status"`
	PriceZone        string `json:"price_zone"`
	SiteId           string `json:"site_id"`
	SiteName         string `json:"site_name"`
	SiteAddress      string `json:"site_address"`
	TripZone         string `json:"trip_zone"`
	ShipToStatus     bool   `json:"ship_to_status"`
	LastPic          string `json:"last_pic"`
	LastCheckIn      string `json:"last_check_in"`
}

type DataStoreVisitMenuValidation struct {
	IsFreezerCheckedAll string `json:"is_freezer_checked_all"`
	IsStockCheckedAll   string `json:"is_stock_checked_all"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
}

type DataStoreVisitUpdateCoordinate struct {
	CustomerNo string `json:"customer_no"`
	ShipTo     string `json:"ship_to"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	CreatedBy  string `json:"created_by"`
}

type DataStoreVisitOutput struct {
	StartEndDayImages interface{} `json:"start_end_day_images"`
}

type DataStoreVisitOutputURL struct {
	FileName   string `json:"file_name"`
	PreSignURL string `json:"presign_url"`
}
