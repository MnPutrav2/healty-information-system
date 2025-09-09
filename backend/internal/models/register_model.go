package models

type ResponseRegisterPatient struct {
	CareNumber      string `json:"care_number"`
	RegisterNumber  string `json:"register_number"`
	RegisterDate    string `json:"register_date"`
	MedicalRecord   string `json:"medical_record"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	PaymentMethod   string `json:"payment_method"`
	Policlinic_id   string `json:"policlinic_id"`
	Policlinic_name string `json:"policlinic_name"`
	Doctor_id       string `json:"doctor_id"`
	Doctor_name     string `json:"doctor_name"`
	Status          string `json:"status"`
}

type RequestRegisterPatient struct {
	CareNumber     string `json:"care_number"`
	RegisterNumber string `json:"register_number"`
	RegisterDate   string `json:"register_date"`
	MedicalRecord  string `json:"medical_record"`
	PaymentMethod  string `json:"payment_method"`
	Policlinic     string `json:"policlinic"`
	Doctor         string `json:"doctor"`
}

type RegisterStatus struct {
	Status string `json:"status"`
}

type Policlinics struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	PoliType string `json:"poli_type"`
}
