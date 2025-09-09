package repository

import (
	"database/sql"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type DoctorRepository interface {
	GetDoctors(id string) ([]models.DoctorData, error)
	GetDoctorsAll() ([]models.DoctorData, error)
	GetDoctorSchedule(id string) ([]models.DoctorSchedule, error)
}

type doctorRepository struct {
	sql *sql.DB
}

func NewDoctorRepository(sql *sql.DB) DoctorRepository {
	return &doctorRepository{sql}
}

func (q *doctorRepository) GetDoctors(id string) ([]models.DoctorData, error) {
	res, err := q.sql.Query("SELECT id, name, specialis, type_specialis FROM doctors WHERE type_specialis = $1", id)
	if err != nil {
		return nil, err
	}

	var d []models.DoctorData
	for res.Next() {
		var x models.DoctorData

		err := res.Scan(&x.ID, &x.Name, &x.Specialis, &x.TypeSpecialis)
		if err != nil {
			return nil, err
		}

		d = append(d, x)
	}

	return d, nil
}

func (q *doctorRepository) GetDoctorsAll() ([]models.DoctorData, error) {
	res, err := q.sql.Query("SELECT id, name, specialis, type_specialis FROM doctors")
	if err != nil {
		return nil, err
	}

	var d []models.DoctorData
	for res.Next() {
		var x models.DoctorData

		err := res.Scan(&x.ID, &x.Name, &x.Specialis, &x.TypeSpecialis)
		if err != nil {
			return nil, err
		}

		d = append(d, x)
	}

	return d, nil
}

func (q *doctorRepository) GetDoctorSchedule(id string) ([]models.DoctorSchedule, error) {
	res, err := q.sql.Query("SELECT day, first_time, last_time FROM schedule WHERE doctor_id = $1", id)
	if err != nil {
		return nil, err
	}

	var data []models.DoctorSchedule
	for res.Next() {
		var d models.DoctorSchedule

		err := res.Scan(&d.Day, &d.FirstTime, &d.LastTime)
		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}
