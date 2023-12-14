package entities

type DataCustomer struct {
	CustomerNo       string `json:"customer_no"`
	CustomerName     string `json:"customer_name"`
	CustomerAddress  string `json:"customer_address"`
	CustomerType     string `json:"customer_type"`
	CustomerTypeDesc string `json:"customer_type_desc"`
	CustomerStatus   bool   `json:"customer_status"`
}

type DataCustomerBillTo struct {
	BillToId      int    `json:"bill_to_id"`
	BillToName    string `json:"bill_to_name"`
	BillToAddress string `json:"bill_to_address"`
}

type DataCustomerShipTo struct {
	ShipToId      int    `json:"ship_to_id"`
	ShipToName    string `json:"ship_to_name"`
	ShipToAddress string `json:"ship_to_address"`
	PriceZone     string `json:"price_zone"`
	SiteId        string `json:"site_id"`
	SiteName      string `json:"site_name"`
	SiteAddress   string `json:"site_address"`
	ShipToStatus  bool   `json:"ship_to_status"`
	TripZone      string `json:"trip_zone"`
}

type DataCustomerPromoActive struct {
	PromoCode      string `json:"promo_code"`
	PromoName      string `json:"promo_name"`
	PromoDesc      string `json:"promo_desc"`
	StartPromoDate string `json:"start_promo_date"`
	EndPromoDate   string `json:"end_promo_date"`
}

type DataApprovedCustomer struct {
	CustomerNo    string `json:"customer_no"`
	Status        string `json:"status"`
	CustomerName  string `json:"customer_name"`
	CustomerType  string `json:"customer_type"`
	VatRegNo      string `json:"vat_reg_no"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	District      string `json:"district"`
	OutletCount   string `json:"outlet_count"`
	PicCount      string `json:"pic_count"`
}

type CustomerNo struct {
	CustomerNo string `json:"customer_no"`
}

type DataApprovedCustomerDetail struct {
	CustomerNo             string `json:"customer_no"`
	CustomerName           string `json:"customer_name"`
	CustomerType           string `json:"customer_type"`
	Type                   string `json:"type"`
	EstablishedIn          string `json:"established_in"`
	NoOfEmployee           string `json:"no_of_employee"`
	AnnualSales            string `json:"annual_sales"`
	PhoneNo                string `json:"phone_no"`
	Fax                    string `json:"fax"`
	Email                  string `json:"email"`
	Website                string `json:"website"`
	StreetAddress          string `json:"street_address"`
	ProvinceId             string `json:"province_id"`
	ProvinceName           string `json:"province_name"`
	CityId                 string `json:"city_id"`
	CityName               string `json:"city_name"`
	DistrictId             string `json:"district_id"`
	DistrictName           string `json:"district_name"`
	SubdistrictId          string `json:"subdistrict_id"`
	SubdistrictName        string `json:"subdistrict_name"`
	PostalCode             string `json:"postal_code"`
	TaxStatus              string `json:"tax_status"`
	TaxMandatory           string `json:"tax_mandatory"`
	VatRegNo               string `json:"vat_reg_no"`
	NpwpName               string `json:"npwp_name"`
	VatAddress             string `json:"vat_address"`
	VatImg                 string `json:"vat_img"`
	AliasName              string `json:"alias_name"`
	AliasNameFull          string `json:"alias_name_full"`
	BillingAddress         string `json:"billing_address"`
	BillingProvince        string `json:"billing_province"`
	BillingProvinceName    string `json:"billing_province_name"`
	BillingCity            string `json:"billing_city"`
	BillingCityName        string `json:"billing_city_name"`
	BillingDistrict        string `json:"billing_district"`
	BillingDistrctName     string `json:"billing_district_name"`
	BillingSubdistrict     string `json:"billing_subdistrict"`
	BillingSubdistrictName string `json:"billing_subdistrict_name"`
	BillingPostal          string `json:"billing_postal"`
	InvoiceMethod          string `json:"invoice_method"`
	PaymentTermsId         string `json:"payment_terms_id"`
	InvoiceDocument        string `json:"invoice_document"`
	ReturnStatus           string `json:"return_status"`
	InvoiceSchedule        string `json:"invoice_schedule"`
	Status                 string `json:"status"`
}
