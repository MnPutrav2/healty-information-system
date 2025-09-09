package models

type DoctorSchedule struct {
	Day       string `json:"day"`
	FirstTime string `json:"first_time"`
	LastTime  string `json:"last_time"`
}

type DoctorData struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Specialis     bool   `json:"specialis"`
	TypeSpecialis string `json:"type_specialis"`
}
