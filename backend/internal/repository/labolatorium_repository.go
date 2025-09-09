package repository

import (
	"database/sql"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type LabolatoriumRepository interface {
	CreateLabData(lab models.LabRequestData) (string, error)
	GetLabolatoriumData(limit int, search string) ([]models.LabResponseData, error)
	DeleteLabolatoriumData(id string) error
	UpdateLabolatoriumData(data models.LabUpdate) error
	CreateLabTemplate(lab models.LabTemplateRequestData) error
	GetLabolatoriumTemplate(id string) ([]models.LabolatoriumTemplate, error)
	CreateLabolatoriumRequest(data models.LabolatoriumRequest) (string, error)
	DeleteLabolatoriumTemplate(id string, unit string) error
	GetLabolatoriumRequest(date1 string, date2 string) ([]models.LabolatoriumRequestData, error)
	UpdateLabSample(data models.LabSampleRequest, id string) error
	GetLabolatoriumDataDetail(id string) ([]models.LabDetailData, error)
	UpdateLabolatoriumValue(id string, val models.LabValueRequest) error
}

type labolatoriumRepository struct {
	sql *sql.DB
}

func NewLabolatoriumRepository(sql *sql.DB) LabolatoriumRepository {
	return &labolatoriumRepository{sql}
}

func (q *labolatoriumRepository) GetLabolatoriumData(limit int, search string) ([]models.LabResponseData, error) {
	result, err := q.sql.Query("SELECT id, name, referral_fee, officer_fee, management, total FROM labolatorium_datas WHERE name LIKE $1", search)
	if err != nil {
		return nil, err
	}

	var lab []models.LabResponseData
	for result.Next() {
		var l models.LabResponseData

		err := result.Scan(&l.ID, &l.Name, &l.ReferralFee, &l.OfficerFee, &l.Management, &l.Total)
		if err != nil {
			return nil, err
		}

		lab = append(lab, l)
	}

	return lab, nil
}

func (q *labolatoriumRepository) CreateLabData(lab models.LabRequestData) (string, error) {
	var check bool
	err := q.sql.QueryRow("SELECT EXISTS (SELECT 1 FROM labolatorium_datas WHERE id = $1)", lab.ID).Scan(&check)
	if err != nil {
		return "", err
	}

	if check {
		return "duplicate", nil
	}

	total := lab.OfficerFee + lab.ReferralFee + lab.Management
	_, err = q.sql.Exec("INSERT INTO labolatorium_datas(id, name, referral_fee, officer_fee, management, total) VALUES($1, $2, $3, $4, $5, $6)", lab.ID, lab.Name, lab.ReferralFee, lab.OfficerFee, lab.Management, total)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (q *labolatoriumRepository) DeleteLabolatoriumData(id string) error {
	_, err := q.sql.Exec("DELETE FROM labolatorium_datas WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (q *labolatoriumRepository) UpdateLabolatoriumData(data models.LabUpdate) error {
	total := data.Data.ReferralFee + data.Data.OfficerFee + data.Data.Management
	_, err := q.sql.Exec("UPDATE labolatorium_datas SET id = $1, name = $2, referral_fee = $3, officer_fee = $4, management = $5, total = $6 WHERE id = $7", data.Data.ID, data.Data.Name, data.Data.ReferralFee, data.Data.OfficerFee, data.Data.Management, total, data.ID)
	if err != nil {
		return err
	}

	return nil
}

func (q *labolatoriumRepository) CreateLabTemplate(lab models.LabTemplateRequestData) error {
	_, err := q.sql.Exec("INSERT INTO labolatorium_template(id, name, unit, normal_value) VALUES($1, $2, $3, $4)", lab.ID, lab.Value.Name, lab.Value.Unit, lab.Value.NormalValue)
	if err != nil {
		return err
	}

	return nil
}

func (q *labolatoriumRepository) GetLabolatoriumTemplate(id string) ([]models.LabolatoriumTemplate, error) {
	res, err := q.sql.Query("SELECT id, name, unit, normal_value FROM labolatorium_template WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	var t []models.LabolatoriumTemplate
	for res.Next() {
		var m models.LabolatoriumTemplate

		err := res.Scan(&m.ID, &m.Name, &m.Unit, &m.NormalValue)
		if err != nil {
			return nil, err
		}

		t = append(t, m)
	}

	return t, nil
}

func (q *labolatoriumRepository) CreateLabolatoriumRequest(data models.LabolatoriumRequest) (string, error) {
	var check bool
	err := q.sql.QueryRow("SELECT EXISTS (SELECT 1 FROM labolatorium_request WHERE lab_id = $1)", data.LabID).Scan(&check)
	if err != nil {
		return "", err
	}

	if check {
		return "duplicate", nil
	}

	_, err = q.sql.Exec("INSERT INTO labolatorium_request(care_number, lab_id, referral_id, officer_id, request_date, sample, validate, validate_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", data.CareNumber, data.LabID, data.ReferralID, data.OfficerID, data.Date, nil, false, nil)
	if err != nil {
		return "", err
	}

	for _, d := range data.Data {
		_, err = q.sql.Exec("INSERT INTO labolatorium_detail (lab_request, lab_id, referral_fee, officer_fee, management, total) VALUES ($1, $2, $3, $4, $5, $6)", data.LabID, d.Id, d.ReferralFee, d.OfficerFee, d.Management, d.Total)
		if err != nil {
			return "", err
		}
	}

	return "success", nil
}

func (q *labolatoriumRepository) DeleteLabolatoriumTemplate(id string, unit string) error {
	_, err := q.sql.Exec("DELETE FROM labolatorium_template WHERE id = $1 AND unit = $2", id, unit)
	if err != nil {
		return err
	}

	return nil
}

func (q *labolatoriumRepository) GetLabolatoriumRequest(date1 string, date2 string) ([]models.LabolatoriumRequestData, error) {
	query := `SELECT labolatorium_request.lab_id, labolatorium_request.care_number, patients.name, employees.name AS referral_name, labolatorium_request.request_date, labolatorium_request.sample, labolatorium_request.validate, labolatorium_request.validate_date 
				FROM labolatorium_request 
				INNER JOIN doctors ON labolatorium_request.referral_id = doctors.id 
				INNER JOIN employees ON doctors.employee_id = employees.id 
				INNER JOIN registration ON labolatorium_request.care_number = registration.care_number 
				INNER JOIN patients ON registration.medical_record = patients.medical_record
				WHERE labolatorium_request.request_date BETWEEN $1 AND $2;`

	result, err := q.sql.Query(query, date1, date2)
	if err != nil {
		return nil, err
	}

	var datas []models.LabolatoriumRequestData
	for result.Next() {
		var dt models.LabolatoriumRequestData

		err := result.Scan(&dt.LabID, &dt.CareNumber, &dt.PatientName, &dt.ReferralName, &dt.RequestDate, &dt.Sample, &dt.Validate, &dt.ValidateDate)
		if err != nil {
			return nil, err
		}

		datas = append(datas, dt)
	}

	return datas, nil
}

func (q *labolatoriumRepository) UpdateLabSample(data models.LabSampleRequest, id string) error {
	if data.Status {
		_, err := q.sql.Exec("UPDATE labolatorium_request SET sample = $1 WHERE lab_id = $2", data.Value, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *labolatoriumRepository) GetLabolatoriumDataDetail(id string) ([]models.LabDetailData, error) {
	result, err := q.sql.Query("SELECT labolatorium_detail.lab_id, labolatorium_datas.name FROM labolatorium_detail INNER JOIN labolatorium_datas ON labolatorium_detail.lab_id = labolatorium_datas.id WHERE labolatorium_detail.lab_request = $1", id)
	if err != nil {
		return nil, err
	}

	var datas []models.LabDetailData
	for result.Next() {
		var fin models.LabDetailData

		err = result.Scan(&fin.ID, &fin.Name)
		if err != nil {
			return nil, err
		}

		res, err := q.sql.Query("SELECT template_id, id, name, unit, normal_value FROM labolatorium_template WHERE id = $1", fin.ID)
		if err != nil {
			return nil, err
		}

		var tem []models.LabolatoriumTemplate
		for res.Next() {
			var temp models.LabolatoriumTemplate

			err = res.Scan(&temp.TempID, &temp.ID, &temp.Name, &temp.Unit, &temp.NormalValue)
			if err != nil {
				return nil, err
			}

			tem = append(tem, temp)
		}

		fin.Template = append(fin.Template, tem...)
		datas = append(datas, fin)
	}

	return datas, nil
}

func (q *labolatoriumRepository) UpdateLabolatoriumValue(id string, val models.LabValueRequest) error {
	_, err := q.sql.Exec("UPDATE labolatorium_request SET validate = $1, validate_date = $2 WHERE lab_id = $3", true, val.Date, id)
	if err != nil {
		return err
	}

	for _, d := range val.Value {
		_, err := q.sql.Exec("INSERT INTO labolatorium_detail_template(lab_request, lab_id, template_id, value) VALUES($1, $2, $3, $4)", id, d.ID, d.TempID, d.Value)
		if err != nil {
			return err
		}
	}

	return nil
}
