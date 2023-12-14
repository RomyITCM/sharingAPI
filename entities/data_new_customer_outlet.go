package entities

type DataNewCustomerOutlet struct {
	OutletId     string `json:"outlet_id"`
	OutletName   string `json:"outlet_name"`
	Address      string `json:"address"`
	ProvinceId   string `json:"province_id"`
	CityId       string `json:"city_id"`
	DistrictId   string `json:"district_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	StoreFreezer string `json:"store_freezer"`
	BenFreezer   string `json:"ben_freezer"`
}

type DataNewCustomerOutletDetail struct {
	OutletId        string `json:"outlet_id"`
	OutletCode      string `json:"outlet_code"`
	OutletName      string `json:"outlet_name"`
	CustomerId      string `json:"customer_id"`
	StoreArea       string `json:"store_area"`
	ShipSchedule    string `json:"ship_schedule"`
	StoreImg        string `json:"store_img"`
	Address         string `json:"address"`
	ProvinceId      string `json:"province_id"`
	ProvinceName    string `json:"province_name"`
	CityId          string `json:"city_id"`
	CityName        string `json:"city_name"`
	DistrictId      string `json:"district_id"`
	DistrictName    string `json:"district_name"`
	SubdistrictId   string `json:"subdistrict_id"`
	SubdistrictName string `json:"subdistrict_name"`
	Zipcode         string `json:"zipcode"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	PhoneNo         string `json:"phone_no"`
	CellphoneNo     string `json:"cellphone_no"`
	Email           string `json:"email"`
	Salesman        string `json:"salesman"`
	RegionId        string `json:"region_id"`
	RegionName      string `json:"region_name"`
	AreaId          string `json:"area_id"`
	AreaName        string `json:"area_name"`
	ZoneId          string `json:"zone_id"`
	ZoneName        string `json:"zone_name"`
	Status          string `json:"status"`
	ApprovalStatus  string `json:"approval_status"`
}

type OutletId struct {
	OutletId string `json:"outlet_id"`
}

type InfoNewCustomerOutlet struct {
	OutletCode   string `json:"outlet_code"`
	OutletName   string `json:"outlet_name"`
	CustomerId   string `json:"customer_id"`
	StoreArea    string `json:"store_area"`
	ShipSchedule string `json:"ship_schedule"`
	StoreImg     string `json:"store_img"`
	Address      string `json:"address"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Subdistrict  string `json:"subdistrict"`
	Zipcode      string `json:"zipcode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	PhoneNo      string `json:"phone_no"`
	CellphoneNo  string `json:"cellphone_no"`
	Email        string `json:"email"`
	Salesman     string `json:"salesman"`
	RegionId     string `json:"region_id"`
	AreaId       string `json:"area_id"`
	ZoneId       string `json:"zone_id"`
	BillTo       string `json:"bill_to"`
	CreatedBy    string `json:"created_by"`
	CreatedByIp  string `json:"created_by_ip"`
}

type InfoNewCustomerOutletUpdate struct {
	OutletId     string `json:"outlet_id"`
	OutletCode   string `json:"outlet_code"`
	OutletName   string `json:"outlet_name"`
	StoreArea    string `json:"store_area"`
	ShipSchedule string `json:"ship_schedule"`
	StoreImg     string `json:"store_img"`
	Address      string `json:"address"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Subdistrict  string `json:"subdistrict"`
	Zipcode      string `json:"zipcode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	PhoneNo      string `json:"phone_no"`
	CellphoneNo  string `json:"cellphone_no"`
	Email        string `json:"email"`
	Salesman     string `json:"salesman"`
	RegionId     string `json:"region_id"`
	AreaId       string `json:"area_id"`
	ZoneId       string `json:"zone_id"`
	Status       string `json:"status"`
	CreatedBy    string `json:"created_by"`
	CreatedByIp  string `json:"created_by_ip"`
}

type DataNeedApproveOutlet struct {
	OutletId     string `json:"outlet_id"`
	OutletName   string `json:"outlet_name"`
	Address      string `json:"address"`
	ProvinceId   string `json:"province_id"`
	CityId       string `json:"city_id"`
	DistrictId   string `json:"district_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	StoreFreezer string `json:"store_freezer"`
	BenFreezer   string `json:"ben_freezer"`
}
