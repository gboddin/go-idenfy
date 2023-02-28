package idenfy

type AML struct {
	Status           ServiceStatus            `json:"status"`
	Data             []map[string]interface{} `json:"data"`
	ServiceName      string                   `json:"serviceName"`
	ServiceGroupType string                   `json:"serviceGroupType"`
	Uid              string                   `json:"uid"`
	ErrorMessage     string                   `json:"errorMessage"`
}

type ServiceStatus struct {
	ServiceSuspected bool   `json:"serviceSuspected"`
	CheckSuccessful  bool   `json:"checkSuccessful"`
	ServiceFound     bool   `json:"serviceFound"`
	ServiceUsed      bool   `json:"serviceUsed"`
	OverallStatus    string `json:"overallStatus"`
}
