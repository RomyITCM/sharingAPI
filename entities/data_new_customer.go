package entities

type DataNewCustomerRequest struct {
	CustomerRequestNo string `json:"customer_request_no"`
	Status            string `json:"status"`
	CustomerName      string `json:"customer_name"`
	CustomerType      string `json:"customer_type"`
	VatRegNo          string `json:"vat_reg_no"`
	StreetAddress     string `json:"street_address"`
	CityId            string `json:"city_id"`
	DistrictId        string `json:"district_id"`
	OutletCount       string `json:"outlet_count"`
	PicCount          string `json:"pic_count"`
}

type DataPreviewCustomer struct {
	CustomerRequestNo      string `json:"customer_request_no"`
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
	BillingCode            string `json:"billing_code"`
	BillingName            string `json:"billing_name"`
	BillingAddress         string `json:"billing_address"`
	BillingProvince        string `json:"billing_province"`
	BillingProvinceName    string `json:"billing_province_name"`
	BillingCity            string `json:"billing_city"`
	BillingCityName        string `json:"billing_city_name"`
	BillingDistrict        string `json:"billing_district"`
	BillingDistrictName    string `json:"billing_district_name"`
	BillingSubdistrict     string `json:"billing_subdistrict"`
	BillingSubdistrictName string `json:"billing_subdistrict_name"`
	BillingPostal          string `json:"billing_postal"`
	BillingMethod          string `json:"billing_method"`
	PaymentTermsId         string `json:"payment_terms_id"`
	BillingDocuments       string `json:"billing_documents"`
	ReturAvailable         string `json:"retur_available"`
	FactureSchedule        string `json:"facture_schedule"`
	Status                 string `json:"status"`
	ApprovalStatus         string `json:"approval_status"`
}

type CustomerRequestNo struct {
	CustomerRequestNo string `json:"customer_request_no"`
}

type InfoNewCustomerRequest struct {
	CustomerName       string `json:"customer_name"`
	CustomerType       string `json:"customer_type"`
	YearEstablished    string `json:"year_established"`
	EmployeeAmount     string `json:"employee_amount"`
	AnnualSales        string `json:"annual_sales"`
	PhoneNo            string `json:"phone_no"`
	Fax                string `json:"fax"`
	Email              string `json:"email"`
	Website            string `json:"website"`
	StreetAddress      string `json:"street_address"`
	ProvinceId         string `json:"province_id"`
	CityId             string `json:"city_id"`
	DistrictId         string `json:"district_id"`
	SubdistrictId      string `json:"subdistrict_id"`
	Postal             string `json:"postal"`
	TaxMandatory       string `json:"tax_mandatory"`
	TaxStatus          string `json:"tax_status"`
	NpwpNo             string `json:"npwp_no"`
	NpwpName           string `json:"npwp_name"`
	NpwpAddress        string `json:"npwp_address"`
	VatImg             string `json:"vat_img"`
	CreateBy           string `json:"create_by"`
	CreateByIp         string `json:"create_by_ip"`
	BillingCode        string `json:"billing_code"`
	BillingName        string `json:"billing_name"`
	BillingMethod      string `json:"billing_method"`
	FactureSchedule    string `json:"facture_schedule"`
	BillingDoc         string `json:"billing_doc"`
	ReturAvailable     string `json:"retur_available"`
	BillingAddress     string `json:"billing_address"`
	BillingProvince    string `json:"billing_province"`
	BillingCity        string `json:"billing_city"`
	BillingDistrict    string `json:"billing_district"`
	BillingSubdistrict string `json:"billing_subdistrict"`
	BillingPostal      string `json:"billing_postal"`
	PaymentTerm        string `json:"payment_term"`
}

type InfoNewCustomerRequestDelete struct {
	CustomerRequestNo string `json:"customer_request_no"`
}

type InfoNewCustomerRequestUpdate struct {
	CustomerRequestNo  string `json:"customer_request_no"`
	CustomerName       string `json:"customer_name"`
	CustomerType       string `json:"customer_type"`
	YearEstablished    string `json:"year_established"`
	EmployeeAmount     string `json:"employee_amount"`
	AnnualSales        string `json:"annual_sales"`
	PhoneNo            string `json:"phone_no"`
	Fax                string `json:"fax"`
	Email              string `json:"email"`
	Website            string `json:"website"`
	StreetAddress      string `json:"street_address"`
	ProvinceId         string `json:"province_id"`
	CityId             string `json:"city_id"`
	DistrictId         string `json:"district_id"`
	SubdistrictId      string `json:"subdistrict_id"`
	Postal             string `json:"postal"`
	TaxMandatory       string `json:"tax_mandatory"`
	TaxStatus          string `json:"tax_status"`
	NpwpNo             string `json:"npwp_no"`
	NpwpName           string `json:"npwp_name"`
	NpwpAddress        string `json:"npwp_address"`
	VatImg             string `json:"vat_img"`
	CreateBy           string `json:"create_by"`
	CreateByIp         string `json:"create_by_ip"`
	BillingCode        string `json:"billing_code"`
	BillingName        string `json:"billing_name"`
	BillingMethod      string `json:"billing_method"`
	FactureSchedule    string `json:"facture_schedule"`
	BillingDoc         string `json:"billing_doc"`
	ReturAvailable     string `json:"retur_available"`
	BillingAddress     string `json:"billing_address"`
	BillingProvince    string `json:"billing_province"`
	BillingCity        string `json:"billing_city"`
	BillingDistrict    string `json:"billing_district"`
	BillingSubdistrict string `json:"billing_subdistrict"`
	BillingPostal      string `json:"billing_postal"`
	PaymentTerm        string `json:"payment_term"`
	Status             string `json:"status"`
}

type DataNeedApproveCustomerRequest struct {
	CustomerRequestNo string `json:"customer_request_no"`
	Status            string `json:"status"`
	CustomerName      string `json:"customer_name"`
	CustomerType      string `json:"customer_type"`
	VatRegNo          string `json:"vat_reg_no"`
	StreetAddress     string `json:"street_address"`
	CityId            string `json:"city_id"`
	DistrictId        string `json:"district_id"`
	OutletCount       string `json:"outlet_count"`
	PicCount          string `json:"pic_count"`
}

type InfoNewCustomerRequestApproval struct {
	CustomerRequestNo string `json:"customer_request_no"`
}
