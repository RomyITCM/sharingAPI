package entities

type DataFindingInsert struct {
	EmailTo     string `json:"email_to"`
	Subject     string `json:"subject"`
	Remark      string `json:"remark"`
	CreatedBy   string `json:"created_by"`
	CreatedByIP string `json:"created_by_ip"`
}
