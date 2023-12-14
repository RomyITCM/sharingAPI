package entities

type DataCustomerAddress struct {
	Id            string `json:"id"`
	AliasName     string `json:"alias_name"`
	AliasNameFull string `json:"alias_name_full"`
	StreetAddress string `json:"street_address"`
	ProvinceId    string `json:"province_id"`
	ProvinceName  string `json:"province_name"`
	CityId        string `json:"city_id"`
	CityName      string `json:"city_name"`
	DistrictId    string `json:"district_id"`
	DistrictName  string `json:"district_name"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	StoreFreezer  string `json:"store_freezer"`
	BenFreezer    string `json:"ben_freezer"`
}

type CustomerAddressFreezer struct {
	FreezerId       string `json:"freezer_id"`
	OutletId        string `json:"outlet_id"`
	CustomerId      string `json:"customer_id"`
	FreezerOrigin   string `json:"freezer_origin"`
	FreezerType     string `json:"freezer_type"`
	RequestedAmount string `json:"requested_amount"`
}

type CustomerAddressId struct {
	CustomerAddressId string `json:"customer_address_id"`
}

type CustomerAddressDetail struct {
	Id            string `json:"id"`
	AliasName     string `json:"alias_name"`
	ALiasNameFull string `json:"alias_name_full"`
	CustomerNo    string `json:"customer_no"`
	StoreArea     string `json:"store_area"`
	StoreImg      string `json:"store_img"`
	StreetAddress string `json:"street_address"`
	ProvinceId    string `json:"province_id"`
	Province      string `json:"province"`
	CityId        string `json:"city_id"`
	City          string `json:"city"`
	DistrictId    string `json:"district_id"`
	District      string `json:"district"`
	SubdistrictId string `json:"subdistrict_id"`
	Subdistrict   string `json:"subdistrict"`
	ZipCode       string `json:"zip_code"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	PhoneNo       string `json:"phone_no"`
	MobilePhoneNo string `json:"mobile_phone_no"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	Salesman      string `json:"salesman"`
	RegionId      string `json:"region_id"`
	AreaId        string `json:"area_id"`
	ZoneId        string `json:"zone_id"`
	BillTo        string `json:"bill_to"`
	ShipSchedule  string `json:"ship_schedule"`
}

type CustomerAddressBillToByCustomer struct {
	Id          string `json:"id"`
	BillingName string `json:"billing_name"`
}
