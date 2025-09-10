package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetHumanResourceRequest(w http.ResponseWriter, r *http.Request) (models.EmployeeData, error) {
	// get client request body
	var patient models.EmployeeData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.EmployeeData{}, err
	}

	return patient, nil
}
