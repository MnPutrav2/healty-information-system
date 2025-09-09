package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetRequestBodyRegisterData(w http.ResponseWriter, r *http.Request, path string) (models.RequestRegisterPatient, error) {
	// get client request body
	var patient models.RequestRegisterPatient

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RequestRegisterPatient{}, err
	}

	return patient, nil
}

func GetRequestBodyRegisterUpdateStatus(w http.ResponseWriter, r *http.Request, path string) (models.RegisterStatus, error) {
	// get client request body
	var patient models.RegisterStatus

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RegisterStatus{}, err
	}

	return patient, nil
}
