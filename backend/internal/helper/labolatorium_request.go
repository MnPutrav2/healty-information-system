package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetLabolatoriumResquestBody(w http.ResponseWriter, r *http.Request) (models.LabRequestData, error) {
	var lab models.LabRequestData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabRequestData{}, err
	}

	return lab, nil
}

func GetLabolatoriumResquestBodyData(w http.ResponseWriter, r *http.Request) (models.LabolatoriumRequest, error) {
	var lab models.LabolatoriumRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabolatoriumRequest{}, err
	}

	return lab, nil
}

func GetLabolatoriumTemplateResquestBody(w http.ResponseWriter, r *http.Request) (models.LabTemplateRequestData, error) {
	var lab models.LabTemplateRequestData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabTemplateRequestData{}, err
	}

	return lab, nil
}

func UpdateLabolatoriumResquestBody(w http.ResponseWriter, r *http.Request) (models.LabUpdate, error) {
	var lab models.LabUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabUpdate{}, err
	}

	return lab, nil
}

func UpdateLabolatoriumSampleRequestBody(w http.ResponseWriter, r *http.Request) (models.LabSampleRequest, error) {
	var lab models.LabSampleRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabSampleRequest{}, err
	}

	return lab, nil
}

func UpdateLabolatoriumValueRequestBody(w http.ResponseWriter, r *http.Request) (models.LabValueRequest, error) {
	var lab models.LabValueRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&lab)
	if err != nil {
		return models.LabValueRequest{}, err
	}

	return lab, nil
}
