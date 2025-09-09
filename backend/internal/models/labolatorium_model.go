package models

type LabRequestData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ReferralFee int    `json:"referral_fee"`
	OfficerFee  int    `json:"officer_fee"`
	Management  int    `json:"management"`
	Total       int    `json:"total"`
}

type LabUpdate struct {
	ID   string         `json:"id"`
	Data LabRequestData `json:"data"`
}

type LabResponseData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ReferralFee int    `json:"referral_fee"`
	OfficerFee  int    `json:"officer_fee"`
	Management  int    `json:"management"`
	Total       int    `json:"total"`
}

// type LabResponseData struct {
// 	ID          string                 `json:"id"`
// 	Name        string                 `json:"name"`
// 	RefferalFee int                    `json:"refferal_fee"`
// 	OfficerFee  int                    `json:"officer_fee"`
// 	Management  int                    `json:"management"`
// 	Total       string                 `json:"total"`
// 	Template    []LabolatoriumTemplate `json:"template"`
// }

type LabolatoriumTemplate struct {
	TempID      string `json:"template_id"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	NormalValue string `json:"normal_value"`
}

type LabolatoriumRequest struct {
	CareNumber string           `json:"care_number"`
	LabID      string           `json:"lab_id"`
	OfficerID  string           `json:"officer_id"`
	ReferralID string           `json:"referral_id"`
	Date       string           `json:"date"`
	Data       []DataLabRequest `json:"data"`
}

type DataLabRequest struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	ReferralFee int                    `json:"referral_fee"`
	OfficerFee  int                    `json:"officer_fee"`
	Management  int                    `json:"management"`
	Total       int                    `json:"total"`
	Template    []LabolatoriumTemplate `json:"template"`
}

type LabTemplateRequestData struct {
	TempID string               `json:"template_id"`
	ID     string               `json:"id"`
	Value  LabolatoriumTemplate `json:"value"`
}

type LabSampleRequest struct {
	Status bool   `json:"status"`
	Value  string `json:"value"`
}

type LabolatoriumRequestData struct {
	LabID        string  `json:"lab_id"`
	CareNumber   string  `json:"care_number"`
	PatientName  string  `json:"patient_name"`
	ReferralName string  `json:"referral_name"`
	RequestDate  string  `json:"request_date"`
	Sample       *string `json:"sample"`
	Validate     bool    `json:"validate"`
	ValidateDate *string `json:"validate_date"`
}

type LabDetailData struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Template []LabolatoriumTemplate `json:"template"`
}

type LabValueRequest struct {
	Date  string `json:"date"`
	Value []Labc `json:"value"`
}

type Labc struct {
	TempID string `json:"template_id"`
	ID     string `json:"id"`
	Value  string `json:"value"`
}
