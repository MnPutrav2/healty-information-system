package repository

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type CommonRepository interface {
	GetCurrentMedicalRecord() string
	GetCurrentRegisterNumber(date string, policlinic string) (int, error)
	GetCurrentCareNumber(date string) (int, error)
	GetPoliclinic() ([]models.Policlinics, error)
	AddRecipeNumber(care string) (string, error)
	GetCurrentRecipeNumber(date string) (int, error)
	GetCurrentLabNumber(date string) (int, error)
}

type commonRepository struct {
	w   http.ResponseWriter
	r   *http.Request
	sql *sql.DB
}

func NewCommonRepository(w http.ResponseWriter, r *http.Request, sql *sql.DB) CommonRepository {
	return &commonRepository{w, r, sql}
}

func (q *commonRepository) GetCurrentMedicalRecord() string {
	var mr string

	err := q.sql.QueryRow("SELECT COUNT(medical_record) + 1 FROM patients").Scan(&mr)
	if err != nil {
		panic(err.Error())
	}

	return mr
}

func (q *commonRepository) GetCurrentRegisterNumber(date string, policlinic string) (int, error) {
	var data int

	err := q.sql.QueryRow("SELECT COUNT(*) FROM registration WHERE register_date >= $1::date AND register_date < ($1::date + INTERVAL '1 day') AND registration.policlinic = $2", date, policlinic).Scan(&data)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (q *commonRepository) GetCurrentCareNumber(date string) (int, error) {
	var data int

	err := q.sql.QueryRow("SELECT COUNT(*) FROM registration WHERE register_date >= $1::date AND register_date < ($1::date + INTERVAL '1 day')", date).Scan(&data)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (q *commonRepository) GetPoliclinic() ([]models.Policlinics, error) {
	res, err := q.sql.Query("SELECT id, name, poli_type FROM policlinic")
	if err != nil {
		return nil, err
	}

	var d []models.Policlinics
	for res.Next() {
		var x models.Policlinics

		err := res.Scan(&x.ID, &x.Name, &x.PoliType)
		if err != nil {
			return nil, err
		}

		d = append(d, x)
	}

	return d, nil
}

func (q *commonRepository) AddRecipeNumber(care string) (string, error) {
	var rec string
	err := q.sql.QueryRow("SELECT recipe_id FROM recipes WHERE care_number = $1", care).Scan(&rec)
	if err != nil {
		return "", err
	}

	return rec, nil
}

func (q *commonRepository) GetCurrentRecipeNumber(date string) (int, error) {
	var current int
	dt := date + "%"
	err := q.sql.QueryRow("SELECT COUNT(*) FROM recipes WHERE date::date = $1", dt).Scan(&current)
	if err != nil {
		return 0, err
	}

	return current, nil
}

func (q *commonRepository) GetCurrentLabNumber(date string) (int, error) {
	var current int
	dt := date + "%"
	err := q.sql.QueryRow("SELECT COUNT(*) FROM labolatorium_request WHERE request_date::date = $1", dt).Scan(&current)
	if err != nil {
		return 0, err
	}

	return current, nil
}
