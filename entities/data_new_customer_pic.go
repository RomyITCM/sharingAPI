package entities

type DataNewCustomerPic struct {
	PicId               string `json:"pic_id"`
	Honorific           string `json:"honorific"`
	PicName             string `json:"pic_name"`
	Position            string `json:"position"`
	PhoneNo             string `json:"phone_no"`
	Email               string `json:"email"`
	OutletId            string `json:"outlet_id"`
	OutletName          string `json:"outlet_name"`
	CustomerAddressId   string `json:"customer_address_id"`
	CustomerAddressName string `json:"customer_address_name"`
}

type DataNewCustomerPicDetail struct {
	PicId             string `json:"pic_id"`
	Honorific         string `json:"honorific"`
	PicName           string `json:"pic_name"`
	Position          string `json:"position"`
	PhoneNo           string `json:"phone_no"`
	Email             string `json:"email"`
	Status            string `json:"status"`
	OutletId          string `json:"outlet_id"`
	CustomerAddressId string `json:"customer_address_id"`
}

type PicId struct {
	PicId string `json:"pic_id"`
}

type InfoNewCustomerPic struct {
	CustomerId        string `json:"customer_id"`
	Honorific         string `json:"honorific"`
	PicName           string `json:"pic_name"`
	Position          string `json:"position"`
	PhoneNo           string `json:"phone_no"`
	Email             string `json:"email"`
	OutletId          string `json:"outlet_id"`
	CustomerAddressId string `json:"customer_address_id"`
	CreatedBy         string `json:"created_by"`
	CreatedByIp       string `json:"created_by_ip"`
}

type InfoNewCustomerPicUpdate struct {
	PicId             string `json:"pic_id"`
	Honorific         string `json:"honorific"`
	PicName           string `json:"pic_name"`
	Position          string `json:"position"`
	PhoneNo           string `json:"phone_no"`
	Email             string `json:"email"`
	OutletId          string `json:"outlet_id"`
	CustomerAddressId string `json:"customer_address_id"`
	CreatedBy         string `json:"created_by"`
	CreatedByIp       string `json:"created_by_ip"`
}

type DataPicResponsibility struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
