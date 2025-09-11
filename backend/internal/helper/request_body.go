package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetRequestBodyDrugData(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugData, error) {
	var drug models.RequestBodyDrugData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugData{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugDataUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugDataUpdate, error) {
	var drug models.RequestBodyDrugDataUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugDataUpdate{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipe(w http.ResponseWriter, r *http.Request, path string) (models.RecipeRequest, error) {
	var drug models.RecipeRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeRequest{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeCompound(w http.ResponseWriter, r *http.Request, path string) (models.RecipeCompoundRequest, error) {
	var drug models.RecipeCompoundRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeCompoundRequest{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeValidate(w http.ResponseWriter, r *http.Request, path string) (models.ValidateRecipe, error) {
	var drug models.ValidateRecipe

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.ValidateRecipe{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeHandover(w http.ResponseWriter, r *http.Request, path string) (models.RecipeHandover, error) {
	var drug models.RecipeHandover

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeHandover{}, err
	}

	return drug, nil
}

func GetAmbulatoryRequest(w http.ResponseWriter, r *http.Request, path string) (models.RequestAmbulatoryCare, error) {
	// get client request body
	var patient models.RequestAmbulatoryCare

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RequestAmbulatoryCare{}, err
	}

	return patient, nil
}

func GetAmbulatoryRequestUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestUpdateAmbulatorCare, error) {
	// get client request body
	var patient models.RequestUpdateAmbulatorCare

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RequestUpdateAmbulatorCare{}, err
	}

	return patient, nil
}

func GetExaminationRequest(w http.ResponseWriter, r *http.Request, path string) (models.ExaminationRequest, error) {
	// get client request body
	var patient models.ExaminationRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.ExaminationRequest{}, err
	}

	return patient, nil
}

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

func GetUserRequest(w http.ResponseWriter, r *http.Request) (models.UserReq, error) {
	// get client request body
	var patient models.UserReq

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.UserReq{}, err
	}

	return patient, nil
}

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

func GetRequestBodyUserAccount(w http.ResponseWriter, r *http.Request, path string) (models.UserAccount, error) {
	// get client request body
	var user models.UserAccount

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&user)
	if err != nil {
		return models.UserAccount{}, err
	}

	return user, nil
}

func GetRequestBodyPatientData(w http.ResponseWriter, r *http.Request, path string) (models.PatientData, error) {
	// get client request body
	var patient models.PatientData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.PatientData{}, err
	}

	return patient, nil
}

func GetRequestBodyPatientDataUpdate(w http.ResponseWriter, r *http.Request, path string) (models.PatientDataUpdate, error) {
	// get client request body
	var patient models.PatientDataUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.PatientDataUpdate{}, err
	}

	return patient, nil
}
