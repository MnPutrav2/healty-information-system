package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetExaminationRequestBody(w http.ResponseWriter, r *http.Request) (models.ExaminationCostRequest, error) {
	var exam models.ExaminationCostRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&exam)
	if err != nil {
		return models.ExaminationCostRequest{}, err
	}

	return exam, nil
}

func GetExaminationRequestBody2(w http.ResponseWriter, r *http.Request) (models.ExaminationCost, error) {
	var exam models.ExaminationCost

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&exam)
	if err != nil {
		return models.ExaminationCost{}, err
	}

	return exam, nil
}
