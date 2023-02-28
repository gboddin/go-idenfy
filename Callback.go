package idenfy

type IdentityCallbackResp struct {
	Final           *bool             `json:"final"`
	Platform        string            `json:"platform"`
	Status          Status            `json:"status"`
	Data            Data              `json:"data"`
	FileUrls        map[string]string `json:"fileUrls"`
	AML             []AML             `json:"AML"`
	LID             []LID             `json:"LID"`
	ScanRef         string            `json:"scanRef"`
	ExternalRef     string            `json:"externalRef"`
	ClientId        string            `json:"clientId"`
	StartTime       int64             `json:"startTime"`
	FinishTime      int64             `json:"finishTime"`
	ClientIp        string            `json:"clientIp"`
	ClientIpCountry string            `json:"clientIpCountry"`
	ClientLocation  string            `json:"clientLocation"`
	GdcMatch        bool              `json:"gdcMatch"`
	ManualAddress   string            `json:"manualAddress"`
}

type Status struct {
	Overall         string   `json:"overall"`
	FraudTags       []string `json:"fraudTags"`
	MismatchTags    []string `json:"mismatchTags"`
	AutoFace        string   `json:"autoFace"`
	ManualFace      string   `json:"manualFace"`
	AutoDocument    string   `json:"autoDocument"`
	ManualDocument  string   `json:"manualDocument"`
	AdditionalSteps string   `json:"additionalSteps"`
}

type Data struct {
	DocFirstname           string   `json:"docFirstName"`
	DocLastname            string   `json:"docLastName"`
	DocNumber              string   `json:"docNumber"`
	DocPersonalCode        string   `json:"docPersonalCode"`
	DocExpiry              string   `json:"docExpiry"`
	DocDob                 string   `json:"docDob"`
	DocDateOfIssue         string   `json:"docDateOfIssue"`
	DocType                Document `json:"docType"`
	DocSex                 string   `json:"docSex"`
	DocNationality         string   `json:"docNationality"`
	DocIssuingCountry      string   `json:"docIssuingCountry"`
	DocTemporaryAddress    string   `json:"docTemporaryAddress"`
	DocBirthName           string   `json:"docBirthName"`
	BirthPlace             string   `json:"birthPlace"`
	Authority              string   `json:"authority"`
	Address                string   `json:"address"`
	MothersMaidenName      string   `json:"mothersMaidenName"`
	DriverLicenseCategory  string   `json:"driverLicenseCategory"`
	ManuallyDataChanged    bool     `json:"manuallyDataChanged"`
	FullName               string   `json:"fullName"`
	SelectedCountry        string   `json:"selectedCountry"`
	OrgFirstName           string   `json:"orgFirstName"`
	OrgLastName            string   `json:"orgLastName"`
	OrgNationality         string   `json:"orgNationality"`
	OrgBirthPlace          string   `json:"orgBirthPlace"`
	OrgAuthority           string   `json:"orgAuthority"`
	OrgAddress             string   `json:"orgAddress"`
	OrgTemporaryAddress    string   `json:"OrgTemporaryAddress"`
	OrgMothersMaidenName   string   `json:"orgMothersMaidenName"`
	OrgBirthName           string   `json:"orgBirthName"`
	AgeEstimate            string   `json:"ageEstimate"`
	ClientIpProxyRiskLevel string   `json:"clientIpProxyRiskLevel"`
	DuplicateFaces         []string `json:"duplicateFaces"`
	DuplicateDocFaces      []string `json:"duplicateDocFaces"`
}
