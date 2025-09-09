package models

type ExaminationCostRequest struct {
	ExaminationName string `json:"examination_name"`
	PaymentMethod   string `json:"payment_method"`
	DoctorCost      int    `json:"doctor_cost"`
	NurseCost       int    `json:"nurse_cost"`
	ManagementCost  int    `json:"management_cost"`
}

type ExaminationCost struct {
	ID              string `json:"id"`
	ExaminationName string `json:"examination_name"`
	PaymentMethod   string `json:"payment_method"`
	DoctorCost      int    `json:"doctor_cost"`
	NurseCost       int    `json:"nurse_cost"`
	ManagementCost  int    `json:"management_cost"`
}
