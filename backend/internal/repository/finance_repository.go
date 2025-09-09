package repository

import (
	"database/sql"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

type FinanceRepository interface {
	CreateExaminationData(data models.ExaminationCost) error
	GetExaminationData(search string, limit int) ([]models.ExaminationCost, error)
	DeleteExaminationData(id int) error
	UpdateExaminationData(data models.ExaminationCost, id int) error
}

type financeRepository struct {
	sql *sql.DB
}

func NewFinanceRepository(sql *sql.DB) FinanceRepository {
	return &financeRepository{sql}
}

func (q *financeRepository) CreateExaminationData(data models.ExaminationCost) error {
	_, err := q.sql.Exec("INSERT INTO examination_cost(id, examination_name, payment_method, doctor_cost, nurse_cost, management_cost) VALUES($1, $2, $3, $4, $5, $6)", data.ID, data.ExaminationName, data.PaymentMethod, data.DoctorCost, data.NurseCost, data.ManagementCost)
	if err != nil {
		return err
	}

	return nil
}

func (q *financeRepository) GetExaminationData(search string, limit int) ([]models.ExaminationCost, error) {
	res, err := q.sql.Query("SELECT id, examination_name, payment_method, doctor_cost, nurse_cost, management_cost FROM examination_cost WHERE examination_name LIKE $1 LIMIT $2", search, limit)
	if err != nil {
		return nil, err
	}

	var datas []models.ExaminationCost

	for res.Next() {
		var data models.ExaminationCost

		err := res.Scan(&data.ID, &data.ExaminationName, &data.PaymentMethod, &data.DoctorCost, &data.NurseCost, &data.ManagementCost)

		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (q *financeRepository) DeleteExaminationData(id int) error {
	_, err := q.sql.Exec("DELETE FROM examination_cost WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (q *financeRepository) UpdateExaminationData(data models.ExaminationCost, id int) error {
	_, err := q.sql.Exec("UPDATE examination_cost SET id = $1, examination_name = $2, payment_method = $3, doctor_cost = $4, nurse_cost = $5, management_cost = $6 WHERE id = $7", data.ID, data.ExaminationName, data.PaymentMethod, data.DoctorCost, data.NurseCost, data.ManagementCost, id)
	if err != nil {
		return err
	}

	return nil
}
