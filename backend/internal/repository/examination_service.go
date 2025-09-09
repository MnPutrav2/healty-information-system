package repository

import (
	"database/sql"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type ExaminationRepository interface {
	AddExamination(data models.ExaminationRequest) error
	GetExamination(id string) ([]models.ExaminationRequestDetail, error)
	DeleteExamination(id string) error
}

type examinationRepository struct {
	sql *sql.DB
}

func NewExaminationRepository(sql *sql.DB) ExaminationRepository {
	return &examinationRepository{sql}
}

func (q *examinationRepository) AddExamination(data models.ExaminationRequest) error {

	var did *string
	var dido int
	if data.DoctorID != "" {
		var dc int

		_ = q.sql.QueryRow("SELECT doctor_cost FROM examination_cost WHERE id = $1", data.Examination).Scan(&dc)

		did = &data.DoctorID
		dido = dc
	} else {
		did = nil
		dido = 0
	}

	var nid *string
	var nido int
	if data.NurseID != "" {
		var nc int

		_ = q.sql.QueryRow("SELECT nurse_cost FROM examination_cost WHERE id = $1", data.Examination).Scan(&nc)

		nid = &data.NurseID
		nido = nc
	} else {
		nid = nil
		nido = 0
	}

	var mnj int
	_ = q.sql.QueryRow("SELECT management_cost FROM examination_cost WHERE id = $1", data.Examination).Scan(&mnj)

	_, err := q.sql.Exec("INSERT INTO examination(care_number, examination, doctor_id, nurse_id, service_type, date, total) VALUES($1, $2, $3, $4, $5, $6, $7)", data.CareNumber, data.Examination, did, nid, data.ServiceType, data.Date, dido+nido+mnj)
	if err != nil {
		return err
	}

	return nil
}

func (q *examinationRepository) GetExamination(id string) ([]models.ExaminationRequestDetail, error) {
	res, err := q.sql.Query("SELECT examination.id, examination.care_number, examination.examination, examination_cost.examination_name, COALESCE(examination.doctor_id, '') AS doctor_id, COALESCE(doctors.name, '') AS name_dr, COALESCE(examination.nurse_id, '') AS nurse_id, COALESCE(employees.name, '') AS name_pr, examination.service_type, examination.date, examination.total FROM examination INNER JOIN examination_cost ON examination.examination = examination_cost.id LEFT JOIN doctors ON examination.doctor_id = doctors.id LEFT JOIN employees ON examination.nurse_id = employees.id WHERE examination.care_number = $1", id)
	if err != nil {
		return nil, err
	}

	var datas []models.ExaminationRequestDetail

	for res.Next() {
		var data models.ExaminationRequestDetail

		err := res.Scan(&data.ID, &data.CareNumber, &data.ExaminationID, &data.Examination, &data.DoctorID, &data.DoctorName, &data.NurseID, &data.NurseName, &data.ServiceType, &data.Date, &data.Total)
		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (q *examinationRepository) DeleteExamination(id string) error {
	_, err := q.sql.Exec("DELETE FROM examination WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

//  SELECT SUM(CASE WHEN examination.nurse_id IS NULL THEN 0 ELSE examination_cost.nurse_cost END AS nurse_cost, CASE WHEN examination.doctor_id IS NULL THEN 0 ELSE examination_cost.doctor_cost END AS doctor_cost, SUM(nure_cost + doctor_cost), examination.id, examination.care_number, examination.examination, examination_cost.examination_name, COALESCE(examination.doctor_id, '') AS doctor_id, doctors.name, COALESCE(examination.nurse_id, '') AS nurse_id, employees.name, examination.service_type, examination.date FROM examination INNER JOIN examination_cost ON examination.examination = examination_cost.id LEFT JOIN doctors ON examination.doctor_id = doctors.id LEFT JOIN employees ON examination.nurse_id = employees.id WHERE examination.care_number = '20250907000001';
