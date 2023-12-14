package entities

type DataSMDList struct {
	ShipToName    string `json:"ship_to_name"`
	ShipToNo      string `json:"ship_to_no"`
	CustomerName  string `json:"customer_name"`
	CustomerNo    string `json:"customer_no"`
	BillToName    string `json:"bill_to_name"`
	BillToNo      string `json:"bill_to_no"`
	ShipToAddress string `json:"ship_to_address"`
	LastReport    string `json:"last_report"`
}

type DataSMDFreezerList struct {
	Freezer          string `json:"freezer"`
	Location         string `json:"location"`
	Power            string `json:"power"`
	Condition        string `json:"condition"`
	OutsideCondition string `json:"outside_condition"`
	TransNo          string `json:"trans_no"`
	SerialNo         string `json:"serial_no"`
}

type DataSMDFreezer struct {
	FreezerAvailable     int    `json:"freezer_available"`
	SerialNo             string `json:"serial_no"`
	Merk                 int    `json:"merk"`
	NoteMerk             string `json:"note_merk"`
	Capacity             int    `json:"capacity"`
	Location             int    `json:"location"`
	Access               int    `json:"access"`
	NoteAccess           string `json:"note_access"`
	Power                int    `json:"power"`
	NotePower            string `json:"note_power"`
	Condition            int    `json:"condition"`
	NoteCondition        string `json:"note_condition"`
	OutsideCondition     int    `json:"outside_condition"`
	NoteOutsideCondition string `json:"note_outside_condition"`
	Suhu                 int    `json:"suhu"`
}

type DataSMDFreezerInsert struct {
	Header        Header        `json:"header"`
	DetailFreezer DetailFreezer `json:"detail_freezer"`
}

type DataSMDFreezerUpdate struct {
	TransNo       TransNo       `json:"trans_no"`
	DetailFreezer DetailFreezer `json:"detail_freezer"`
}

type DataSMDSKU struct {
	Header    Header      `json:"header"`
	DetailSKU interface{} `json:"detail_sku"`
}

type TransNo struct {
	TransNo     string `json:"trans_no"`
	MsgType     string `json:"msg_type"`
	MsgError    string `json:"msg_error"`
	CreatedBy   string `json:"created_by"`
	CreatedByIP string `json:"created_by_ip"`
}

type Header struct {
	CustomerNo      string `json:"customer_no"`
	BillTo          string `json:"bill_to"`
	ShipTo          string `json:"ship_to"`
	FrezerAvailable int    `json:"frezer_available"`
	CreatedBy       string `json:"created_by"`
	CreatedByIP     string `json:"created_by_ip"`
}

type DetailFreezer struct {
	SerialNo             string `json:"serial_no"`
	Merk                 int    `json:"merk"`
	NoteMerk             string `json:"note_merk"`
	Capacity             int    `json:"capacity"`
	Location             int    `json:"location"`
	Access               int    `json:"access"`
	NoteAccess           string `json:"note_access"`
	Power                int    `json:"power"`
	NotePower            string `json:"note_power"`
	Condition            int    `json:"condition"`
	NoteCondition        string `json:"note_condition"`
	OutsideCondition     int    `json:"outside_condition"`
	NoteOutsideCondition string `json:"note_outside_condition"`
	Suhu                 int    `json:"suhu"`
}
