package repository

import (
	"database/sql"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type HumanResourceRepository interface {
	AddEmployeeData(data models.EmployeeData) error
	GetEmployeeData(qs string, l int) ([]models.EmployeeData, error)
	DeleteEmployeeData(id string) error
	UpdateEmployeeData(data models.EmployeeData, id string) error
}

type humanResourceRepository struct {
	sql *sql.DB
}

func NewHumanResourceRepository(sql *sql.DB) HumanResourceRepository {
	return &humanResourceRepository{sql}
}

func (q *humanResourceRepository) AddEmployeeData(data models.EmployeeData) error {
	_, err := q.sql.Exec("INSERT INTO employees(id, name, gender, birth_place, birth_date, wedding, address, nik, bpjs, npwp, phone_number, email) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", &data.ID, &data.Name, &data.Gender, &data.BirthPlace, &data.BirthDate, &data.Wedding, &data.Address, &data.NIK, &data.BPJS, &data.NPWP, &data.PhoneNumber, &data.Email)
	if err != nil {
		return err
	}

	return nil
}

func (q *humanResourceRepository) GetEmployeeData(qs string, l int) ([]models.EmployeeData, error) {
	res, err := q.sql.Query("SELECT id, name, gender, birth_place, birth_date, wedding, address, nik, bpjs, npwp, phone_number, email FROM employees WHERE name LIKE $1 LIMIT $2", qs, l)
	if err != nil {
		return nil, err
	}

	var datas []models.EmployeeData

	for res.Next() {
		var data models.EmployeeData

		err := res.Scan(&data.ID, &data.Name, &data.Gender, &data.BirthPlace, &data.BirthDate, &data.Wedding, &data.Address, &data.NIK, &data.BPJS, &data.NPWP, &data.PhoneNumber, &data.Email)

		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (q *humanResourceRepository) DeleteEmployeeData(id string) error {
	_, err := q.sql.Exec("DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (q *humanResourceRepository) UpdateEmployeeData(data models.EmployeeData, id string) error {
	_, err := q.sql.Exec("UPDATE employees SET id = $1, name = $2, gender = $3, birth_place = $4, birth_date = $5, wedding = $6, address = $7, nik = $8, bpjs = $9, npwp = $10, phone_number = $11, email = $12 WHERE id = $13", data.ID, data.Name, data.Gender, data.BirthPlace, data.BirthDate, data.Wedding, data.Address, data.NIK, data.BPJS, data.NPWP, data.PhoneNumber, data.Email, id)
	if err != nil {
		return err
	}

	return nil
}
