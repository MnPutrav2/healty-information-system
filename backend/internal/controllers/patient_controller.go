package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

func CreatePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	// get client request body
	patient, err := helper.GetRequestBodyPatientData(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.CreatePatientData(patient, val.Id, path); err != nil {
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, "", "error server", err.Error()+" : 500", 500, path)
		return
	}

	helper.ResponseSuccess(w, "", "create patient : 201", path, s, 201)
}

func GetPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	// get client request body
	param := r.URL.Query()

	limit := param.Get("limit")
	search := "%" + param.Get("search") + "%"

	patientRepo := repository.NewPatientRepository(w, r, sql)
	patients, err := patientRepo.GetPatientData(limit, search, val.Token, path)
	if err != nil {
		return
	}

	s, err := json.Marshal(patients)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get patient data ", path, s, 200)
}

func UpdatePatientData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	patient, err := helper.GetRequestBodyPatientDataUpdate(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.UpdatePatientData(patient); err != nil {
		helper.ResponseError(w, val.Id, "patient data not found", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "update patient data", path, s, 200)
}

func DeletePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	// get client request body
	param := r.URL.Query().Get("mr")

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.DeletePatientData(param); err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "delete patient data", path, s, 200)
}

func GetPatientStatus(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckPatientLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	patientRepo := repository.NewPatientRepository(w, r, sql)
	status, err := patientRepo.GetPatientAccountData(val.Token)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get user status", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(status)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get user status", path, s, 200)
}
