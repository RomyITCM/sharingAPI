package entities

type DataNewCustomerDocument struct {
	DocId             string `json:"doc_id"`
	DocumentType      string `json:"document_type"`
	DocumentNo        string `json:"document_no"`
	DocumentName      string `json:"document_name"`
	DocumentAddress   string `json:"document_address"`
	DocImg            string `json:"doc_img"`
	CustomerRequestNo string `json:"customer_request_no"`
	CustomerNo        string `json:"customer_no"`
	BankCode          string `json:"bank_code"`
	BankName          string `json:"bank_name"`
}

type InfoNewCustomerDocument struct {
	DocumentType      string `json:"document_type"`
	DocumentNo        string `json:"document_no"`
	DocumentName      string `json:"document_name"`
	DocumentAddress   string `json:"document_address"`
	DocImg            string `json:"doc_img"`
	CustomerRequestNo string `json:"customer_request_no"`
	CustomerNo        string `json:"customer_no"`
	BankCode          string `json:"bank_code"`
	CreatedBy         string `json:"created_by"`
	CreatedByIp       string `json:"created_by_ip"`
}

type DocId struct {
	DocId string `json:"doc_id"`
}

type InfoNewCustomerDocumentUpdate struct {
	CustomerRequestNo string `json:"customer_request_no"`
	DocumentType      string `json:"document_type"`
	DocumentNo        string `json:"document_no"`
	DocumentName      string `json:"document_name"`
	DocumentAddress   string `json:"document_address"`
	DocImg            string `json:"doc_img"`
	BankCode          string `json:"bank_code"`
	CreatedBy         string `json:"created_by"`
	CreatedByIp       string `json:"created_by_ip"`
}

type DataCustomerBanks struct {
	Code     string `json:"code"`
	BankName string `json:"bank_name"`
}

type InfoCustomerDocument struct {
	DocumentNo string `json:"document_no"`
	CustomerNo string `json:"customer_no"`
	DocImg     string `json:"doc_img"`
	// DocPath    string `json:"doc_path"`
	CreatedBy   string `json:"created_by"`
	CreatedByIp string `json:"created_by_ip"`
}
