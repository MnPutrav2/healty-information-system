package models

type PatientData struct {
	MedicalRecord string `json:"medical_record"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Wedding       string `json:"wedding"`
	Religion      string `json:"religion"`
	Education     string `json:"education"`
	BirthPlace    string `json:"birth_place"`
	BirthDate     string `json:"birth_date"`
	Work          string `json:"work"`
	Address       string `json:"address"`
	Village       int    `json:"village" validate:"max=4"`
	District      int    `json:"district" validate:"max=2"`
	Regencie      int    `json:"regencie" validate:"max=2"`
	Province      int    `json:"province" validate:"max=2"`
	NIK           string `json:"nik"`
	BPJS          string `json:"bpjs"`
	PhoneNumber   string `json:"phone_number"`
	ParentName    string `json:"parent_name"`
	Relationship  string `json:"relationship"`
	ParentGender  string `json:"parent_gender"`
}

type PatientDataRes struct {
	MedicalRecord string     `json:"medical_record"`
	Name          string     `json:"name"`
	Gender        string     `json:"gender"`
	Wedding       string     `json:"wedding"`
	Religion      string     `json:"religion"`
	Education     string     `json:"education"`
	BirthPlace    string     `json:"birth_place"`
	BirthDate     string     `json:"birth_date"`
	Work          string     `json:"work"`
	Address       string     `json:"address"`
	Village       int        `json:"village" validate:"max=4"`
	District      int        `json:"district" validate:"max=2"`
	Regencie      int        `json:"regencie" validate:"max=2"`
	Province      int        `json:"province" validate:"max=2"`
	NIK           string     `json:"nik"`
	BPJS          string     `json:"bpjs"`
	PhoneNumber   string     `json:"phone_number"`
	ParentName    string     `json:"parent_name"`
	Relationship  string     `json:"relationship"`
	ParentGender  string     `json:"parent_gender"`
	Schedule      []Schedule `json:"schedule"`
}

type Schedule struct {
	DoctorID   string `json:"doctor_id"`
	DoctorName string `json:"doctor_name"`
	Policlinic string `json:"policlinic"`
	DateTime   string `json:"date"`
}

type PatientDataUpdate struct {
	MedicalRecordKey string      `json:"medical_record"`
	Update           PatientData `json:"update"`
}

type PatientBundle struct {
	ResourceType string  `json:"resourceType"`
	Type         string  `json:"type"`
	Total        int     `json:"total"`
	Link         []Link  `json:"link"`
	Entry        []Entry `json:"entry"`
}

type Link struct {
	Relation string `json:"relation"`
	URL      string `json:"url"`
}

type Entry struct {
	FullURL  string   `json:"fullUrl"`
	Resource *Patient `json:"resource"`
}

type Patient struct {
	ResourceType string        `json:"resourceType"`
	ID           string        `json:"id"`
	Active       bool          `json:"active"`
	Identifier   []Identifier  `json:"identifier"`
	Name         []HumanName   `json:"name"`
	Link         []PatientLink `json:"link"`
	Meta         Meta          `json:"meta"`
}

type Identifier struct {
	System string `json:"system"`
	Use    string `json:"use"`
	Value  string `json:"value"`
}

type HumanName struct {
	Use  string `json:"use"`
	Text string `json:"text"`
}

type PatientLink struct {
	Other Reference `json:"other"`
	Type  string    `json:"type"`
}

type Reference struct {
	Reference string `json:"reference"`
}

type Meta struct {
	LastUpdated string   `json:"lastUpdated"`
	Profile     []string `json:"profile"`
	VersionID   string   `json:"versionId"`
}
