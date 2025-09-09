package models

type Examination struct {
	ID              string `json:"id"`
	ExaminationName string `json:"examination_name"`
	PaymentMethod   string `json:"payment_method"`
	DoctorCost      int    `json:"doctor_cost"`
	NurseCost       int    `json:"nurse_cost"`
	ManagementCost  int    `json:"management_cost"`
}

type ExaminationRequest struct {
	CareNumber  string `json:"care_number"`
	Examination string `json:"examination"`
	DoctorID    string `json:"doctor_id"`
	NurseID     string `json:"nurse_id"`
	ServiceType string `json:"service_type"`
	Date        string `json:"date"`
}

type ExaminationRequestDetail struct {
	ID            string `json:"id"`
	CareNumber    string `json:"care_number"`
	ExaminationID string `json:"examination_id"`
	Examination   string `json:"examination"`
	DoctorID      string `json:"doctor_id"`
	DoctorName    string `json:"doctor_name"`
	NurseID       string `json:"nurse_id"`
	NurseName     string `json:"nurse_name"`
	ServiceType   string `json:"service_type"`
	Date          string `json:"date"`
	Total         int    `json:"total"`
}
