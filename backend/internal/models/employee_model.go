package models

type EmployeeData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Wedding     string `json:"wedding"`
	Address     string `json:"address"`
	NIK         string `json:"nik"`
	BPJS        string `json:"bpjs"`
	NPWP        string `json:"npwp"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Status      bool   `json:"status"`
}

type UserReq struct {
	ID       string `json:"id"`
	EmpID    string `json:"employee_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
