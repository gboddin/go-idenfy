package idenfy

type LID struct {
	Status           ServiceStatus `json:"status"`
	Data             []LIDData     `json:"data"`
	ServiceName      string        `json:"serviceName"`
	ServiceGroupType string        `json:"serviceGroupType"`
	Uid              string        `json:"uid"`
	ErrorMessage     string        `json:"errorMessage"`
}

type LIDData struct {
	DocumentNumber string   `json:"documentNumber"`
	DocumentType   Document `json:"documentType"`
	Valid          bool     `json:"valid"`
	ExpiryDate     string   `json:"expiryDate"`
	CheckDate      string   `json:"checkDate"`
}
