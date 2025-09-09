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

func CreateAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	care, err := helper.GetAmbulatoryRequest(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid json format", err.Error(), 400, path)
		return
	}

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err = ambulatoryRepo.CreateAmbulatoryCareData(care)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, val.Id, "create ambulatory care", path, s, 201)
}

func DeleteAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	query := r.URL.Query()
	careNum := query.Get("care-number")
	date := query.Get("date")

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err := ambulatoryRepo.DeleteAmbulatoryCareData(careNum, date)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed delete data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, val.Id, "delete ambulatory care", path, s, 200)
}

func GetAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	query := r.URL.Query()
	careNum := query.Get("care-number")
	date1 := query.Get("date1")
	date2 := query.Get("date2")

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	res, err := ambulatoryRepo.GetAmbulatoryCareData(careNum, date1, date2)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(res)
	helper.ResponseSuccess(w, val.Id, "get ambulatory care", path, s, 200)
}

func UpdateAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	care, err := helper.GetAmbulatoryRequestUpdate(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid json format", err.Error(), 400, path)
		return
	}

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err = ambulatoryRepo.UpdateAmbulatoryCareData(care)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "update"})
	helper.ResponseSuccess(w, val.Id, "update ambulatory care", path, s, 200)
}
